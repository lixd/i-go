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

// ServerStreamClient is the client API for ServerStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerStreamClient interface {
	// 客户端传入一个数,服务端分别返回该数的0到9次方
	Pow(ctx context.Context, in *ServerStreamReq, opts ...grpc.CallOption) (ServerStream_PowClient, error)
}

type serverStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewServerStreamClient(cc grpc.ClientConnInterface) ServerStreamClient {
	return &serverStreamClient{cc}
}

func (c *serverStreamClient) Pow(ctx context.Context, in *ServerStreamReq, opts ...grpc.CallOption) (ServerStream_PowClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ServerStream_serviceDesc.Streams[0], "/stream.ServerStream/Pow", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverStreamPowClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServerStream_PowClient interface {
	Recv() (*ServerStreamResp, error)
	grpc.ClientStream
}

type serverStreamPowClient struct {
	grpc.ClientStream
}

func (x *serverStreamPowClient) Recv() (*ServerStreamResp, error) {
	m := new(ServerStreamResp)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerStreamServer is the server API for ServerStream service.
// All implementations must embed UnimplementedServerStreamServer
// for forward compatibility
type ServerStreamServer interface {
	// 客户端传入一个数,服务端分别返回该数的0到9次方
	Pow(*ServerStreamReq, ServerStream_PowServer) error
	mustEmbedUnimplementedServerStreamServer()
}

// UnimplementedServerStreamServer must be embedded to have forward compatible implementations.
type UnimplementedServerStreamServer struct {
}

func (UnimplementedServerStreamServer) Pow(*ServerStreamReq, ServerStream_PowServer) error {
	return status.Errorf(codes.Unimplemented, "method Pow not implemented")
}
func (UnimplementedServerStreamServer) mustEmbedUnimplementedServerStreamServer() {}

// UnsafeServerStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerStreamServer will
// result in compilation errors.
type UnsafeServerStreamServer interface {
	mustEmbedUnimplementedServerStreamServer()
}

func RegisterServerStreamServer(s grpc.ServiceRegistrar, srv ServerStreamServer) {
	s.RegisterService(&_ServerStream_serviceDesc, srv)
}

func _ServerStream_Pow_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ServerStreamReq)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerStreamServer).Pow(m, &serverStreamPowServer{stream})
}

type ServerStream_PowServer interface {
	Send(*ServerStreamResp) error
	grpc.ServerStream
}

type serverStreamPowServer struct {
	grpc.ServerStream
}

func (x *serverStreamPowServer) Send(m *ServerStreamResp) error {
	return x.ServerStream.SendMsg(m)
}

var _ServerStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "stream.ServerStream",
	HandlerType: (*ServerStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Pow",
			Handler:       _ServerStream_Pow_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server_stream.proto",
}
