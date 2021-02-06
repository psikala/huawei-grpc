// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package huawei_dialin

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

// GRPCConfigOperClient is the client API for GRPCConfigOper service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GRPCConfigOperClient interface {
	Subscribe(ctx context.Context, in *SubsArgs, opts ...grpc.CallOption) (GRPCConfigOper_SubscribeClient, error)
	Cancel(ctx context.Context, in *CancelArgs, opts ...grpc.CallOption) (*CancelReply, error)
}

type gRPCConfigOperClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCConfigOperClient(cc grpc.ClientConnInterface) GRPCConfigOperClient {
	return &gRPCConfigOperClient{cc}
}

func (c *gRPCConfigOperClient) Subscribe(ctx context.Context, in *SubsArgs, opts ...grpc.CallOption) (GRPCConfigOper_SubscribeClient, error) {
	stream, err := c.cc.NewStream(ctx, &GRPCConfigOper_ServiceDesc.Streams[0], "/huawei_dialin.gRPCConfigOper/Subscribe", opts...)
	if err != nil {
		return nil, err
	}
	x := &gRPCConfigOperSubscribeClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type GRPCConfigOper_SubscribeClient interface {
	Recv() (*SubsReply, error)
	grpc.ClientStream
}

type gRPCConfigOperSubscribeClient struct {
	grpc.ClientStream
}

func (x *gRPCConfigOperSubscribeClient) Recv() (*SubsReply, error) {
	m := new(SubsReply)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gRPCConfigOperClient) Cancel(ctx context.Context, in *CancelArgs, opts ...grpc.CallOption) (*CancelReply, error) {
	out := new(CancelReply)
	err := c.cc.Invoke(ctx, "/huawei_dialin.gRPCConfigOper/Cancel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCConfigOperServer is the server API for GRPCConfigOper service.
// All implementations must embed UnimplementedGRPCConfigOperServer
// for forward compatibility
type GRPCConfigOperServer interface {
	Subscribe(*SubsArgs, GRPCConfigOper_SubscribeServer) error
	Cancel(context.Context, *CancelArgs) (*CancelReply, error)
	mustEmbedUnimplementedGRPCConfigOperServer()
}

// UnimplementedGRPCConfigOperServer must be embedded to have forward compatible implementations.
type UnimplementedGRPCConfigOperServer struct {
}

func (UnimplementedGRPCConfigOperServer) Subscribe(*SubsArgs, GRPCConfigOper_SubscribeServer) error {
	return status.Errorf(codes.Unimplemented, "method Subscribe not implemented")
}
func (UnimplementedGRPCConfigOperServer) Cancel(context.Context, *CancelArgs) (*CancelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Cancel not implemented")
}
func (UnimplementedGRPCConfigOperServer) mustEmbedUnimplementedGRPCConfigOperServer() {}

// UnsafeGRPCConfigOperServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCConfigOperServer will
// result in compilation errors.
type UnsafeGRPCConfigOperServer interface {
	mustEmbedUnimplementedGRPCConfigOperServer()
}

func RegisterGRPCConfigOperServer(s grpc.ServiceRegistrar, srv GRPCConfigOperServer) {
	s.RegisterService(&GRPCConfigOper_ServiceDesc, srv)
}

func _GRPCConfigOper_Subscribe_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubsArgs)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(GRPCConfigOperServer).Subscribe(m, &gRPCConfigOperSubscribeServer{stream})
}

type GRPCConfigOper_SubscribeServer interface {
	Send(*SubsReply) error
	grpc.ServerStream
}

type gRPCConfigOperSubscribeServer struct {
	grpc.ServerStream
}

func (x *gRPCConfigOperSubscribeServer) Send(m *SubsReply) error {
	return x.ServerStream.SendMsg(m)
}

func _GRPCConfigOper_Cancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelArgs)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCConfigOperServer).Cancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/huawei_dialin.gRPCConfigOper/Cancel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCConfigOperServer).Cancel(ctx, req.(*CancelArgs))
	}
	return interceptor(ctx, in, info, handler)
}

// GRPCConfigOper_ServiceDesc is the grpc.ServiceDesc for GRPCConfigOper service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRPCConfigOper_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "huawei_dialin.gRPCConfigOper",
	HandlerType: (*GRPCConfigOperServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Cancel",
			Handler:    _GRPCConfigOper_Cancel_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Subscribe",
			Handler:       _GRPCConfigOper_Subscribe_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "proto/huawei-dialin/huawei-grpc-dialin.proto",
}