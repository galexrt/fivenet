// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: services/citizenstore/citizenstore.proto

package citizenstore

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

// CitizenStoreServiceClient is the client API for CitizenStoreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CitizenStoreServiceClient interface {
	// @perm: Attrs=Fields/StringList:["PhoneNumber","Licenses","UserProps.Wanted","UserProps.Job"]
	ListCitizens(ctx context.Context, in *ListCitizensRequest, opts ...grpc.CallOption) (*ListCitizensResponse, error)
	// @perm: Attrs=Jobs/JobRankList
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// @perm: Attrs=Fields/StringList:["SourceUser"]
	ListUserActivity(ctx context.Context, in *ListUserActivityRequest, opts ...grpc.CallOption) (*ListUserActivityResponse, error)
	// @perm: Attrs=Fields/StringList:["Wanted","Job"]
	SetUserProps(ctx context.Context, in *SetUserPropsRequest, opts ...grpc.CallOption) (*SetUserPropsResponse, error)
}

type citizenStoreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCitizenStoreServiceClient(cc grpc.ClientConnInterface) CitizenStoreServiceClient {
	return &citizenStoreServiceClient{cc}
}

func (c *citizenStoreServiceClient) ListCitizens(ctx context.Context, in *ListCitizensRequest, opts ...grpc.CallOption) (*ListCitizensResponse, error) {
	out := new(ListCitizensResponse)
	err := c.cc.Invoke(ctx, "/services.citizenstore.CitizenStoreService/ListCitizens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *citizenStoreServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/services.citizenstore.CitizenStoreService/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *citizenStoreServiceClient) ListUserActivity(ctx context.Context, in *ListUserActivityRequest, opts ...grpc.CallOption) (*ListUserActivityResponse, error) {
	out := new(ListUserActivityResponse)
	err := c.cc.Invoke(ctx, "/services.citizenstore.CitizenStoreService/ListUserActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *citizenStoreServiceClient) SetUserProps(ctx context.Context, in *SetUserPropsRequest, opts ...grpc.CallOption) (*SetUserPropsResponse, error) {
	out := new(SetUserPropsResponse)
	err := c.cc.Invoke(ctx, "/services.citizenstore.CitizenStoreService/SetUserProps", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CitizenStoreServiceServer is the server API for CitizenStoreService service.
// All implementations must embed UnimplementedCitizenStoreServiceServer
// for forward compatibility
type CitizenStoreServiceServer interface {
	// @perm: Attrs=Fields/StringList:["PhoneNumber","Licenses","UserProps.Wanted","UserProps.Job"]
	ListCitizens(context.Context, *ListCitizensRequest) (*ListCitizensResponse, error)
	// @perm: Attrs=Jobs/JobRankList
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// @perm: Attrs=Fields/StringList:["SourceUser"]
	ListUserActivity(context.Context, *ListUserActivityRequest) (*ListUserActivityResponse, error)
	// @perm: Attrs=Fields/StringList:["Wanted","Job"]
	SetUserProps(context.Context, *SetUserPropsRequest) (*SetUserPropsResponse, error)
	mustEmbedUnimplementedCitizenStoreServiceServer()
}

// UnimplementedCitizenStoreServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCitizenStoreServiceServer struct {
}

func (UnimplementedCitizenStoreServiceServer) ListCitizens(context.Context, *ListCitizensRequest) (*ListCitizensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCitizens not implemented")
}
func (UnimplementedCitizenStoreServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedCitizenStoreServiceServer) ListUserActivity(context.Context, *ListUserActivityRequest) (*ListUserActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUserActivity not implemented")
}
func (UnimplementedCitizenStoreServiceServer) SetUserProps(context.Context, *SetUserPropsRequest) (*SetUserPropsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetUserProps not implemented")
}
func (UnimplementedCitizenStoreServiceServer) mustEmbedUnimplementedCitizenStoreServiceServer() {}

// UnsafeCitizenStoreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CitizenStoreServiceServer will
// result in compilation errors.
type UnsafeCitizenStoreServiceServer interface {
	mustEmbedUnimplementedCitizenStoreServiceServer()
}

func RegisterCitizenStoreServiceServer(s grpc.ServiceRegistrar, srv CitizenStoreServiceServer) {
	s.RegisterService(&CitizenStoreService_ServiceDesc, srv)
}

func _CitizenStoreService_ListCitizens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCitizensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitizenStoreServiceServer).ListCitizens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.citizenstore.CitizenStoreService/ListCitizens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitizenStoreServiceServer).ListCitizens(ctx, req.(*ListCitizensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CitizenStoreService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitizenStoreServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.citizenstore.CitizenStoreService/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitizenStoreServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CitizenStoreService_ListUserActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListUserActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitizenStoreServiceServer).ListUserActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.citizenstore.CitizenStoreService/ListUserActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitizenStoreServiceServer).ListUserActivity(ctx, req.(*ListUserActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CitizenStoreService_SetUserProps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetUserPropsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CitizenStoreServiceServer).SetUserProps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/services.citizenstore.CitizenStoreService/SetUserProps",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CitizenStoreServiceServer).SetUserProps(ctx, req.(*SetUserPropsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CitizenStoreService_ServiceDesc is the grpc.ServiceDesc for CitizenStoreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CitizenStoreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.citizenstore.CitizenStoreService",
	HandlerType: (*CitizenStoreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListCitizens",
			Handler:    _CitizenStoreService_ListCitizens_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _CitizenStoreService_GetUser_Handler,
		},
		{
			MethodName: "ListUserActivity",
			Handler:    _CitizenStoreService_ListUserActivity_Handler,
		},
		{
			MethodName: "SetUserProps",
			Handler:    _CitizenStoreService_SetUserProps_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/citizenstore/citizenstore.proto",
}
