// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.31.1
// source: github.com/openconfig/gnoi/packet_link_qualification/packet_link_qualification.proto

package linkqual

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
	LinkQualification_Create_FullMethodName       = "/gnoi.packet_link_qualification.LinkQualification/Create"
	LinkQualification_Get_FullMethodName          = "/gnoi.packet_link_qualification.LinkQualification/Get"
	LinkQualification_Capabilities_FullMethodName = "/gnoi.packet_link_qualification.LinkQualification/Capabilities"
	LinkQualification_Delete_FullMethodName       = "/gnoi.packet_link_qualification.LinkQualification/Delete"
	LinkQualification_List_FullMethodName         = "/gnoi.packet_link_qualification.LinkQualification/List"
)

// LinkQualificationClient is the client API for LinkQualification service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkQualificationClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	Capabilities(ctx context.Context, in *CapabilitiesRequest, opts ...grpc.CallOption) (*CapabilitiesResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
}

type linkQualificationClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkQualificationClient(cc grpc.ClientConnInterface) LinkQualificationClient {
	return &linkQualificationClient{cc}
}

func (c *linkQualificationClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, LinkQualification_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkQualificationClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, LinkQualification_Get_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkQualificationClient) Capabilities(ctx context.Context, in *CapabilitiesRequest, opts ...grpc.CallOption) (*CapabilitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CapabilitiesResponse)
	err := c.cc.Invoke(ctx, LinkQualification_Capabilities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkQualificationClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, LinkQualification_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkQualificationClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, LinkQualification_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkQualificationServer is the server API for LinkQualification service.
// All implementations should embed UnimplementedLinkQualificationServer
// for forward compatibility.
type LinkQualificationServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	Capabilities(context.Context, *CapabilitiesRequest) (*CapabilitiesResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

// UnimplementedLinkQualificationServer should be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedLinkQualificationServer struct{}

func (UnimplementedLinkQualificationServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedLinkQualificationServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedLinkQualificationServer) Capabilities(context.Context, *CapabilitiesRequest) (*CapabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Capabilities not implemented")
}
func (UnimplementedLinkQualificationServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedLinkQualificationServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedLinkQualificationServer) testEmbeddedByValue() {}

// UnsafeLinkQualificationServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkQualificationServer will
// result in compilation errors.
type UnsafeLinkQualificationServer interface {
	mustEmbedUnimplementedLinkQualificationServer()
}

func RegisterLinkQualificationServer(s grpc.ServiceRegistrar, srv LinkQualificationServer) {
	// If the following call pancis, it indicates UnimplementedLinkQualificationServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&LinkQualification_ServiceDesc, srv)
}

func _LinkQualification_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkQualificationServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkQualification_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkQualificationServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkQualification_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkQualificationServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkQualification_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkQualificationServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkQualification_Capabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkQualificationServer).Capabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkQualification_Capabilities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkQualificationServer).Capabilities(ctx, req.(*CapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkQualification_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkQualificationServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkQualification_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkQualificationServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LinkQualification_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkQualificationServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: LinkQualification_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkQualificationServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LinkQualification_ServiceDesc is the grpc.ServiceDesc for LinkQualification service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LinkQualification_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gnoi.packet_link_qualification.LinkQualification",
	HandlerType: (*LinkQualificationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _LinkQualification_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _LinkQualification_Get_Handler,
		},
		{
			MethodName: "Capabilities",
			Handler:    _LinkQualification_Capabilities_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _LinkQualification_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _LinkQualification_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openconfig/gnoi/packet_link_qualification/packet_link_qualification.proto",
}
