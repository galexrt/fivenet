// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: services/completion/completion.proto

package completion

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

// CompletionServiceClient is the client API for CompletionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CompletionServiceClient interface {
	CompleteJobNames(ctx context.Context, in *CompleteJobNamesRequest, opts ...grpc.CallOption) (*CompleteJobNamesResponse, error)
	CompleteJobGrades(ctx context.Context, in *CompleteJobGradesRequest, opts ...grpc.CallOption) (*CompleteJobGradesResponse, error)
	CompleteDocumentCategory(ctx context.Context, in *CompleteDocumentCategoryRequest, opts ...grpc.CallOption) (*CompleteDocumentCategoryResponse, error)
}

type completionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCompletionServiceClient(cc grpc.ClientConnInterface) CompletionServiceClient {
	return &completionServiceClient{cc}
}

func (c *completionServiceClient) CompleteJobNames(ctx context.Context, in *CompleteJobNamesRequest, opts ...grpc.CallOption) (*CompleteJobNamesResponse, error) {
	out := new(CompleteJobNamesResponse)
	err := c.cc.Invoke(ctx, "/services.completion.CompletionService/CompleteJobNames", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *completionServiceClient) CompleteJobGrades(ctx context.Context, in *CompleteJobGradesRequest, opts ...grpc.CallOption) (*CompleteJobGradesResponse, error) {
	out := new(CompleteJobGradesResponse)
	err := c.cc.Invoke(ctx, "/services.completion.CompletionService/CompleteJobGrades", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *completionServiceClient) CompleteDocumentCategory(ctx context.Context, in *CompleteDocumentCategoryRequest, opts ...grpc.CallOption) (*CompleteDocumentCategoryResponse, error) {
	out := new(CompleteDocumentCategoryResponse)
	err := c.cc.Invoke(ctx, "/services.completion.CompletionService/CompleteDocumentCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CompletionServiceServer is the server API for CompletionService service.
// All implementations must embed UnimplementedCompletionServiceServer
// for forward compatibility
type CompletionServiceServer interface {
	CompleteJobNames(context.Context, *CompleteJobNamesRequest) (*CompleteJobNamesResponse, error)
	CompleteJobGrades(context.Context, *CompleteJobGradesRequest) (*CompleteJobGradesResponse, error)
	CompleteDocumentCategory(context.Context, *CompleteDocumentCategoryRequest) (*CompleteDocumentCategoryResponse, error)
	mustEmbedUnimplementedCompletionServiceServer()
}

// UnimplementedCompletionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCompletionServiceServer struct {
}

func (UnimplementedCompletionServiceServer) CompleteJobNames(context.Context, *CompleteJobNamesRequest) (*CompleteJobNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteJobNames not implemented")
}
func (UnimplementedCompletionServiceServer) CompleteJobGrades(context.Context, *CompleteJobGradesRequest) (*CompleteJobGradesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteJobGrades not implemented")
}
func (UnimplementedCompletionServiceServer) CompleteDocumentCategory(context.Context, *CompleteDocumentCategoryRequest) (*CompleteDocumentCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompleteDocumentCategory not implemented")
}
func (UnimplementedCompletionServiceServer) mustEmbedUnimplementedCompletionServiceServer() {}

// UnsafeCompletionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CompletionServiceServer will
// result in compilation errors.
type UnsafeCompletionServiceServer interface {
	mustEmbedUnimplementedCompletionServiceServer()
}

func RegisterCompletionServiceServer(s grpc.ServiceRegistrar, srv CompletionServiceServer) {
	s.RegisterService(&CompletionService_ServiceDesc, srv)
}

func _CompletionService_CompleteJobNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteJobNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompletionServiceServer).CompleteJobNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.completion.CompletionService/CompleteJobNames",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompletionServiceServer).CompleteJobNames(ctx, req.(*CompleteJobNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompletionService_CompleteJobGrades_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteJobGradesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompletionServiceServer).CompleteJobGrades(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.completion.CompletionService/CompleteJobGrades",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompletionServiceServer).CompleteJobGrades(ctx, req.(*CompleteJobGradesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CompletionService_CompleteDocumentCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompleteDocumentCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CompletionServiceServer).CompleteDocumentCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.completion.CompletionService/CompleteDocumentCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CompletionServiceServer).CompleteDocumentCategory(ctx, req.(*CompleteDocumentCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CompletionService_ServiceDesc is the grpc.ServiceDesc for CompletionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CompletionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.completion.CompletionService",
	HandlerType: (*CompletionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CompleteJobNames",
			Handler:    _CompletionService_CompleteJobNames_Handler,
		},
		{
			MethodName: "CompleteJobGrades",
			Handler:    _CompletionService_CompleteJobGrades_Handler,
		},
		{
			MethodName: "CompleteDocumentCategory",
			Handler:    _CompletionService_CompleteDocumentCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/completion/completion.proto",
}
