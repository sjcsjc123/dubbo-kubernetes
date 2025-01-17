// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: registry/v1alpha1/display.proto

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
	DisplayService_DisplayOrganizationElements_FullMethodName       = "/bufman.dubbo.apache.org.registry.v1alpha1.DisplayService/DisplayOrganizationElements"
	DisplayService_DisplayRepositoryElements_FullMethodName         = "/bufman.dubbo.apache.org.registry.v1alpha1.DisplayService/DisplayRepositoryElements"
	DisplayService_DisplayUserElements_FullMethodName               = "/bufman.dubbo.apache.org.registry.v1alpha1.DisplayService/DisplayUserElements"
	DisplayService_DisplayServerElements_FullMethodName             = "/bufman.dubbo.apache.org.registry.v1alpha1.DisplayService/DisplayServerElements"
	DisplayService_DisplayOwnerEntitledElements_FullMethodName      = "/bufman.dubbo.apache.org.registry.v1alpha1.DisplayService/DisplayOwnerEntitledElements"
	DisplayService_DisplayRepositoryEntitledElements_FullMethodName = "/bufman.dubbo.apache.org.registry.v1alpha1.DisplayService/DisplayRepositoryEntitledElements"
	DisplayService_ListManageableRepositoryRoles_FullMethodName     = "/bufman.dubbo.apache.org.registry.v1alpha1.DisplayService/ListManageableRepositoryRoles"
	DisplayService_ListManageableUserRepositoryRoles_FullMethodName = "/bufman.dubbo.apache.org.registry.v1alpha1.DisplayService/ListManageableUserRepositoryRoles"
)

// DisplayServiceClient is the client API for DisplayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DisplayServiceClient interface {
	// DisplayOrganizationElements returns which organization elements should be displayed to the user.
	DisplayOrganizationElements(ctx context.Context, in *DisplayOrganizationElementsRequest, opts ...grpc.CallOption) (*DisplayOrganizationElementsResponse, error)
	// DisplayRepositoryElements returns which repository elements should be displayed to the user.
	DisplayRepositoryElements(ctx context.Context, in *DisplayRepositoryElementsRequest, opts ...grpc.CallOption) (*DisplayRepositoryElementsResponse, error)
	// DisplayUserElements returns which user elements should be displayed to the user.
	DisplayUserElements(ctx context.Context, in *DisplayUserElementsRequest, opts ...grpc.CallOption) (*DisplayUserElementsResponse, error)
	// DisplayServerElements returns which server elements should be displayed to the user.
	DisplayServerElements(ctx context.Context, in *DisplayServerElementsRequest, opts ...grpc.CallOption) (*DisplayServerElementsResponse, error)
	// DisplayOwnerEntitledElements returns which owner elements are entitled to be displayed to the user.
	DisplayOwnerEntitledElements(ctx context.Context, in *DisplayOwnerEntitledElementsRequest, opts ...grpc.CallOption) (*DisplayOwnerEntitledElementsResponse, error)
	// DisplayRepositoryEntitledElements returns which repository elements are entitled to be displayed to the user.
	DisplayRepositoryEntitledElements(ctx context.Context, in *DisplayRepositoryEntitledElementsRequest, opts ...grpc.CallOption) (*DisplayRepositoryEntitledElementsResponse, error)
	// ListManageableRepositoryRoles returns which roles should be displayed
	// to the user when they are managing contributors on the repository.
	ListManageableRepositoryRoles(ctx context.Context, in *ListManageableRepositoryRolesRequest, opts ...grpc.CallOption) (*ListManageableRepositoryRolesResponse, error)
	// ListManageableUserRepositoryRoles returns which roles should be displayed
	// to the user when they are managing a specific contributor on the repository.
	ListManageableUserRepositoryRoles(ctx context.Context, in *ListManageableUserRepositoryRolesRequest, opts ...grpc.CallOption) (*ListManageableUserRepositoryRolesResponse, error)
}

type displayServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDisplayServiceClient(cc grpc.ClientConnInterface) DisplayServiceClient {
	return &displayServiceClient{cc}
}

