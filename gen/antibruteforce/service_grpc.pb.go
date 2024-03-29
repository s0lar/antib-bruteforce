// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/service.proto

package antibruteforce

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

// CheckerClient is the client API for Checker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CheckerClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
	Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error)
	AddBlacklist(ctx context.Context, in *NetListRequest, opts ...grpc.CallOption) (*NetListResponse, error)
	RemoveBlacklist(ctx context.Context, in *NetListRequest, opts ...grpc.CallOption) (*NetListResponse, error)
	AddWhitelist(ctx context.Context, in *NetListRequest, opts ...grpc.CallOption) (*NetListResponse, error)
	RemoveWhitelist(ctx context.Context, in *NetListRequest, opts ...grpc.CallOption) (*NetListResponse, error)
}

type checkerClient struct {
	cc grpc.ClientConnInterface
}

func NewCheckerClient(cc grpc.ClientConnInterface) CheckerClient {
	return &checkerClient{cc}
}

func (c *checkerClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/antibruteforce.Checker/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) Reset(ctx context.Context, in *ResetRequest, opts ...grpc.CallOption) (*ResetResponse, error) {
	out := new(ResetResponse)
	err := c.cc.Invoke(ctx, "/antibruteforce.Checker/Reset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) AddBlacklist(ctx context.Context, in *NetListRequest, opts ...grpc.CallOption) (*NetListResponse, error) {
	out := new(NetListResponse)
	err := c.cc.Invoke(ctx, "/antibruteforce.Checker/AddBlacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) RemoveBlacklist(ctx context.Context, in *NetListRequest, opts ...grpc.CallOption) (*NetListResponse, error) {
	out := new(NetListResponse)
	err := c.cc.Invoke(ctx, "/antibruteforce.Checker/RemoveBlacklist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) AddWhitelist(ctx context.Context, in *NetListRequest, opts ...grpc.CallOption) (*NetListResponse, error) {
	out := new(NetListResponse)
	err := c.cc.Invoke(ctx, "/antibruteforce.Checker/AddWhitelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *checkerClient) RemoveWhitelist(ctx context.Context, in *NetListRequest, opts ...grpc.CallOption) (*NetListResponse, error) {
	out := new(NetListResponse)
	err := c.cc.Invoke(ctx, "/antibruteforce.Checker/RemoveWhitelist", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CheckerServer is the server API for Checker service.
// All implementations must embed UnimplementedCheckerServer
// for forward compatibility
type CheckerServer interface {
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
	Reset(context.Context, *ResetRequest) (*ResetResponse, error)
	AddBlacklist(context.Context, *NetListRequest) (*NetListResponse, error)
	RemoveBlacklist(context.Context, *NetListRequest) (*NetListResponse, error)
	AddWhitelist(context.Context, *NetListRequest) (*NetListResponse, error)
	RemoveWhitelist(context.Context, *NetListRequest) (*NetListResponse, error)
	mustEmbedUnimplementedCheckerServer()
}

// UnimplementedCheckerServer must be embedded to have forward compatible implementations.
type UnimplementedCheckerServer struct {
}

func (UnimplementedCheckerServer) Check(context.Context, *CheckRequest) (*CheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (UnimplementedCheckerServer) Reset(context.Context, *ResetRequest) (*ResetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reset not implemented")
}
func (UnimplementedCheckerServer) AddBlacklist(context.Context, *NetListRequest) (*NetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBlacklist not implemented")
}
func (UnimplementedCheckerServer) RemoveBlacklist(context.Context, *NetListRequest) (*NetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveBlacklist not implemented")
}
func (UnimplementedCheckerServer) AddWhitelist(context.Context, *NetListRequest) (*NetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddWhitelist not implemented")
}
func (UnimplementedCheckerServer) RemoveWhitelist(context.Context, *NetListRequest) (*NetListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveWhitelist not implemented")
}
func (UnimplementedCheckerServer) mustEmbedUnimplementedCheckerServer() {}

// UnsafeCheckerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CheckerServer will
// result in compilation errors.
type UnsafeCheckerServer interface {
	mustEmbedUnimplementedCheckerServer()
}

func RegisterCheckerServer(s grpc.ServiceRegistrar, srv CheckerServer) {
	s.RegisterService(&Checker_ServiceDesc, srv)
}

func _Checker_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/antibruteforce.Checker/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_Reset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).Reset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/antibruteforce.Checker/Reset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).Reset(ctx, req.(*ResetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_AddBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).AddBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/antibruteforce.Checker/AddBlacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).AddBlacklist(ctx, req.(*NetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_RemoveBlacklist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).RemoveBlacklist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/antibruteforce.Checker/RemoveBlacklist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).RemoveBlacklist(ctx, req.(*NetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_AddWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).AddWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/antibruteforce.Checker/AddWhitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).AddWhitelist(ctx, req.(*NetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Checker_RemoveWhitelist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NetListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CheckerServer).RemoveWhitelist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/antibruteforce.Checker/RemoveWhitelist",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CheckerServer).RemoveWhitelist(ctx, req.(*NetListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Checker_ServiceDesc is the grpc.ServiceDesc for Checker service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Checker_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "antibruteforce.Checker",
	HandlerType: (*CheckerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Checker_Check_Handler,
		},
		{
			MethodName: "Reset",
			Handler:    _Checker_Reset_Handler,
		},
		{
			MethodName: "AddBlacklist",
			Handler:    _Checker_AddBlacklist_Handler,
		},
		{
			MethodName: "RemoveBlacklist",
			Handler:    _Checker_RemoveBlacklist_Handler,
		},
		{
			MethodName: "AddWhitelist",
			Handler:    _Checker_AddWhitelist_Handler,
		},
		{
			MethodName: "RemoveWhitelist",
			Handler:    _Checker_RemoveWhitelist_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/service.proto",
}
