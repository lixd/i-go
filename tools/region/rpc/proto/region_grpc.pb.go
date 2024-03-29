// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package proto

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// RegionServerClient is the client API for RegionServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RegionServerClient interface {
	IP2Region(ctx context.Context, in *IP, opts ...grpc.CallOption) (*Region, error)
	IP2LatLong(ctx context.Context, in *IP, opts ...grpc.CallOption) (*LatLong, error)
}

type regionServerClient struct {
	cc grpc.ClientConnInterface
}

func NewRegionServerClient(cc grpc.ClientConnInterface) RegionServerClient {
	return &regionServerClient{cc}
}

func (c *regionServerClient) IP2Region(ctx context.Context, in *IP, opts ...grpc.CallOption) (*Region, error) {
	out := new(Region)
	err := c.cc.Invoke(ctx, "/proto.RegionServer/IP2Region", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *regionServerClient) IP2LatLong(ctx context.Context, in *IP, opts ...grpc.CallOption) (*LatLong, error) {
	out := new(LatLong)
	err := c.cc.Invoke(ctx, "/proto.RegionServer/IP2LatLong", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RegionServerServer is the server API for RegionServer service.
// All implementations must embed UnimplementedRegionServerServer
// for forward compatibility
type RegionServerServer interface {
	IP2Region(context.Context, *IP) (*Region, error)
	IP2LatLong(context.Context, *IP) (*LatLong, error)
	mustEmbedUnimplementedRegionServerServer()
}

// UnimplementedRegionServerServer must be embedded to have forward compatible implementations.
type UnimplementedRegionServerServer struct {
}

func (UnimplementedRegionServerServer) IP2Region(context.Context, *IP) (*Region, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IP2Region not implemented")
}
func (UnimplementedRegionServerServer) IP2LatLong(context.Context, *IP) (*LatLong, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IP2LatLong not implemented")
}
func (UnimplementedRegionServerServer) mustEmbedUnimplementedRegionServerServer() {}

// UnsafeRegionServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RegionServerServer will
// result in compilation errors.
type UnsafeRegionServerServer interface {
	mustEmbedUnimplementedRegionServerServer()
}

func RegisterRegionServerServer(s grpc.ServiceRegistrar, srv RegionServerServer) {
	s.RegisterService(&_RegionServer_serviceDesc, srv)
}

func _RegionServer_IP2Region_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServerServer).IP2Region(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RegionServer/IP2Region",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServerServer).IP2Region(ctx, req.(*IP))
	}
	return interceptor(ctx, in, info, handler)
}

func _RegionServer_IP2LatLong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IP)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RegionServerServer).IP2LatLong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.RegionServer/IP2LatLong",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RegionServerServer).IP2LatLong(ctx, req.(*IP))
	}
	return interceptor(ctx, in, info, handler)
}

var _RegionServer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.RegionServer",
	HandlerType: (*RegionServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IP2Region",
			Handler:    _RegionServer_IP2Region_Handler,
		},
		{
			MethodName: "IP2LatLong",
			Handler:    _RegionServer_IP2LatLong_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "region.proto",
}
