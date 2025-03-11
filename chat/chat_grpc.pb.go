// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: chat.proto

package chat

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChatMessageService_Chat_FullMethodName = "/chat.ChatMessageService/Chat"
)

// ChatMessageServiceClient is the client API for ChatMessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatMessageServiceClient interface {
	Chat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ClientMessage, ServerResponse], error)
}

type chatMessageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatMessageServiceClient(cc grpc.ClientConnInterface) ChatMessageServiceClient {
	return &chatMessageServiceClient{cc}
}

func (c *chatMessageServiceClient) Chat(ctx context.Context, opts ...grpc.CallOption) (grpc.BidiStreamingClient[ClientMessage, ServerResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ChatMessageService_ServiceDesc.Streams[0], ChatMessageService_Chat_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ClientMessage, ServerResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatMessageService_ChatClient = grpc.BidiStreamingClient[ClientMessage, ServerResponse]

// ChatMessageServiceServer is the server API for ChatMessageService service.
// All implementations must embed UnimplementedChatMessageServiceServer
// for forward compatibility.
type ChatMessageServiceServer interface {
	Chat(grpc.BidiStreamingServer[ClientMessage, ServerResponse]) error
	mustEmbedUnimplementedChatMessageServiceServer()
}

// UnimplementedChatMessageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChatMessageServiceServer struct{}

func (UnimplementedChatMessageServiceServer) Chat(grpc.BidiStreamingServer[ClientMessage, ServerResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
}
func (UnimplementedChatMessageServiceServer) mustEmbedUnimplementedChatMessageServiceServer() {}
func (UnimplementedChatMessageServiceServer) testEmbeddedByValue()                            {}

// UnsafeChatMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatMessageServiceServer will
// result in compilation errors.
type UnsafeChatMessageServiceServer interface {
	mustEmbedUnimplementedChatMessageServiceServer()
}

func RegisterChatMessageServiceServer(s grpc.ServiceRegistrar, srv ChatMessageServiceServer) {
	// If the following call pancis, it indicates UnimplementedChatMessageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChatMessageService_ServiceDesc, srv)
}

func _ChatMessageService_Chat_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatMessageServiceServer).Chat(&grpc.GenericServerStream[ClientMessage, ServerResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ChatMessageService_ChatServer = grpc.BidiStreamingServer[ClientMessage, ServerResponse]

// ChatMessageService_ServiceDesc is the grpc.ServiceDesc for ChatMessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatMessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chat.ChatMessageService",
	HandlerType: (*ChatMessageServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Chat",
			Handler:       _ChatMessageService_Chat_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat.proto",
}
