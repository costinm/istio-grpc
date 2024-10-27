// Copyright 2015 gRPC authors.
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

// An integration test service that covers all the method signature permutations
// of unary/streaming requests/responses.

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: grpc/testing/report_qps_scenario_service.proto

package testingconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	testing "github.com/costinm/grpc-mesh/gen/proto/go/grpc/testing"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ReportQpsScenarioServiceName is the fully-qualified name of the ReportQpsScenarioService service.
	ReportQpsScenarioServiceName = "grpc.testing.ReportQpsScenarioService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ReportQpsScenarioServiceReportScenarioProcedure is the fully-qualified name of the
	// ReportQpsScenarioService's ReportScenario RPC.
	ReportQpsScenarioServiceReportScenarioProcedure = "/grpc.testing.ReportQpsScenarioService/ReportScenario"
)

// ReportQpsScenarioServiceClient is a client for the grpc.testing.ReportQpsScenarioService service.
type ReportQpsScenarioServiceClient interface {
	// Report results of a QPS test benchmark scenario.
	ReportScenario(context.Context, *connect_go.Request[testing.ScenarioResult]) (*connect_go.Response[testing.Void], error)
}

// NewReportQpsScenarioServiceClient constructs a client for the
// grpc.testing.ReportQpsScenarioService service. By default, it uses the Connect protocol with the
// binary Protobuf Codec, asks for gzipped responses, and sends uncompressed requests. To use the
// gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewReportQpsScenarioServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ReportQpsScenarioServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &reportQpsScenarioServiceClient{
		reportScenario: connect_go.NewClient[testing.ScenarioResult, testing.Void](
			httpClient,
			baseURL+ReportQpsScenarioServiceReportScenarioProcedure,
			opts...,
		),
	}
}

// reportQpsScenarioServiceClient implements ReportQpsScenarioServiceClient.
type reportQpsScenarioServiceClient struct {
	reportScenario *connect_go.Client[testing.ScenarioResult, testing.Void]
}

// ReportScenario calls grpc.testing.ReportQpsScenarioService.ReportScenario.
func (c *reportQpsScenarioServiceClient) ReportScenario(ctx context.Context, req *connect_go.Request[testing.ScenarioResult]) (*connect_go.Response[testing.Void], error) {
	return c.reportScenario.CallUnary(ctx, req)
}

// ReportQpsScenarioServiceHandler is an implementation of the grpc.testing.ReportQpsScenarioService
// service.
type ReportQpsScenarioServiceHandler interface {
	// Report results of a QPS test benchmark scenario.
	ReportScenario(context.Context, *connect_go.Request[testing.ScenarioResult]) (*connect_go.Response[testing.Void], error)
}

// NewReportQpsScenarioServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewReportQpsScenarioServiceHandler(svc ReportQpsScenarioServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	reportQpsScenarioServiceReportScenarioHandler := connect_go.NewUnaryHandler(
		ReportQpsScenarioServiceReportScenarioProcedure,
		svc.ReportScenario,
		opts...,
	)
	return "/grpc.testing.ReportQpsScenarioService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ReportQpsScenarioServiceReportScenarioProcedure:
			reportQpsScenarioServiceReportScenarioHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedReportQpsScenarioServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedReportQpsScenarioServiceHandler struct{}

func (UnimplementedReportQpsScenarioServiceHandler) ReportScenario(context.Context, *connect_go.Request[testing.ScenarioResult]) (*connect_go.Response[testing.Void], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("grpc.testing.ReportQpsScenarioService.ReportScenario is not implemented"))
}
