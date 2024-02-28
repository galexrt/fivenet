// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.4
// source: services/jobs/jobs.proto

package jobs

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

const (
	JobsService_ListColleagues_FullMethodName        = "/services.jobs.JobsService/ListColleagues"
	JobsService_GetSelf_FullMethodName               = "/services.jobs.JobsService/GetSelf"
	JobsService_GetColleague_FullMethodName          = "/services.jobs.JobsService/GetColleague"
	JobsService_ListColleagueActivity_FullMethodName = "/services.jobs.JobsService/ListColleagueActivity"
	JobsService_SetJobsUserProps_FullMethodName      = "/services.jobs.JobsService/SetJobsUserProps"
	JobsService_GetMOTD_FullMethodName               = "/services.jobs.JobsService/GetMOTD"
	JobsService_SetMOTD_FullMethodName               = "/services.jobs.JobsService/SetMOTD"
)

// JobsServiceClient is the client API for JobsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JobsServiceClient interface {
	// @perm
	ListColleagues(ctx context.Context, in *ListColleaguesRequest, opts ...grpc.CallOption) (*ListColleaguesResponse, error)
	// @perm: Name=ListColleagues
	GetSelf(ctx context.Context, in *GetSelfRequest, opts ...grpc.CallOption) (*GetSelfResponse, error)
	// @perm
	GetColleague(ctx context.Context, in *GetColleagueRequest, opts ...grpc.CallOption) (*GetColleagueResponse, error)
	// @perm: Name=GetColleague
	ListColleagueActivity(ctx context.Context, in *ListColleagueActivityRequest, opts ...grpc.CallOption) (*ListColleagueActivityResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "Lower_Rank", "Same_Rank", "Any"}
	SetJobsUserProps(ctx context.Context, in *SetJobsUserPropsRequest, opts ...grpc.CallOption) (*SetJobsUserPropsResponse, error)
	// @perm: Name=Any
	GetMOTD(ctx context.Context, in *GetMOTDRequest, opts ...grpc.CallOption) (*GetMOTDResponse, error)
	// @perm
	SetMOTD(ctx context.Context, in *SetMOTDRequest, opts ...grpc.CallOption) (*SetMOTDResponse, error)
}

type jobsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJobsServiceClient(cc grpc.ClientConnInterface) JobsServiceClient {
	return &jobsServiceClient{cc}
}

