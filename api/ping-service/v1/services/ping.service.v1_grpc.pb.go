// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.6
// source: api/ping-service/v1/services/ping.service.v1.proto

package servicev1

import (
	context "context"
	resources "github.com/go-micro-saas/service-api/api/ping-service/v1/resources"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	SrvPing_Ping_FullMethodName = "/saas.api.ping.servicev1.SrvPing/Ping"
)

// SrvPingClient is the client API for SrvPing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SrvPingClient interface {
	// Ping ping
	//
	// 测试PingV1
	Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error)
}

type srvPingClient struct {
	cc grpc.ClientConnInterface
}

func NewSrvPingClient(cc grpc.ClientConnInterface) SrvPingClient {
	return &srvPingClient{cc}
}

func (c *srvPingClient) Ping(ctx context.Context, in *resources.PingReq, opts ...grpc.CallOption) (*resources.PingResp, error) {
	out := new(resources.PingResp)
	err := c.cc.Invoke(ctx, SrvPing_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SrvPingServer is the server API for SrvPing service.
// All implementations must embed UnimplementedSrvPingServer
// for forward compatibility
type SrvPingServer interface {
	// Ping ping
	//
	// 测试PingV1
	Ping(context.Context, *resources.PingReq) (*resources.PingResp, error)
	mustEmbedUnimplementedSrvPingServer()
}

// UnimplementedSrvPingServer must be embedded to have forward compatible implementations.
type UnimplementedSrvPingServer struct {
}

func (UnimplementedSrvPingServer) Ping(context.Context, *resources.PingReq) (*resources.PingResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedSrvPingServer) mustEmbedUnimplementedSrvPingServer() {}

// UnsafeSrvPingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SrvPingServer will
// result in compilation errors.
type UnsafeSrvPingServer interface {
	mustEmbedUnimplementedSrvPingServer()
}

func RegisterSrvPingServer(s grpc.ServiceRegistrar, srv SrvPingServer) {
	s.RegisterService(&SrvPing_ServiceDesc, srv)
}

func _SrvPing_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(resources.PingReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SrvPingServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SrvPing_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SrvPingServer).Ping(ctx, req.(*resources.PingReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SrvPing_ServiceDesc is the grpc.ServiceDesc for SrvPing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SrvPing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "saas.api.ping.servicev1.SrvPing",
	HandlerType: (*SrvPingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SrvPing_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/ping-service/v1/services/ping.service.v1.proto",
}
