// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: modules/video/pb/rpc.proto

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

// VideoClient is the client API for Video service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoClient interface {
	Healthz(ctx context.Context, in *HealthzRequest, opts ...grpc.CallOption) (*HealthzResponse, error)
	GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*GetVideoResponse, error)
	ListVideo(ctx context.Context, in *ListVideoRequest, opts ...grpc.CallOption) (*ListVideoResponse, error)
	UploadVideo(ctx context.Context, opts ...grpc.CallOption) (Video_UploadVideoClient, error)
	DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*DeleteVideoResponse, error)
}

type videoClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoClient(cc grpc.ClientConnInterface) VideoClient {
	return &videoClient{cc}
}

func (c *videoClient) Healthz(ctx context.Context, in *HealthzRequest, opts ...grpc.CallOption) (*HealthzResponse, error) {
	out := new(HealthzResponse)
	err := c.cc.Invoke(ctx, "/pb.Video/Healthz", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) GetVideo(ctx context.Context, in *GetVideoRequest, opts ...grpc.CallOption) (*GetVideoResponse, error) {
	out := new(GetVideoResponse)
	err := c.cc.Invoke(ctx, "/pb.Video/GetVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) ListVideo(ctx context.Context, in *ListVideoRequest, opts ...grpc.CallOption) (*ListVideoResponse, error) {
	out := new(ListVideoResponse)
	err := c.cc.Invoke(ctx, "/pb.Video/ListVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoClient) UploadVideo(ctx context.Context, opts ...grpc.CallOption) (Video_UploadVideoClient, error) {
	stream, err := c.cc.NewStream(ctx, &Video_ServiceDesc.Streams[0], "/pb.Video/UploadVideo", opts...)
	if err != nil {
		return nil, err
	}
	x := &videoUploadVideoClient{stream}
	return x, nil
}

type Video_UploadVideoClient interface {
	Send(*UploadVideoRequest) error
	CloseAndRecv() (*UploadVideoResponse, error)
	grpc.ClientStream
}

type videoUploadVideoClient struct {
	grpc.ClientStream
}

func (x *videoUploadVideoClient) Send(m *UploadVideoRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *videoUploadVideoClient) CloseAndRecv() (*UploadVideoResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UploadVideoResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *videoClient) DeleteVideo(ctx context.Context, in *DeleteVideoRequest, opts ...grpc.CallOption) (*DeleteVideoResponse, error) {
	out := new(DeleteVideoResponse)
	err := c.cc.Invoke(ctx, "/pb.Video/DeleteVideo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServer is the server API for Video service.
// All implementations must embed UnimplementedVideoServer
// for forward compatibility
type VideoServer interface {
	Healthz(context.Context, *HealthzRequest) (*HealthzResponse, error)
	GetVideo(context.Context, *GetVideoRequest) (*GetVideoResponse, error)
	ListVideo(context.Context, *ListVideoRequest) (*ListVideoResponse, error)
	UploadVideo(Video_UploadVideoServer) error
	DeleteVideo(context.Context, *DeleteVideoRequest) (*DeleteVideoResponse, error)
	mustEmbedUnimplementedVideoServer()
}

// UnimplementedVideoServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServer struct {
}

func (UnimplementedVideoServer) Healthz(context.Context, *HealthzRequest) (*HealthzResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Healthz not implemented")
}
func (UnimplementedVideoServer) GetVideo(context.Context, *GetVideoRequest) (*GetVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideo not implemented")
}
func (UnimplementedVideoServer) ListVideo(context.Context, *ListVideoRequest) (*ListVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListVideo not implemented")
}
func (UnimplementedVideoServer) UploadVideo(Video_UploadVideoServer) error {
	return status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (UnimplementedVideoServer) DeleteVideo(context.Context, *DeleteVideoRequest) (*DeleteVideoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVideo not implemented")
}
func (UnimplementedVideoServer) mustEmbedUnimplementedVideoServer() {}

// UnsafeVideoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServer will
// result in compilation errors.
type UnsafeVideoServer interface {
	mustEmbedUnimplementedVideoServer()
}

func RegisterVideoServer(s grpc.ServiceRegistrar, srv VideoServer) {
	s.RegisterService(&Video_ServiceDesc, srv)
}

func _Video_Healthz_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthzRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).Healthz(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Video/Healthz",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).Healthz(ctx, req.(*HealthzRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_GetVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).GetVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Video/GetVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).GetVideo(ctx, req.(*GetVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_ListVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).ListVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Video/ListVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).ListVideo(ctx, req.(*ListVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Video_UploadVideo_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VideoServer).UploadVideo(&videoUploadVideoServer{stream})
}

type Video_UploadVideoServer interface {
	SendAndClose(*UploadVideoResponse) error
	Recv() (*UploadVideoRequest, error)
	grpc.ServerStream
}

type videoUploadVideoServer struct {
	grpc.ServerStream
}

func (x *videoUploadVideoServer) SendAndClose(m *UploadVideoResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *videoUploadVideoServer) Recv() (*UploadVideoRequest, error) {
	m := new(UploadVideoRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Video_DeleteVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteVideoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServer).DeleteVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Video/DeleteVideo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServer).DeleteVideo(ctx, req.(*DeleteVideoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Video_ServiceDesc is the grpc.ServiceDesc for Video service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Video_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Video",
	HandlerType: (*VideoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Healthz",
			Handler:    _Video_Healthz_Handler,
		},
		{
			MethodName: "GetVideo",
			Handler:    _Video_GetVideo_Handler,
		},
		{
			MethodName: "ListVideo",
			Handler:    _Video_ListVideo_Handler,
		},
		{
			MethodName: "DeleteVideo",
			Handler:    _Video_DeleteVideo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadVideo",
			Handler:       _Video_UploadVideo_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "modules/video/pb/rpc.proto",
}
