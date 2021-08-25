// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// CommandServiceClient is the client API for CommandService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandServiceClient interface {
	Execute(ctx context.Context, opts ...grpc.CallOption) (CommandService_ExecuteClient, error)
}

type commandServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommandServiceClient(cc grpc.ClientConnInterface) CommandServiceClient {
	return &commandServiceClient{cc}
}

func (c *commandServiceClient) Execute(ctx context.Context, opts ...grpc.CallOption) (CommandService_ExecuteClient, error) {
	stream, err := c.cc.NewStream(ctx, &CommandService_ServiceDesc.Streams[0], "/CommandService/Execute", opts...)
	if err != nil {
		return nil, err
	}
	x := &commandServiceExecuteClient{stream}
	return x, nil
}

type CommandService_ExecuteClient interface {
	Send(*ServerMsg) error
	Recv() (*ClientMsg, error)
	grpc.ClientStream
}

type commandServiceExecuteClient struct {
	grpc.ClientStream
}

func (x *commandServiceExecuteClient) Send(m *ServerMsg) error {
	return x.ClientStream.SendMsg(m)
}

func (x *commandServiceExecuteClient) Recv() (*ClientMsg, error) {
	m := new(ClientMsg)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommandServiceServer is the server API for CommandService service.
// All implementations must embed UnimplementedCommandServiceServer
// for forward compatibility
type CommandServiceServer interface {
	Execute(CommandService_ExecuteServer) error
	mustEmbedUnimplementedCommandServiceServer()
}

// UnimplementedCommandServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommandServiceServer struct {
}

func (UnimplementedCommandServiceServer) Execute(CommandService_ExecuteServer) error {
	return status.Errorf(codes.Unimplemented, "method Execute not implemented")
}
func (UnimplementedCommandServiceServer) mustEmbedUnimplementedCommandServiceServer() {}

// UnsafeCommandServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandServiceServer will
// result in compilation errors.
type UnsafeCommandServiceServer interface {
	mustEmbedUnimplementedCommandServiceServer()
}

func RegisterCommandServiceServer(s grpc.ServiceRegistrar, srv CommandServiceServer) {
	s.RegisterService(&CommandService_ServiceDesc, srv)
}

func _CommandService_Execute_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CommandServiceServer).Execute(&commandServiceExecuteServer{stream})
}

type CommandService_ExecuteServer interface {
	Send(*ClientMsg) error
	Recv() (*ServerMsg, error)
	grpc.ServerStream
}

type commandServiceExecuteServer struct {
	grpc.ServerStream
}

func (x *commandServiceExecuteServer) Send(m *ClientMsg) error {
	return x.ServerStream.SendMsg(m)
}

func (x *commandServiceExecuteServer) Recv() (*ServerMsg, error) {
	m := new(ServerMsg)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommandService_ServiceDesc is the grpc.ServiceDesc for CommandService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CommandService",
	HandlerType: (*CommandServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Execute",
			Handler:       _CommandService_Execute_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "test.proto",
}
