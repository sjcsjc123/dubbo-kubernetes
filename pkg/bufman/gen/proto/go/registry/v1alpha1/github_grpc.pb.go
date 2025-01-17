// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: registry/v1alpha1/github.proto

package registryv1alpha1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GithubService_GetGithubAppConfig_FullMethodName = "/bufman.dubbo.apache.org.registry.v1alpha1.GithubService/GetGithubAppConfig"
)

// GithubServiceClient is the client API for GithubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GithubServiceClient interface {
	// GetGithubAppConfig returns a Github Application Configuration.
	GetGithubAppConfig(ctx context.Context, in *GetGithubAppConfigRequest, opts ...grpc.CallOption) (*GetGithubAppConfigResponse, error)
}

type githubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGithubServiceClient(cc grpc.ClientConnInterface) GithubServiceClient {
	return &githubServiceClient{cc}
}

func (c *githubServiceClient) GetGithubAppConfig(ctx context.Context, in *GetGithubAppConfigRequest, opts ...grpc.CallOption) (*GetGithubAppConfigResponse, error) {
	out := new(GetGithubAppConfigResponse)
	err := c.cc.Invoke(ctx, GithubService_GetGithubAppConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GithubServiceServer is the server API for GithubService service.
// All implementations must embed UnimplementedGithubServiceServer
// for forward compatibility
type GithubServiceServer interface {
	// GetGithubAppConfig returns a Github Application Configuration.
	GetGithubAppConfig(context.Context, *GetGithubAppConfigRequest) (*GetGithubAppConfigResponse, error)
	mustEmbedUnimplementedGithubServiceServer()
}

// UnimplementedGithubServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGithubServiceServer struct {
}

func (UnimplementedGithubServiceServer) GetGithubAppConfig(context.Context, *GetGithubAppConfigRequest) (*GetGithubAppConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGithubAppConfig not implemented")
}
func (UnimplementedGithubServiceServer) mustEmbedUnimplementedGithubServiceServer() {}

// UnsafeGithubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GithubServiceServer will
// result in compilation errors.
type UnsafeGithubServiceServer interface {
	mustEmbedUnimplementedGithubServiceServer()
}

func RegisterGithubServiceServer(s grpc.ServiceRegistrar, srv GithubServiceServer) {
	s.RegisterService(&GithubService_ServiceDesc, srv)
}

func _GithubService_GetGithubAppConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGithubAppConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GithubServiceServer).GetGithubAppConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GithubService_GetGithubAppConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GithubServiceServer).GetGithubAppConfig(ctx, req.(*GetGithubAppConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GithubService_ServiceDesc is the grpc.ServiceDesc for GithubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GithubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bufman.dubbo.apache.org.registry.v1alpha1.GithubService",
	HandlerType: (*GithubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGithubAppConfig",
			Handler:    _GithubService_GetGithubAppConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/v1alpha1/github.proto",
}
