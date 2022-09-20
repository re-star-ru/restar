// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.14.0
// source: api/proto/v1/diagnostic.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DiagnosticServiceClient is the client API for DiagnosticService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiagnosticServiceClient interface {
	Create(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Diagnostic, error)
	Read(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Diagnostic, error)
	Update(ctx context.Context, in *Diagnostic, opts ...grpc.CallOption) (*Diagnostic, error)
	List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DiagnosticList, error)
}

type diagnosticServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDiagnosticServiceClient(cc grpc.ClientConnInterface) DiagnosticServiceClient {
	return &diagnosticServiceClient{cc}
}

func (c *diagnosticServiceClient) Create(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*Diagnostic, error) {
	out := new(Diagnostic)
	err := c.cc.Invoke(ctx, "/DiagnosticService/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diagnosticServiceClient) Read(ctx context.Context, in *ID, opts ...grpc.CallOption) (*Diagnostic, error) {
	out := new(Diagnostic)
	err := c.cc.Invoke(ctx, "/DiagnosticService/Read", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diagnosticServiceClient) Update(ctx context.Context, in *Diagnostic, opts ...grpc.CallOption) (*Diagnostic, error) {
	out := new(Diagnostic)
	err := c.cc.Invoke(ctx, "/DiagnosticService/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diagnosticServiceClient) List(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DiagnosticList, error) {
	out := new(DiagnosticList)
	err := c.cc.Invoke(ctx, "/DiagnosticService/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiagnosticServiceServer is the server API for DiagnosticService service.
// All implementations must embed UnimplementedDiagnosticServiceServer
// for forward compatibility
type DiagnosticServiceServer interface {
	Create(context.Context, *emptypb.Empty) (*Diagnostic, error)
	Read(context.Context, *ID) (*Diagnostic, error)
	Update(context.Context, *Diagnostic) (*Diagnostic, error)
	List(context.Context, *emptypb.Empty) (*DiagnosticList, error)
	mustEmbedUnimplementedDiagnosticServiceServer()
}

// UnimplementedDiagnosticServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDiagnosticServiceServer struct {
}

func (UnimplementedDiagnosticServiceServer) Create(context.Context, *emptypb.Empty) (*Diagnostic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedDiagnosticServiceServer) Read(context.Context, *ID) (*Diagnostic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedDiagnosticServiceServer) Update(context.Context, *Diagnostic) (*Diagnostic, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedDiagnosticServiceServer) List(context.Context, *emptypb.Empty) (*DiagnosticList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedDiagnosticServiceServer) mustEmbedUnimplementedDiagnosticServiceServer() {}

// UnsafeDiagnosticServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiagnosticServiceServer will
// result in compilation errors.
type UnsafeDiagnosticServiceServer interface {
	mustEmbedUnimplementedDiagnosticServiceServer()
}

func RegisterDiagnosticServiceServer(s grpc.ServiceRegistrar, srv DiagnosticServiceServer) {
	s.RegisterService(&DiagnosticService_ServiceDesc, srv)
}

func _DiagnosticService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiagnosticServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DiagnosticService/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiagnosticServiceServer).Create(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiagnosticService_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiagnosticServiceServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DiagnosticService/Read",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiagnosticServiceServer).Read(ctx, req.(*ID))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiagnosticService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Diagnostic)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiagnosticServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DiagnosticService/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiagnosticServiceServer).Update(ctx, req.(*Diagnostic))
	}
	return interceptor(ctx, in, info, handler)
}

func _DiagnosticService_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiagnosticServiceServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DiagnosticService/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiagnosticServiceServer).List(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// DiagnosticService_ServiceDesc is the grpc.ServiceDesc for DiagnosticService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DiagnosticService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DiagnosticService",
	HandlerType: (*DiagnosticServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _DiagnosticService_Create_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _DiagnosticService_Read_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _DiagnosticService_Update_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DiagnosticService_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/diagnostic.proto",
}
