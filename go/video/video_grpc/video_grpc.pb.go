// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: video.proto

package video_grpc

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	VideoService_Videos_FullMethodName      = "/video.VideoService/Videos"
	VideoService_Video_FullMethodName       = "/video.VideoService/Video"
	VideoService_UploadVideo_FullMethodName = "/video.VideoService/UploadVideo"
)

// VideoServiceClient is the client API for VideoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VideoServiceClient interface {
	Videos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VideosResponse, error)
	Video(ctx context.Context, in *VideoID, opts ...grpc.CallOption) (*VideoPayload, error)
	UploadVideo(ctx context.Context, in *UploadVideoInput, opts ...grpc.CallOption) (*VideoPayload, error)
}

type videoServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewVideoServiceClient(cc grpc.ClientConnInterface) VideoServiceClient {
	return &videoServiceClient{cc}
}

func (c *videoServiceClient) Videos(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*VideosResponse, error) {
	out := new(VideosResponse)
	err := c.cc.Invoke(ctx, VideoService_Videos_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) Video(ctx context.Context, in *VideoID, opts ...grpc.CallOption) (*VideoPayload, error) {
	out := new(VideoPayload)
	err := c.cc.Invoke(ctx, VideoService_Video_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *videoServiceClient) UploadVideo(ctx context.Context, in *UploadVideoInput, opts ...grpc.CallOption) (*VideoPayload, error) {
	out := new(VideoPayload)
	err := c.cc.Invoke(ctx, VideoService_UploadVideo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VideoServiceServer is the server API for VideoService service.
// All implementations must embed UnimplementedVideoServiceServer
// for forward compatibility
type VideoServiceServer interface {
	Videos(context.Context, *empty.Empty) (*VideosResponse, error)
	Video(context.Context, *VideoID) (*VideoPayload, error)
	UploadVideo(context.Context, *UploadVideoInput) (*VideoPayload, error)
	mustEmbedUnimplementedVideoServiceServer()
}

// UnimplementedVideoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedVideoServiceServer struct {
}

func (UnimplementedVideoServiceServer) Videos(context.Context, *empty.Empty) (*VideosResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Videos not implemented")
}
func (UnimplementedVideoServiceServer) Video(context.Context, *VideoID) (*VideoPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Video not implemented")
}
func (UnimplementedVideoServiceServer) UploadVideo(context.Context, *UploadVideoInput) (*VideoPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadVideo not implemented")
}
func (UnimplementedVideoServiceServer) mustEmbedUnimplementedVideoServiceServer() {}

// UnsafeVideoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VideoServiceServer will
// result in compilation errors.
type UnsafeVideoServiceServer interface {
	mustEmbedUnimplementedVideoServiceServer()
}

func RegisterVideoServiceServer(s grpc.ServiceRegistrar, srv VideoServiceServer) {
	s.RegisterService(&VideoService_ServiceDesc, srv)
}

func _VideoService_Videos_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).Videos(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_Videos_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).Videos(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_Video_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).Video(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_Video_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).Video(ctx, req.(*VideoID))
	}
	return interceptor(ctx, in, info, handler)
}

func _VideoService_UploadVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadVideoInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VideoServiceServer).UploadVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: VideoService_UploadVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VideoServiceServer).UploadVideo(ctx, req.(*UploadVideoInput))
	}
	return interceptor(ctx, in, info, handler)
}

// VideoService_ServiceDesc is the grpc.ServiceDesc for VideoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VideoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "video.VideoService",
	HandlerType: (*VideoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Videos",
			Handler:    _VideoService_Videos_Handler,
		},
		{
			MethodName: "Video",
			Handler:    _VideoService_Video_Handler,
		},
		{
			MethodName: "UploadVideo",
			Handler:    _VideoService_UploadVideo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "video.proto",
}