func (c *jobsServiceClient) ListColleagues(ctx context.Context, in *ListColleaguesRequest, opts ...grpc.CallOption) (*ListColleaguesResponse, error) {
	out := new(ListColleaguesResponse)
	err := c.cc.Invoke(ctx, JobsService_ListColleagues_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) GetSelf(ctx context.Context, in *GetSelfRequest, opts ...grpc.CallOption) (*GetSelfResponse, error) {
	out := new(GetSelfResponse)
	err := c.cc.Invoke(ctx, JobsService_GetSelf_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) GetColleague(ctx context.Context, in *GetColleagueRequest, opts ...grpc.CallOption) (*GetColleagueResponse, error) {
	out := new(GetColleagueResponse)
	err := c.cc.Invoke(ctx, JobsService_GetColleague_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) ListColleagueActivity(ctx context.Context, in *ListColleagueActivityRequest, opts ...grpc.CallOption) (*ListColleagueActivityResponse, error) {
	out := new(ListColleagueActivityResponse)
	err := c.cc.Invoke(ctx, JobsService_ListColleagueActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) SetJobsUserProps(ctx context.Context, in *SetJobsUserPropsRequest, opts ...grpc.CallOption) (*SetJobsUserPropsResponse, error) {
	out := new(SetJobsUserPropsResponse)
	err := c.cc.Invoke(ctx, JobsService_SetJobsUserProps_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) GetMOTD(ctx context.Context, in *GetMOTDRequest, opts ...grpc.CallOption) (*GetMOTDResponse, error) {
	out := new(GetMOTDResponse)
	err := c.cc.Invoke(ctx, JobsService_GetMOTD_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jobsServiceClient) SetMOTD(ctx context.Context, in *SetMOTDRequest, opts ...grpc.CallOption) (*SetMOTDResponse, error) {
	out := new(SetMOTDResponse)
	err := c.cc.Invoke(ctx, JobsService_SetMOTD_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JobsServiceServer is the server API for JobsService service.
// All implementations must embed UnimplementedJobsServiceServer
// for forward compatibility
type JobsServiceServer interface {
	// @perm
	ListColleagues(context.Context, *ListColleaguesRequest) (*ListColleaguesResponse, error)
	// @perm: Name=ListColleagues
	GetSelf(context.Context, *GetSelfRequest) (*GetSelfResponse, error)
	// @perm
	GetColleague(context.Context, *GetColleagueRequest) (*GetColleagueResponse, error)
	// @perm: Name=GetColleague
	ListColleagueActivity(context.Context, *ListColleagueActivityRequest) (*ListColleagueActivityResponse, error)
	// @perm: Attrs=Access/StringList:[]string{"Own", "Lower_Rank", "Same_Rank", "Any"}
	SetJobsUserProps(context.Context, *SetJobsUserPropsRequest) (*SetJobsUserPropsResponse, error)
	// @perm: Name=Any
	GetMOTD(context.Context, *GetMOTDRequest) (*GetMOTDResponse, error)
	// @perm
	SetMOTD(context.Context, *SetMOTDRequest) (*SetMOTDResponse, error)
	mustEmbedUnimplementedJobsServiceServer()
}

// UnimplementedJobsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJobsServiceServer struct {
}

func (UnimplementedJobsServiceServer) ListColleagues(context.Context, *ListColleaguesRequest) (*ListColleaguesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListColleagues not implemented")
}
func (UnimplementedJobsServiceServer) GetSelf(context.Context, *GetSelfRequest) (*GetSelfResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSelf not implemented")
}
func (UnimplementedJobsServiceServer) GetColleague(context.Context, *GetColleagueRequest) (*GetColleagueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetColleague not implemented")
}
func (UnimplementedJobsServiceServer) ListColleagueActivity(context.Context, *ListColleagueActivityRequest) (*ListColleagueActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListColleagueActivity not implemented")
}
func (UnimplementedJobsServiceServer) SetJobsUserProps(context.Context, *SetJobsUserPropsRequest) (*SetJobsUserPropsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetJobsUserProps not implemented")
}
func (UnimplementedJobsServiceServer) GetMOTD(context.Context, *GetMOTDRequest) (*GetMOTDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMOTD not implemented")
}
func (UnimplementedJobsServiceServer) SetMOTD(context.Context, *SetMOTDRequest) (*SetMOTDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetMOTD not implemented")
}
func (UnimplementedJobsServiceServer) mustEmbedUnimplementedJobsServiceServer() {}

// UnsafeJobsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JobsServiceServer will
// result in compilation errors.
type UnsafeJobsServiceServer interface {
	mustEmbedUnimplementedJobsServiceServer()
}

func RegisterJobsServiceServer(s grpc.ServiceRegistrar, srv JobsServiceServer) {
	s.RegisterService(&JobsService_ServiceDesc, srv)
}

func _JobsService_ListColleagues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListColleaguesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).ListColleagues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_ListColleagues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).ListColleagues(ctx, req.(*ListColleaguesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_GetSelf_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSelfRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).GetSelf(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_GetSelf_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).GetSelf(ctx, req.(*GetSelfRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_GetColleague_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetColleagueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).GetColleague(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_GetColleague_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).GetColleague(ctx, req.(*GetColleagueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_ListColleagueActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListColleagueActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).ListColleagueActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_ListColleagueActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).ListColleagueActivity(ctx, req.(*ListColleagueActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_SetJobsUserProps_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetJobsUserPropsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).SetJobsUserProps(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_SetJobsUserProps_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).SetJobsUserProps(ctx, req.(*SetJobsUserPropsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_GetMOTD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMOTDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).GetMOTD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_GetMOTD_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).GetMOTD(ctx, req.(*GetMOTDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JobsService_SetMOTD_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetMOTDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JobsServiceServer).SetMOTD(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: JobsService_SetMOTD_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JobsServiceServer).SetMOTD(ctx, req.(*SetMOTDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JobsService_ServiceDesc is the grpc.ServiceDesc for JobsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JobsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.jobs.JobsService",
	HandlerType: (*JobsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListColleagues",
			Handler:    _JobsService_ListColleagues_Handler,
		},
		{
			MethodName: "GetSelf",
			Handler:    _JobsService_GetSelf_Handler,
		},
		{
			MethodName: "GetColleague",
			Handler:    _JobsService_GetColleague_Handler,
		},
		{
			MethodName: "ListColleagueActivity",
			Handler:    _JobsService_ListColleagueActivity_Handler,
		},
		{
			MethodName: "SetJobsUserProps",
			Handler:    _JobsService_SetJobsUserProps_Handler,
		},
		{
			MethodName: "GetMOTD",
			Handler:    _JobsService_GetMOTD_Handler,
		},
		{
			MethodName: "SetMOTD",
			Handler:    _JobsService_SetMOTD_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services/jobs/jobs.proto",
}
