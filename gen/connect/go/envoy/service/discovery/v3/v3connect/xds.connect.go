// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: envoy/service/discovery/v3/xds.proto

// GRPC package - part of the URL. Service is added.
// URL: /PACKAGE.SERVICE/METHOD
package v3connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	_ "github.com/costinm/grpc-mesh/gen/proto/go/envoy/service/discovery/v3"
	xds "github.com/costinm/grpc-mesh/gen/proto/go/xds"
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
	// AggregatedDiscoveryServiceName is the fully-qualified name of the AggregatedDiscoveryService
	// service.
	AggregatedDiscoveryServiceName = "envoy.service.discovery.v3.AggregatedDiscoveryService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// AggregatedDiscoveryServiceStreamAggregatedResourcesProcedure is the fully-qualified name of the
	// AggregatedDiscoveryService's StreamAggregatedResources RPC.
	AggregatedDiscoveryServiceStreamAggregatedResourcesProcedure = "/envoy.service.discovery.v3.AggregatedDiscoveryService/StreamAggregatedResources"
	// AggregatedDiscoveryServiceDeltaAggregatedResourcesProcedure is the fully-qualified name of the
	// AggregatedDiscoveryService's DeltaAggregatedResources RPC.
	AggregatedDiscoveryServiceDeltaAggregatedResourcesProcedure = "/envoy.service.discovery.v3.AggregatedDiscoveryService/DeltaAggregatedResources"
)

// AggregatedDiscoveryServiceClient is a client for the
// envoy.service.discovery.v3.AggregatedDiscoveryService service.
type AggregatedDiscoveryServiceClient interface {
	StreamAggregatedResources(context.Context) *connect_go.BidiStreamForClient[xds.DiscoveryRequest, xds.DiscoveryResponse]
	DeltaAggregatedResources(context.Context) *connect_go.BidiStreamForClient[xds.DeltaDiscoveryRequest, xds.DeltaDiscoveryResponse]
}

// NewAggregatedDiscoveryServiceClient constructs a client for the
// envoy.service.discovery.v3.AggregatedDiscoveryService service. By default, it uses the Connect
// protocol with the binary Protobuf Codec, asks for gzipped responses, and sends uncompressed
// requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewAggregatedDiscoveryServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) AggregatedDiscoveryServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &aggregatedDiscoveryServiceClient{
		streamAggregatedResources: connect_go.NewClient[xds.DiscoveryRequest, xds.DiscoveryResponse](
			httpClient,
			baseURL+AggregatedDiscoveryServiceStreamAggregatedResourcesProcedure,
			opts...,
		),
		deltaAggregatedResources: connect_go.NewClient[xds.DeltaDiscoveryRequest, xds.DeltaDiscoveryResponse](
			httpClient,
			baseURL+AggregatedDiscoveryServiceDeltaAggregatedResourcesProcedure,
			opts...,
		),
	}
}

// aggregatedDiscoveryServiceClient implements AggregatedDiscoveryServiceClient.
type aggregatedDiscoveryServiceClient struct {
	streamAggregatedResources *connect_go.Client[xds.DiscoveryRequest, xds.DiscoveryResponse]
	deltaAggregatedResources  *connect_go.Client[xds.DeltaDiscoveryRequest, xds.DeltaDiscoveryResponse]
}

// StreamAggregatedResources calls
// envoy.service.discovery.v3.AggregatedDiscoveryService.StreamAggregatedResources.
func (c *aggregatedDiscoveryServiceClient) StreamAggregatedResources(ctx context.Context) *connect_go.BidiStreamForClient[xds.DiscoveryRequest, xds.DiscoveryResponse] {
	return c.streamAggregatedResources.CallBidiStream(ctx)
}

// DeltaAggregatedResources calls
// envoy.service.discovery.v3.AggregatedDiscoveryService.DeltaAggregatedResources.
func (c *aggregatedDiscoveryServiceClient) DeltaAggregatedResources(ctx context.Context) *connect_go.BidiStreamForClient[xds.DeltaDiscoveryRequest, xds.DeltaDiscoveryResponse] {
	return c.deltaAggregatedResources.CallBidiStream(ctx)
}

// AggregatedDiscoveryServiceHandler is an implementation of the
// envoy.service.discovery.v3.AggregatedDiscoveryService service.
type AggregatedDiscoveryServiceHandler interface {
	StreamAggregatedResources(context.Context, *connect_go.BidiStream[xds.DiscoveryRequest, xds.DiscoveryResponse]) error
	DeltaAggregatedResources(context.Context, *connect_go.BidiStream[xds.DeltaDiscoveryRequest, xds.DeltaDiscoveryResponse]) error
}

// NewAggregatedDiscoveryServiceHandler builds an HTTP handler from the service implementation. It
// returns the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewAggregatedDiscoveryServiceHandler(svc AggregatedDiscoveryServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	aggregatedDiscoveryServiceStreamAggregatedResourcesHandler := connect_go.NewBidiStreamHandler(
		AggregatedDiscoveryServiceStreamAggregatedResourcesProcedure,
		svc.StreamAggregatedResources,
		opts...,
	)
	aggregatedDiscoveryServiceDeltaAggregatedResourcesHandler := connect_go.NewBidiStreamHandler(
		AggregatedDiscoveryServiceDeltaAggregatedResourcesProcedure,
		svc.DeltaAggregatedResources,
		opts...,
	)
	return "/envoy.service.discovery.v3.AggregatedDiscoveryService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case AggregatedDiscoveryServiceStreamAggregatedResourcesProcedure:
			aggregatedDiscoveryServiceStreamAggregatedResourcesHandler.ServeHTTP(w, r)
		case AggregatedDiscoveryServiceDeltaAggregatedResourcesProcedure:
			aggregatedDiscoveryServiceDeltaAggregatedResourcesHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedAggregatedDiscoveryServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedAggregatedDiscoveryServiceHandler struct{}

func (UnimplementedAggregatedDiscoveryServiceHandler) StreamAggregatedResources(context.Context, *connect_go.BidiStream[xds.DiscoveryRequest, xds.DiscoveryResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("envoy.service.discovery.v3.AggregatedDiscoveryService.StreamAggregatedResources is not implemented"))
}

func (UnimplementedAggregatedDiscoveryServiceHandler) DeltaAggregatedResources(context.Context, *connect_go.BidiStream[xds.DeltaDiscoveryRequest, xds.DeltaDiscoveryResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("envoy.service.discovery.v3.AggregatedDiscoveryService.DeltaAggregatedResources is not implemented"))
}
