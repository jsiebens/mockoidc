// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: mockoidc/v1/mockoidc.proto

package mockoidcv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/jsiebens/mockoidc/pkg/gen/mockoidc/v1"
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
	// MockOIDCServiceName is the fully-qualified name of the MockOIDCService service.
	MockOIDCServiceName = "mockoidc.v1.MockOIDCService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// MockOIDCServicePushUserProcedure is the fully-qualified name of the MockOIDCService's PushUser
	// RPC.
	MockOIDCServicePushUserProcedure = "/mockoidc.v1.MockOIDCService/PushUser"
)

// MockOIDCServiceClient is a client for the mockoidc.v1.MockOIDCService service.
type MockOIDCServiceClient interface {
	PushUser(context.Context, *connect_go.Request[v1.PushUserRequest]) (*connect_go.Response[v1.PushUserResponse], error)
}

// NewMockOIDCServiceClient constructs a client for the mockoidc.v1.MockOIDCService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewMockOIDCServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) MockOIDCServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &mockOIDCServiceClient{
		pushUser: connect_go.NewClient[v1.PushUserRequest, v1.PushUserResponse](
			httpClient,
			baseURL+MockOIDCServicePushUserProcedure,
			opts...,
		),
	}
}

// mockOIDCServiceClient implements MockOIDCServiceClient.
type mockOIDCServiceClient struct {
	pushUser *connect_go.Client[v1.PushUserRequest, v1.PushUserResponse]
}

// PushUser calls mockoidc.v1.MockOIDCService.PushUser.
func (c *mockOIDCServiceClient) PushUser(ctx context.Context, req *connect_go.Request[v1.PushUserRequest]) (*connect_go.Response[v1.PushUserResponse], error) {
	return c.pushUser.CallUnary(ctx, req)
}

// MockOIDCServiceHandler is an implementation of the mockoidc.v1.MockOIDCService service.
type MockOIDCServiceHandler interface {
	PushUser(context.Context, *connect_go.Request[v1.PushUserRequest]) (*connect_go.Response[v1.PushUserResponse], error)
}

// NewMockOIDCServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewMockOIDCServiceHandler(svc MockOIDCServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mockOIDCServicePushUserHandler := connect_go.NewUnaryHandler(
		MockOIDCServicePushUserProcedure,
		svc.PushUser,
		opts...,
	)
	return "/mockoidc.v1.MockOIDCService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case MockOIDCServicePushUserProcedure:
			mockOIDCServicePushUserHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedMockOIDCServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedMockOIDCServiceHandler struct{}

func (UnimplementedMockOIDCServiceHandler) PushUser(context.Context, *connect_go.Request[v1.PushUserRequest]) (*connect_go.Response[v1.PushUserResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("mockoidc.v1.MockOIDCService.PushUser is not implemented"))
}