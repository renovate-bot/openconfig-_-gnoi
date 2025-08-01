// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: github.com/openconfig/gnoi/otdr/otdr.proto

package otdr

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
	OTDR_Initiate_FullMethodName = "/gnoi.optical.OTDR/Initiate"
)

// OTDRClient is the client API for OTDR service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OTDRClient interface {
	Initiate(ctx context.Context, in *InitiateRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[InitiateResponse], error)
}

type oTDRClient struct {
	cc grpc.ClientConnInterface
}

func NewOTDRClient(cc grpc.ClientConnInterface) OTDRClient {
	return &oTDRClient{cc}
}

func (c *oTDRClient) Initiate(ctx context.Context, in *InitiateRequest, opts ...grpc.CallOption) (grpc.ServerStreamingClient[InitiateResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &OTDR_ServiceDesc.Streams[0], OTDR_Initiate_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[InitiateRequest, InitiateResponse]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OTDR_InitiateClient = grpc.ServerStreamingClient[InitiateResponse]

// OTDRServer is the server API for OTDR service.
// All implementations should embed UnimplementedOTDRServer
// for forward compatibility.
type OTDRServer interface {
	Initiate(*InitiateRequest, grpc.ServerStreamingServer[InitiateResponse]) error
}

// UnimplementedOTDRServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOTDRServer struct{}

func (UnimplementedOTDRServer) Initiate(*InitiateRequest, grpc.ServerStreamingServer[InitiateResponse]) error {
	return status.Errorf(codes.Unimplemented, "method Initiate not implemented")
}
func (UnimplementedOTDRServer) testEmbeddedByValue() {}

// UnsafeOTDRServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OTDRServer will
// result in compilation errors.
type UnsafeOTDRServer interface {
	mustEmbedUnimplementedOTDRServer()
}

func RegisterOTDRServer(s grpc.ServiceRegistrar, srv OTDRServer) {
	// If the following call pancis, it indicates UnimplementedOTDRServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OTDR_ServiceDesc, srv)
}

func _OTDR_Initiate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(InitiateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(OTDRServer).Initiate(m, &grpc.GenericServerStream[InitiateRequest, InitiateResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type OTDR_InitiateServer = grpc.ServerStreamingServer[InitiateResponse]

// OTDR_ServiceDesc is the grpc.ServiceDesc for OTDR service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OTDR_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.optical.OTDR",
	HandlerType: (*OTDRServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Initiate",
			Handler:       _OTDR_Initiate_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "github.com/openconfig/gnoi/otdr/otdr.proto",
}
