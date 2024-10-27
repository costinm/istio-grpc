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
// Source: grpc/testing/benchmark_service.proto

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
	// BenchmarkServiceName is the fully-qualified name of the BenchmarkService service.
	BenchmarkServiceName = "grpc.testing.BenchmarkService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// BenchmarkServiceUnaryCallProcedure is the fully-qualified name of the BenchmarkService's
	// UnaryCall RPC.
	BenchmarkServiceUnaryCallProcedure = "/grpc.testing.BenchmarkService/UnaryCall"
	// BenchmarkServiceStreamingCallProcedure is the fully-qualified name of the BenchmarkService's
	// StreamingCall RPC.
	BenchmarkServiceStreamingCallProcedure = "/grpc.testing.BenchmarkService/StreamingCall"
	// BenchmarkServiceStreamingFromClientProcedure is the fully-qualified name of the
	// BenchmarkService's StreamingFromClient RPC.
	BenchmarkServiceStreamingFromClientProcedure = "/grpc.testing.BenchmarkService/StreamingFromClient"
	// BenchmarkServiceStreamingFromServerProcedure is the fully-qualified name of the
	// BenchmarkService's StreamingFromServer RPC.
	BenchmarkServiceStreamingFromServerProcedure = "/grpc.testing.BenchmarkService/StreamingFromServer"
	// BenchmarkServiceStreamingBothWaysProcedure is the fully-qualified name of the BenchmarkService's
	// StreamingBothWays RPC.
	BenchmarkServiceStreamingBothWaysProcedure = "/grpc.testing.BenchmarkService/StreamingBothWays"
)

// BenchmarkServiceClient is a client for the grpc.testing.BenchmarkService service.
type BenchmarkServiceClient interface {
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(context.Context, *connect_go.Request[testing.SimpleRequest]) (*connect_go.Response[testing.SimpleResponse], error)
	// Repeated sequence of one request followed by one response.
	// Should be called streaming ping-pong
	// The server returns the client payload as-is on each response
	StreamingCall(context.Context) *connect_go.BidiStreamForClient[testing.SimpleRequest, testing.SimpleResponse]
	// Single-sided unbounded streaming from client to server
	// The server returns the client payload as-is once the client does WritesDone
	StreamingFromClient(context.Context) *connect_go.ClientStreamForClient[testing.SimpleRequest, testing.SimpleResponse]
	// Single-sided unbounded streaming from server to client
	// The server repeatedly returns the client payload as-is
	StreamingFromServer(context.Context, *connect_go.Request[testing.SimpleRequest]) (*connect_go.ServerStreamForClient[testing.SimpleResponse], error)
	// Two-sided unbounded streaming between server to client
	// Both sides send the content of their own choice to the other
	StreamingBothWays(context.Context) *connect_go.BidiStreamForClient[testing.SimpleRequest, testing.SimpleResponse]
}

// NewBenchmarkServiceClient constructs a client for the grpc.testing.BenchmarkService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBenchmarkServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) BenchmarkServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &benchmarkServiceClient{
		unaryCall: connect_go.NewClient[testing.SimpleRequest, testing.SimpleResponse](
			httpClient,
			baseURL+BenchmarkServiceUnaryCallProcedure,
			opts...,
		),
		streamingCall: connect_go.NewClient[testing.SimpleRequest, testing.SimpleResponse](
			httpClient,
			baseURL+BenchmarkServiceStreamingCallProcedure,
			opts...,
		),
		streamingFromClient: connect_go.NewClient[testing.SimpleRequest, testing.SimpleResponse](
			httpClient,
			baseURL+BenchmarkServiceStreamingFromClientProcedure,
			opts...,
		),
		streamingFromServer: connect_go.NewClient[testing.SimpleRequest, testing.SimpleResponse](
			httpClient,
			baseURL+BenchmarkServiceStreamingFromServerProcedure,
			opts...,
		),
		streamingBothWays: connect_go.NewClient[testing.SimpleRequest, testing.SimpleResponse](
			httpClient,
			baseURL+BenchmarkServiceStreamingBothWaysProcedure,
			opts...,
		),
	}
}

// benchmarkServiceClient implements BenchmarkServiceClient.
type benchmarkServiceClient struct {
	unaryCall           *connect_go.Client[testing.SimpleRequest, testing.SimpleResponse]
	streamingCall       *connect_go.Client[testing.SimpleRequest, testing.SimpleResponse]
	streamingFromClient *connect_go.Client[testing.SimpleRequest, testing.SimpleResponse]
	streamingFromServer *connect_go.Client[testing.SimpleRequest, testing.SimpleResponse]
	streamingBothWays   *connect_go.Client[testing.SimpleRequest, testing.SimpleResponse]
}

// UnaryCall calls grpc.testing.BenchmarkService.UnaryCall.
func (c *benchmarkServiceClient) UnaryCall(ctx context.Context, req *connect_go.Request[testing.SimpleRequest]) (*connect_go.Response[testing.SimpleResponse], error) {
	return c.unaryCall.CallUnary(ctx, req)
}

// StreamingCall calls grpc.testing.BenchmarkService.StreamingCall.
func (c *benchmarkServiceClient) StreamingCall(ctx context.Context) *connect_go.BidiStreamForClient[testing.SimpleRequest, testing.SimpleResponse] {
	return c.streamingCall.CallBidiStream(ctx)
}

