// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: emailmessages.proto

package emailmessages

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
	EmailMessageService_SendEmail_FullMethodName       = "/emailmessages.EmailMessageService/SendEmail"
	EmailMessageService_BroadcastLetter_FullMethodName = "/emailmessages.EmailMessageService/BroadcastLetter"
	EmailMessageService_UpdateScribers_FullMethodName  = "/emailmessages.EmailMessageService/UpdateScribers"
)

// EmailMessageServiceClient is the client API for EmailMessageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmailMessageServiceClient interface {
	SendEmail(ctx context.Context, in *EmailReq, opts ...grpc.CallOption) (*GenericResponse, error)
	BroadcastLetter(ctx context.Context, in *LetterMessage, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Broadcast], error)
	UpdateScribers(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SubscriberList, GenericResponse], error)
}

type emailMessageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmailMessageServiceClient(cc grpc.ClientConnInterface) EmailMessageServiceClient {
	return &emailMessageServiceClient{cc}
}

func (c *emailMessageServiceClient) SendEmail(ctx context.Context, in *EmailReq, opts ...grpc.CallOption) (*GenericResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, EmailMessageService_SendEmail_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emailMessageServiceClient) BroadcastLetter(ctx context.Context, in *LetterMessage, opts ...grpc.CallOption) (grpc.ServerStreamingClient[Broadcast], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &EmailMessageService_ServiceDesc.Streams[0], EmailMessageService_BroadcastLetter_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[LetterMessage, Broadcast]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EmailMessageService_BroadcastLetterClient = grpc.ServerStreamingClient[Broadcast]

func (c *emailMessageServiceClient) UpdateScribers(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[SubscriberList, GenericResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &EmailMessageService_ServiceDesc.Streams[1], EmailMessageService_UpdateScribers_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[SubscriberList, GenericResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EmailMessageService_UpdateScribersClient = grpc.ClientStreamingClient[SubscriberList, GenericResponse]

// EmailMessageServiceServer is the server API for EmailMessageService service.
// All implementations must embed UnimplementedEmailMessageServiceServer
// for forward compatibility.
type EmailMessageServiceServer interface {
	SendEmail(context.Context, *EmailReq) (*GenericResponse, error)
	BroadcastLetter(*LetterMessage, grpc.ServerStreamingServer[Broadcast]) error
	UpdateScribers(grpc.ClientStreamingServer[SubscriberList, GenericResponse]) error
	mustEmbedUnimplementedEmailMessageServiceServer()
}

// UnimplementedEmailMessageServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedEmailMessageServiceServer struct{}

func (UnimplementedEmailMessageServiceServer) SendEmail(context.Context, *EmailReq) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedEmailMessageServiceServer) BroadcastLetter(*LetterMessage, grpc.ServerStreamingServer[Broadcast]) error {
	return status.Errorf(codes.Unimplemented, "method BroadcastLetter not implemented")
}
func (UnimplementedEmailMessageServiceServer) UpdateScribers(grpc.ClientStreamingServer[SubscriberList, GenericResponse]) error {
	return status.Errorf(codes.Unimplemented, "method UpdateScribers not implemented")
}
func (UnimplementedEmailMessageServiceServer) mustEmbedUnimplementedEmailMessageServiceServer() {}
func (UnimplementedEmailMessageServiceServer) testEmbeddedByValue()                             {}

// UnsafeEmailMessageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmailMessageServiceServer will
// result in compilation errors.
type UnsafeEmailMessageServiceServer interface {
	mustEmbedUnimplementedEmailMessageServiceServer()
}

func RegisterEmailMessageServiceServer(s grpc.ServiceRegistrar, srv EmailMessageServiceServer) {
	// If the following call pancis, it indicates UnimplementedEmailMessageServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&EmailMessageService_ServiceDesc, srv)
}

func _EmailMessageService_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmailMessageServiceServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EmailMessageService_SendEmail_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmailMessageServiceServer).SendEmail(ctx, req.(*EmailReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmailMessageService_BroadcastLetter_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LetterMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EmailMessageServiceServer).BroadcastLetter(m, &grpc.GenericServerStream[LetterMessage, Broadcast]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EmailMessageService_BroadcastLetterServer = grpc.ServerStreamingServer[Broadcast]

func _EmailMessageService_UpdateScribers_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EmailMessageServiceServer).UpdateScribers(&grpc.GenericServerStream[SubscriberList, GenericResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type EmailMessageService_UpdateScribersServer = grpc.ClientStreamingServer[SubscriberList, GenericResponse]

// EmailMessageService_ServiceDesc is the grpc.ServiceDesc for EmailMessageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmailMessageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "emailmessages.EmailMessageService",
	HandlerType: (*EmailMessageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendEmail",
			Handler:    _EmailMessageService_SendEmail_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "BroadcastLetter",
			Handler:       _EmailMessageService_BroadcastLetter_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdateScribers",
			Handler:       _EmailMessageService_UpdateScribers_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "emailmessages.proto",
}
