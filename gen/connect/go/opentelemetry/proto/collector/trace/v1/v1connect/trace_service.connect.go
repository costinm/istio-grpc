// Copyright 2019, OpenTelemetry Authors
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

// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: opentelemetry/proto/collector/trace/v1/trace_service.proto

package v1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "go.opentelemetry.io/proto/otlp/collector/trace/v1"
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
	// TraceServiceName is the fully-qualified name of the TraceService service.
	TraceServiceName = "opentelemetry.proto.collector.trace.v1.TraceService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// TraceServiceExportProcedure is the fully-qualified name of the TraceService's Export RPC.
	TraceServiceExportProcedure = "/opentelemetry.proto.collector.trace.v1.TraceService/Export"
)

// TraceServiceClient is a client for the opentelemetry.proto.collector.trace.v1.TraceService
// service.
type TraceServiceClient interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *connect_go.Request[v1.ExportTraceServiceRequest]) (*connect_go.Response[v1.ExportTraceServiceResponse], error)
}

// NewTraceServiceClient constructs a client for the
// opentelemetry.proto.collector.trace.v1.TraceService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewTraceServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) TraceServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &traceServiceClient{
		export: connect_go.NewClient[v1.ExportTraceServiceRequest, v1.ExportTraceServiceResponse](
			httpClient,
			baseURL+TraceServiceExportProcedure,
			opts...,
		),
	}
}

// traceServiceClient implements TraceServiceClient.
type traceServiceClient struct {
	export *connect_go.Client[v1.ExportTraceServiceRequest, v1.ExportTraceServiceResponse]
}

// Export calls opentelemetry.proto.collector.trace.v1.TraceService.Export.
func (c *traceServiceClient) Export(ctx context.Context, req *connect_go.Request[v1.ExportTraceServiceRequest]) (*connect_go.Response[v1.ExportTraceServiceResponse], error) {
	return c.export.CallUnary(ctx, req)
}

// TraceServiceHandler is an implementation of the
// opentelemetry.proto.collector.trace.v1.TraceService service.
type TraceServiceHandler interface {
	// For performance reasons, it is recommended to keep this RPC
	// alive for the entire life of the application.
	Export(context.Context, *connect_go.Request[v1.ExportTraceServiceRequest]) (*connect_go.Response[v1.ExportTraceServiceResponse], error)
}

// NewTraceServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewTraceServiceHandler(svc TraceServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	traceServiceExportHandler := connect_go.NewUnaryHandler(
		TraceServiceExportProcedure,
		svc.Export,
		opts...,
	)
	return "/opentelemetry.proto.collector.trace.v1.TraceService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case TraceServiceExportProcedure:
			traceServiceExportHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedTraceServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedTraceServiceHandler struct{}

func (UnimplementedTraceServiceHandler) Export(context.Context, *connect_go.Request[v1.ExportTraceServiceRequest]) (*connect_go.Response[v1.ExportTraceServiceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("opentelemetry.proto.collector.trace.v1.TraceService.Export is not implemented"))
}