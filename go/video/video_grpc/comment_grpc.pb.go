// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.12.4
// source: comment.proto

package video_grpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	CommentService_CommentsByVideo_FullMethodName = "/comment.CommentService/CommentsByVideo"
	CommentService_Comment_FullMethodName         = "/comment.CommentService/Comment"
	CommentService_PostComment_FullMethodName     = "/comment.CommentService/PostComment"
)

// CommentServiceClient is the client API for CommentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentServiceClient interface {
	CommentsByVideo(ctx context.Context, in *CommentsByVideoInput, opts ...grpc.CallOption) (*CommentsResponse, error)
	Comment(ctx context.Context, in *CommentID, opts ...grpc.CallOption) (*PostCommentPayload, error)
	PostComment(ctx context.Context, in *PostCommentInput, opts ...grpc.CallOption) (*PostCommentPayload, error)
}

type commentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentServiceClient(cc grpc.ClientConnInterface) CommentServiceClient {
	return &commentServiceClient{cc}
}

func (c *commentServiceClient) CommentsByVideo(ctx context.Context, in *CommentsByVideoInput, opts ...grpc.CallOption) (*CommentsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentsResponse)
	err := c.cc.Invoke(ctx, CommentService_CommentsByVideo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) Comment(ctx context.Context, in *CommentID, opts ...grpc.CallOption) (*PostCommentPayload, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostCommentPayload)
	err := c.cc.Invoke(ctx, CommentService_Comment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentServiceClient) PostComment(ctx context.Context, in *PostCommentInput, opts ...grpc.CallOption) (*PostCommentPayload, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostCommentPayload)
	err := c.cc.Invoke(ctx, CommentService_PostComment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentServiceServer is the server API for CommentService service.
// All implementations must embed UnimplementedCommentServiceServer
// for forward compatibility
type CommentServiceServer interface {
	CommentsByVideo(context.Context, *CommentsByVideoInput) (*CommentsResponse, error)
	Comment(context.Context, *CommentID) (*PostCommentPayload, error)
	PostComment(context.Context, *PostCommentInput) (*PostCommentPayload, error)
	mustEmbedUnimplementedCommentServiceServer()
}

// UnimplementedCommentServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCommentServiceServer struct {
}

func (UnimplementedCommentServiceServer) CommentsByVideo(context.Context, *CommentsByVideoInput) (*CommentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommentsByVideo not implemented")
}
func (UnimplementedCommentServiceServer) Comment(context.Context, *CommentID) (*PostCommentPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comment not implemented")
}
func (UnimplementedCommentServiceServer) PostComment(context.Context, *PostCommentInput) (*PostCommentPayload, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostComment not implemented")
}
func (UnimplementedCommentServiceServer) mustEmbedUnimplementedCommentServiceServer() {}

// UnsafeCommentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentServiceServer will
// result in compilation errors.
type UnsafeCommentServiceServer interface {
	mustEmbedUnimplementedCommentServiceServer()
}

func RegisterCommentServiceServer(s grpc.ServiceRegistrar, srv CommentServiceServer) {
	s.RegisterService(&CommentService_ServiceDesc, srv)
}

func _CommentService_CommentsByVideo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentsByVideoInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).CommentsByVideo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_CommentsByVideo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).CommentsByVideo(ctx, req.(*CommentsByVideoInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_Comment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentID)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).Comment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_Comment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).Comment(ctx, req.(*CommentID))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentService_PostComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostCommentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentServiceServer).PostComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CommentService_PostComment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentServiceServer).PostComment(ctx, req.(*PostCommentInput))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentService_ServiceDesc is the grpc.ServiceDesc for CommentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.CommentService",
	HandlerType: (*CommentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CommentsByVideo",
			Handler:    _CommentService_CommentsByVideo_Handler,
		},
		{
			MethodName: "Comment",
			Handler:    _CommentService_Comment_Handler,
		},
		{
			MethodName: "PostComment",
			Handler:    _CommentService_PostComment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}