// StreamingFromClient calls grpc.testing.BenchmarkService.StreamingFromClient.
func (c *benchmarkServiceClient) StreamingFromClient(ctx context.Context) *connect_go.ClientStreamForClient[testing.SimpleRequest, testing.SimpleResponse] {
	return c.streamingFromClient.CallClientStream(ctx)
}

// StreamingFromServer calls grpc.testing.BenchmarkService.StreamingFromServer.
func (c *benchmarkServiceClient) StreamingFromServer(ctx context.Context, req *connect_go.Request[testing.SimpleRequest]) (*connect_go.ServerStreamForClient[testing.SimpleResponse], error) {
	return c.streamingFromServer.CallServerStream(ctx, req)
}

// StreamingBothWays calls grpc.testing.BenchmarkService.StreamingBothWays.
func (c *benchmarkServiceClient) StreamingBothWays(ctx context.Context) *connect_go.BidiStreamForClient[testing.SimpleRequest, testing.SimpleResponse] {
	return c.streamingBothWays.CallBidiStream(ctx)
}

// BenchmarkServiceHandler is an implementation of the grpc.testing.BenchmarkService service.
type BenchmarkServiceHandler interface {
	// One request followed by one response.
	// The server returns the client payload as-is.
	UnaryCall(context.Context, *connect_go.Request[testing.SimpleRequest]) (*connect_go.Response[testing.SimpleResponse], error)
	// Repeated sequence of one request followed by one response.
	// Should be called streaming ping-pong
	// The server returns the client payload as-is on each response
	StreamingCall(context.Context, *connect_go.BidiStream[testing.SimpleRequest, testing.SimpleResponse]) error
	// Single-sided unbounded streaming from client to server
	// The server returns the client payload as-is once the client does WritesDone
	StreamingFromClient(context.Context, *connect_go.ClientStream[testing.SimpleRequest]) (*connect_go.Response[testing.SimpleResponse], error)
	// Single-sided unbounded streaming from server to client
	// The server repeatedly returns the client payload as-is
	StreamingFromServer(context.Context, *connect_go.Request[testing.SimpleRequest], *connect_go.ServerStream[testing.SimpleResponse]) error
	// Two-sided unbounded streaming between server to client
	// Both sides send the content of their own choice to the other
	StreamingBothWays(context.Context, *connect_go.BidiStream[testing.SimpleRequest, testing.SimpleResponse]) error
}

// NewBenchmarkServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBenchmarkServiceHandler(svc BenchmarkServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	benchmarkServiceUnaryCallHandler := connect_go.NewUnaryHandler(
		BenchmarkServiceUnaryCallProcedure,
		svc.UnaryCall,
		opts...,
	)
	benchmarkServiceStreamingCallHandler := connect_go.NewBidiStreamHandler(
		BenchmarkServiceStreamingCallProcedure,
		svc.StreamingCall,
		opts...,
	)
	benchmarkServiceStreamingFromClientHandler := connect_go.NewClientStreamHandler(
		BenchmarkServiceStreamingFromClientProcedure,
		svc.StreamingFromClient,
		opts...,
	)
	benchmarkServiceStreamingFromServerHandler := connect_go.NewServerStreamHandler(
		BenchmarkServiceStreamingFromServerProcedure,
		svc.StreamingFromServer,
		opts...,
	)
	benchmarkServiceStreamingBothWaysHandler := connect_go.NewBidiStreamHandler(
		BenchmarkServiceStreamingBothWaysProcedure,
		svc.StreamingBothWays,
		opts...,
	)
	return "/grpc.testing.BenchmarkService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case BenchmarkServiceUnaryCallProcedure:
			benchmarkServiceUnaryCallHandler.ServeHTTP(w, r)
		case BenchmarkServiceStreamingCallProcedure:
			benchmarkServiceStreamingCallHandler.ServeHTTP(w, r)
		case BenchmarkServiceStreamingFromClientProcedure:
			benchmarkServiceStreamingFromClientHandler.ServeHTTP(w, r)
		case BenchmarkServiceStreamingFromServerProcedure:
			benchmarkServiceStreamingFromServerHandler.ServeHTTP(w, r)
		case BenchmarkServiceStreamingBothWaysProcedure:
			benchmarkServiceStreamingBothWaysHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedBenchmarkServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBenchmarkServiceHandler struct{}

func (UnimplementedBenchmarkServiceHandler) UnaryCall(context.Context, *connect_go.Request[testing.SimpleRequest]) (*connect_go.Response[testing.SimpleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("grpc.testing.BenchmarkService.UnaryCall is not implemented"))
}

func (UnimplementedBenchmarkServiceHandler) StreamingCall(context.Context, *connect_go.BidiStream[testing.SimpleRequest, testing.SimpleResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("grpc.testing.BenchmarkService.StreamingCall is not implemented"))
}

func (UnimplementedBenchmarkServiceHandler) StreamingFromClient(context.Context, *connect_go.ClientStream[testing.SimpleRequest]) (*connect_go.Response[testing.SimpleResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("grpc.testing.BenchmarkService.StreamingFromClient is not implemented"))
}

func (UnimplementedBenchmarkServiceHandler) StreamingFromServer(context.Context, *connect_go.Request[testing.SimpleRequest], *connect_go.ServerStream[testing.SimpleResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("grpc.testing.BenchmarkService.StreamingFromServer is not implemented"))
}

func (UnimplementedBenchmarkServiceHandler) StreamingBothWays(context.Context, *connect_go.BidiStream[testing.SimpleRequest, testing.SimpleResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("grpc.testing.BenchmarkService.StreamingBothWays is not implemented"))
}
