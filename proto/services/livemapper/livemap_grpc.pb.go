// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: services/livemapper/livemap.proto

package livemapper

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

// LivemapperServiceClient is the client API for LivemapperService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LivemapperServiceClient interface {
	// @permission: PerJob=true
	Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (LivemapperService_StreamClient, error)
}

type livemapperServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLivemapperServiceClient(cc grpc.ClientConnInterface) LivemapperServiceClient {
	return &livemapperServiceClient{cc}
}

func (c *livemapperServiceClient) Stream(ctx context.Context, in *StreamRequest, opts ...grpc.CallOption) (LivemapperService_StreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &LivemapperService_ServiceDesc.Streams[0], "/services.livemapper.LivemapperService/Stream", opts...)
	if err != nil {
		return nil, err
	}
	x := &livemapperServiceStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LivemapperService_StreamClient interface {
	Recv() (*ServerStreamResponse, error)
	grpc.ClientStream
}

type livemapperServiceStreamClient struct {
	grpc.ClientStream
}

func (x *livemapperServiceStreamClient) Recv() (*ServerStreamResponse, error) {
	m := new(ServerStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LivemapperServiceServer is the server API for LivemapperService service.
// All implementations must embed UnimplementedLivemapperServiceServer
// for forward compatibility
type LivemapperServiceServer interface {
	// @permission: PerJob=true
	Stream(*StreamRequest, LivemapperService_StreamServer) error
	mustEmbedUnimplementedLivemapperServiceServer()
}

// UnimplementedLivemapperServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLivemapperServiceServer struct {
}

func (UnimplementedLivemapperServiceServer) Stream(*StreamRequest, LivemapperService_StreamServer) error {
	return status.Errorf(codes.Unimplemented, "method Stream not implemented")
}
func (UnimplementedLivemapperServiceServer) mustEmbedUnimplementedLivemapperServiceServer() {}

// UnsafeLivemapperServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LivemapperServiceServer will
// result in compilation errors.
type UnsafeLivemapperServiceServer interface {
	mustEmbedUnimplementedLivemapperServiceServer()
}

func RegisterLivemapperServiceServer(s grpc.ServiceRegistrar, srv LivemapperServiceServer) {
	s.RegisterService(&LivemapperService_ServiceDesc, srv)
}

func _LivemapperService_Stream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LivemapperServiceServer).Stream(m, &livemapperServiceStreamServer{stream})
}

type LivemapperService_StreamServer interface {
	Send(*ServerStreamResponse) error
	grpc.ServerStream
}

type livemapperServiceStreamServer struct {
	grpc.ServerStream
}

func (x *livemapperServiceStreamServer) Send(m *ServerStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

// LivemapperService_ServiceDesc is the grpc.ServiceDesc for LivemapperService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LivemapperService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "services.livemapper.LivemapperService",
	HandlerType: (*LivemapperServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Stream",
			Handler:       _LivemapperService_Stream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "services/livemapper/livemap.proto",
}
