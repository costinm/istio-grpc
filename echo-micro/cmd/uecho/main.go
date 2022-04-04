// Copyright Istio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	ocprom "contrib.go.opencensus.io/exporter/prometheus"
	"github.com/costinm/grpc-mesh/bootstrap"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/plugin/runmetrics"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/zpages"

	"github.com/costinm/grpc-mesh/echo-micro/server"

	//grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_zap "github.com/costinm/grpc-mesh/telemetry/logs/zap"
	"go.uber.org/zap"

	// Instrumentations
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/admin"
	"google.golang.org/grpc/credentials/insecure"
	xdscreds "google.golang.org/grpc/credentials/xds"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/health"
	grpcHealth "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/xds"
)

// Istio echo server with:
// - telemetry/traces (prom/zpages)
// - XDS
// - zap logs
//

var log = grpclog.Component("echo")

type GRPCServer interface {
	RegisterService(*grpc.ServiceDesc, interface{})
	Serve(net.Listener) error
	Stop()
	GracefulStop()
	GetServiceInfo() map[string]grpc.ServiceInfo
}

// TODO: 2 servers, one XDS and one plain
// TODO: get certs, remote config, JWTs
// TODO: tunnels
func Run(lis net.Listener) (func(), error) {
	zl, _ := zap.NewDevelopment(zap.AddCallerSkip(4))
	grpc_zap.ReplaceGrpcLoggerV2WithVerbosity(zl, 99)

	alwaysLoggingDeciderServer := func(ctx context.Context, fullMethodName string, servingObject interface{}) bool { return true }
	alwaysLoggingDeciderClient := func(ctx context.Context, fullMethodName string) bool { return true }

	h := &server.EchoGrpcHandler{
		// Enable OpenTelemetry client side
		DialOptions: []grpc.DialOption{
			grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
			grpc.WithStreamInterceptor(otelgrpc.StreamClientInterceptor()),
			grpc.WithUnaryInterceptor(grpc_zap.UnaryClientInterceptor(zl)),
			grpc.WithStreamInterceptor(grpc_zap.StreamClientInterceptor(zl)),
			grpc.WithUnaryInterceptor(grpc_zap.PayloadUnaryClientInterceptor(zl, alwaysLoggingDeciderClient)),
			grpc.WithStreamInterceptor(grpc_zap.PayloadStreamClientInterceptor(zl, alwaysLoggingDeciderClient)),
		},
	}

	cleanup, err := initTel(context.Background(), "echo")
	if err != nil {
		return nil, err
	}
	var grpcServer GRPCServer
	if os.Getenv("GRPC_XDS_BOOTSTRAP") != "" {
		// Only needs to be set to a file - if the file doesn't exist, create it.
		bootstrap.Generate(&bootstrap.GenerateBootstrapOptions{})
		creds, _ := xdscreds.NewServerCredentials(xdscreds.ServerOptions{FallbackCreds: insecure.NewCredentials()})

		grpcOptions := []grpc.ServerOption{
			grpc.Creds(creds),
			grpc_middleware.WithStreamServerChain(
				otelgrpc.StreamServerInterceptor(),
				grpc_zap.StreamServerInterceptor(zl),
				grpc_zap.PayloadStreamServerInterceptor(zl, alwaysLoggingDeciderServer)),
			grpc_middleware.WithUnaryServerChain(
				otelgrpc.UnaryServerInterceptor(),
				grpc_zap.UnaryServerInterceptor(zl),
				grpc_zap.PayloadUnaryServerInterceptor(zl, alwaysLoggingDeciderServer)),
		}

		// Replaces: grpc.NewServer(grpcOptions...)
		grpcServer = xds.NewGRPCServer(grpcOptions...)

	} else {
		creds := insecure.NewCredentials()

		grpcOptions := []grpc.ServerOption{
			grpc.Creds(creds),
			grpc.StreamInterceptor(otelgrpc.StreamServerInterceptor()),
			grpc.UnaryInterceptor(otelgrpc.UnaryServerInterceptor()),
			grpc.UnaryInterceptor(grpc_zap.UnaryServerInterceptor(zl)),
			grpc.StreamInterceptor(grpc_zap.StreamServerInterceptor(zl)),
		}

		grpcServer = grpc.NewServer(grpcOptions...)
	}

	// Special handling for startup without env variable set

	// Generate the bootstrap if the file is missing ( injection-less )
	// using cloudrun-mesh auto-detection code

	// Generate certs if missing

	h.Register(grpcServer)

	// add the standard grpc health check
	healthServer := health.NewServer()
	grpcHealth.RegisterHealthServer(grpcServer, healthServer)
	reflection.Register(grpcServer)

	// grpcdebug support
	_, err = admin.Register(grpcServer)
	if err != nil {
		log.Info("Failed to register admin", "error", err)
	}

	// Status
	go http.ListenAndServe("127.0.0.1:9081", http.DefaultServeMux)

	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatal(err)
		}
	}()
	return cleanup, nil
}

func initTel(background context.Context, s string) (func(), error) {
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Errorf("Failed to register ocgrpc server views: %v", err)
	}
	if err := view.Register(ocgrpc.DefaultClientViews...); err != nil {
		log.Errorf("Failed to register ocgrpc server views: %v", err)
	}

	// Similar with pilot-agent
	registry := prometheus.NewRegistry()
	wrapped := prometheus.WrapRegistererWithPrefix("uecho_",
		prometheus.Registerer(registry))

	wrapped.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	wrapped.MustRegister(collectors.NewGoCollector())

	//promRegistry = registry
	// go collector metrics collide with other metrics.
	exporter, err := ocprom.NewExporter(ocprom.Options{Registry: registry,
		Registerer: wrapped})
	if err != nil {
		log.Fatalf("could not setup exporter: %v", err)
	}
	view.RegisterExporter(exporter)
	zpages.Handle(http.DefaultServeMux, "/debug")

	http.Handle("/metrics", exporter)
	err = runmetrics.Enable(runmetrics.RunMetricOptions{
		EnableCPU:    true,
		EnableMemory: true,
		Prefix:       "echo/",
	})
	if err != nil {
		log.Error(err)
	}

	return func() {

	}, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9080"
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	c, err := Run(lis)
	if err != nil {
		log.Fatal(err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs

	c()
	// TODO: lame duck, etc
}