func (c *displayServiceClient) DisplayOrganizationElements(ctx context.Context, in *DisplayOrganizationElementsRequest, opts ...grpc.CallOption) (*DisplayOrganizationElementsResponse, error) {
	out := new(DisplayOrganizationElementsResponse)
	err := c.cc.Invoke(ctx, DisplayService_DisplayOrganizationElements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayServiceClient) DisplayRepositoryElements(ctx context.Context, in *DisplayRepositoryElementsRequest, opts ...grpc.CallOption) (*DisplayRepositoryElementsResponse, error) {
	out := new(DisplayRepositoryElementsResponse)
	err := c.cc.Invoke(ctx, DisplayService_DisplayRepositoryElements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayServiceClient) DisplayUserElements(ctx context.Context, in *DisplayUserElementsRequest, opts ...grpc.CallOption) (*DisplayUserElementsResponse, error) {
	out := new(DisplayUserElementsResponse)
	err := c.cc.Invoke(ctx, DisplayService_DisplayUserElements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayServiceClient) DisplayServerElements(ctx context.Context, in *DisplayServerElementsRequest, opts ...grpc.CallOption) (*DisplayServerElementsResponse, error) {
	out := new(DisplayServerElementsResponse)
	err := c.cc.Invoke(ctx, DisplayService_DisplayServerElements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayServiceClient) DisplayOwnerEntitledElements(ctx context.Context, in *DisplayOwnerEntitledElementsRequest, opts ...grpc.CallOption) (*DisplayOwnerEntitledElementsResponse, error) {
	out := new(DisplayOwnerEntitledElementsResponse)
	err := c.cc.Invoke(ctx, DisplayService_DisplayOwnerEntitledElements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayServiceClient) DisplayRepositoryEntitledElements(ctx context.Context, in *DisplayRepositoryEntitledElementsRequest, opts ...grpc.CallOption) (*DisplayRepositoryEntitledElementsResponse, error) {
	out := new(DisplayRepositoryEntitledElementsResponse)
	err := c.cc.Invoke(ctx, DisplayService_DisplayRepositoryEntitledElements_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayServiceClient) ListManageableRepositoryRoles(ctx context.Context, in *ListManageableRepositoryRolesRequest, opts ...grpc.CallOption) (*ListManageableRepositoryRolesResponse, error) {
	out := new(ListManageableRepositoryRolesResponse)
	err := c.cc.Invoke(ctx, DisplayService_ListManageableRepositoryRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *displayServiceClient) ListManageableUserRepositoryRoles(ctx context.Context, in *ListManageableUserRepositoryRolesRequest, opts ...grpc.CallOption) (*ListManageableUserRepositoryRolesResponse, error) {
	out := new(ListManageableUserRepositoryRolesResponse)
	err := c.cc.Invoke(ctx, DisplayService_ListManageableUserRepositoryRoles_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DisplayServiceServer is the server API for DisplayService service.
// All implementations must embed UnimplementedDisplayServiceServer
// for forward compatibility
type DisplayServiceServer interface {
	// DisplayOrganizationElements returns which organization elements should be displayed to the user.
	DisplayOrganizationElements(context.Context, *DisplayOrganizationElementsRequest) (*DisplayOrganizationElementsResponse, error)
	// DisplayRepositoryElements returns which repository elements should be displayed to the user.
	DisplayRepositoryElements(context.Context, *DisplayRepositoryElementsRequest) (*DisplayRepositoryElementsResponse, error)
	// DisplayUserElements returns which user elements should be displayed to the user.
	DisplayUserElements(context.Context, *DisplayUserElementsRequest) (*DisplayUserElementsResponse, error)
	// DisplayServerElements returns which server elements should be displayed to the user.
	DisplayServerElements(context.Context, *DisplayServerElementsRequest) (*DisplayServerElementsResponse, error)
	// DisplayOwnerEntitledElements returns which owner elements are entitled to be displayed to the user.
	DisplayOwnerEntitledElements(context.Context, *DisplayOwnerEntitledElementsRequest) (*DisplayOwnerEntitledElementsResponse, error)
	// DisplayRepositoryEntitledElements returns which repository elements are entitled to be displayed to the user.
	DisplayRepositoryEntitledElements(context.Context, *DisplayRepositoryEntitledElementsRequest) (*DisplayRepositoryEntitledElementsResponse, error)
	// ListManageableRepositoryRoles returns which roles should be displayed
	// to the user when they are managing contributors on the repository.
	ListManageableRepositoryRoles(context.Context, *ListManageableRepositoryRolesRequest) (*ListManageableRepositoryRolesResponse, error)
	// ListManageableUserRepositoryRoles returns which roles should be displayed
	// to the user when they are managing a specific contributor on the repository.
	ListManageableUserRepositoryRoles(context.Context, *ListManageableUserRepositoryRolesRequest) (*ListManageableUserRepositoryRolesResponse, error)
	mustEmbedUnimplementedDisplayServiceServer()
}

// UnimplementedDisplayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDisplayServiceServer struct {
}

func (UnimplementedDisplayServiceServer) DisplayOrganizationElements(context.Context, *DisplayOrganizationElementsRequest) (*DisplayOrganizationElementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayOrganizationElements not implemented")
}
func (UnimplementedDisplayServiceServer) DisplayRepositoryElements(context.Context, *DisplayRepositoryElementsRequest) (*DisplayRepositoryElementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayRepositoryElements not implemented")
}
func (UnimplementedDisplayServiceServer) DisplayUserElements(context.Context, *DisplayUserElementsRequest) (*DisplayUserElementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayUserElements not implemented")
}
func (UnimplementedDisplayServiceServer) DisplayServerElements(context.Context, *DisplayServerElementsRequest) (*DisplayServerElementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayServerElements not implemented")
}
func (UnimplementedDisplayServiceServer) DisplayOwnerEntitledElements(context.Context, *DisplayOwnerEntitledElementsRequest) (*DisplayOwnerEntitledElementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayOwnerEntitledElements not implemented")
}
func (UnimplementedDisplayServiceServer) DisplayRepositoryEntitledElements(context.Context, *DisplayRepositoryEntitledElementsRequest) (*DisplayRepositoryEntitledElementsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisplayRepositoryEntitledElements not implemented")
}
func (UnimplementedDisplayServiceServer) ListManageableRepositoryRoles(context.Context, *ListManageableRepositoryRolesRequest) (*ListManageableRepositoryRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListManageableRepositoryRoles not implemented")
}
func (UnimplementedDisplayServiceServer) ListManageableUserRepositoryRoles(context.Context, *ListManageableUserRepositoryRolesRequest) (*ListManageableUserRepositoryRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListManageableUserRepositoryRoles not implemented")
}
func (UnimplementedDisplayServiceServer) mustEmbedUnimplementedDisplayServiceServer() {}

// UnsafeDisplayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DisplayServiceServer will
// result in compilation errors.
type UnsafeDisplayServiceServer interface {
	mustEmbedUnimplementedDisplayServiceServer()
}

func RegisterDisplayServiceServer(s grpc.ServiceRegistrar, srv DisplayServiceServer) {
	s.RegisterService(&DisplayService_ServiceDesc, srv)
}

func _DisplayService_DisplayOrganizationElements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayOrganizationElementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).DisplayOrganizationElements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisplayService_DisplayOrganizationElements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).DisplayOrganizationElements(ctx, req.(*DisplayOrganizationElementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisplayService_DisplayRepositoryElements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayRepositoryElementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).DisplayRepositoryElements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisplayService_DisplayRepositoryElements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).DisplayRepositoryElements(ctx, req.(*DisplayRepositoryElementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisplayService_DisplayUserElements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayUserElementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).DisplayUserElements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisplayService_DisplayUserElements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).DisplayUserElements(ctx, req.(*DisplayUserElementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisplayService_DisplayServerElements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayServerElementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).DisplayServerElements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisplayService_DisplayServerElements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).DisplayServerElements(ctx, req.(*DisplayServerElementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisplayService_DisplayOwnerEntitledElements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayOwnerEntitledElementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).DisplayOwnerEntitledElements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisplayService_DisplayOwnerEntitledElements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).DisplayOwnerEntitledElements(ctx, req.(*DisplayOwnerEntitledElementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisplayService_DisplayRepositoryEntitledElements_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DisplayRepositoryEntitledElementsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).DisplayRepositoryEntitledElements(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisplayService_DisplayRepositoryEntitledElements_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).DisplayRepositoryEntitledElements(ctx, req.(*DisplayRepositoryEntitledElementsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisplayService_ListManageableRepositoryRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListManageableRepositoryRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).ListManageableRepositoryRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisplayService_ListManageableRepositoryRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).ListManageableRepositoryRoles(ctx, req.(*ListManageableRepositoryRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DisplayService_ListManageableUserRepositoryRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListManageableUserRepositoryRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DisplayServiceServer).ListManageableUserRepositoryRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DisplayService_ListManageableUserRepositoryRoles_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DisplayServiceServer).ListManageableUserRepositoryRoles(ctx, req.(*ListManageableUserRepositoryRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DisplayService_ServiceDesc is the grpc.ServiceDesc for DisplayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DisplayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bufman.dubbo.apache.org.registry.v1alpha1.DisplayService",
	HandlerType: (*DisplayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DisplayOrganizationElements",
			Handler:    _DisplayService_DisplayOrganizationElements_Handler,
		},
		{
			MethodName: "DisplayRepositoryElements",
			Handler:    _DisplayService_DisplayRepositoryElements_Handler,
		},
		{
			MethodName: "DisplayUserElements",
			Handler:    _DisplayService_DisplayUserElements_Handler,
		},
		{
			MethodName: "DisplayServerElements",
			Handler:    _DisplayService_DisplayServerElements_Handler,
		},
		{
			MethodName: "DisplayOwnerEntitledElements",
			Handler:    _DisplayService_DisplayOwnerEntitledElements_Handler,
		},
		{
			MethodName: "DisplayRepositoryEntitledElements",
			Handler:    _DisplayService_DisplayRepositoryEntitledElements_Handler,
		},
		{
			MethodName: "ListManageableRepositoryRoles",
			Handler:    _DisplayService_ListManageableRepositoryRoles_Handler,
		},
		{
			MethodName: "ListManageableUserRepositoryRoles",
			Handler:    _DisplayService_ListManageableUserRepositoryRoles_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "registry/v1alpha1/display.proto",
}
