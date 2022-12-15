// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: comment.proto

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

// CommentRpcClient is the client API for CommentRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommentRpcClient interface {
	// 创建
	CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error)
	// 修改
	UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error)
	// 删除
	DeleteComment(ctx context.Context, in *DeleteCommentByIdRequest, opts ...grpc.CallOption) (*DeleteCommentByIdResponse, error)
	// 根据 parentId 查找
	ListCommentByParent(ctx context.Context, in *ListCommentByParentRequest, opts ...grpc.CallOption) (*ListCommentByParentResponse, error)
	// 根据 id 查找
	RetrieveCommentById(ctx context.Context, in *RetrieveCommentByIdRequest, opts ...grpc.CallOption) (*RetrieveCommentByIdResponse, error)
	// 根据 authorId & type 查找
	ListCommentByAuthorIdAndType(ctx context.Context, in *ListCommentByAuthorIdAndTypeRequest, opts ...grpc.CallOption) (*ListCommentByAuthorIdAndTypeResponse, error)
	// 根据 replyTo & type 查找
	ListCommentByReplyToAndType(ctx context.Context, in *ListCommentByReplyToAndTypeRequest, opts ...grpc.CallOption) (*ListCommentByReplyToAndTypeResponse, error)
}

type commentRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewCommentRpcClient(cc grpc.ClientConnInterface) CommentRpcClient {
	return &commentRpcClient{cc}
}

func (c *commentRpcClient) CreateComment(ctx context.Context, in *CreateCommentRequest, opts ...grpc.CallOption) (*CreateCommentResponse, error) {
	out := new(CreateCommentResponse)
	err := c.cc.Invoke(ctx, "/comment.comment_rpc/createComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) UpdateComment(ctx context.Context, in *UpdateCommentRequest, opts ...grpc.CallOption) (*UpdateCommentResponse, error) {
	out := new(UpdateCommentResponse)
	err := c.cc.Invoke(ctx, "/comment.comment_rpc/updateComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) DeleteComment(ctx context.Context, in *DeleteCommentByIdRequest, opts ...grpc.CallOption) (*DeleteCommentByIdResponse, error) {
	out := new(DeleteCommentByIdResponse)
	err := c.cc.Invoke(ctx, "/comment.comment_rpc/deleteComment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) ListCommentByParent(ctx context.Context, in *ListCommentByParentRequest, opts ...grpc.CallOption) (*ListCommentByParentResponse, error) {
	out := new(ListCommentByParentResponse)
	err := c.cc.Invoke(ctx, "/comment.comment_rpc/listCommentByParent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) RetrieveCommentById(ctx context.Context, in *RetrieveCommentByIdRequest, opts ...grpc.CallOption) (*RetrieveCommentByIdResponse, error) {
	out := new(RetrieveCommentByIdResponse)
	err := c.cc.Invoke(ctx, "/comment.comment_rpc/retrieveCommentById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) ListCommentByAuthorIdAndType(ctx context.Context, in *ListCommentByAuthorIdAndTypeRequest, opts ...grpc.CallOption) (*ListCommentByAuthorIdAndTypeResponse, error) {
	out := new(ListCommentByAuthorIdAndTypeResponse)
	err := c.cc.Invoke(ctx, "/comment.comment_rpc/listCommentByAuthorIdAndType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commentRpcClient) ListCommentByReplyToAndType(ctx context.Context, in *ListCommentByReplyToAndTypeRequest, opts ...grpc.CallOption) (*ListCommentByReplyToAndTypeResponse, error) {
	out := new(ListCommentByReplyToAndTypeResponse)
	err := c.cc.Invoke(ctx, "/comment.comment_rpc/listCommentByReplyToAndType", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CommentRpcServer is the server API for CommentRpc service.
// All implementations must embed UnimplementedCommentRpcServer
// for forward compatibility
type CommentRpcServer interface {
	// 创建
	CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error)
	// 修改
	UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error)
	// 删除
	DeleteComment(context.Context, *DeleteCommentByIdRequest) (*DeleteCommentByIdResponse, error)
	// 根据 parentId 查找
	ListCommentByParent(context.Context, *ListCommentByParentRequest) (*ListCommentByParentResponse, error)
	// 根据 id 查找
	RetrieveCommentById(context.Context, *RetrieveCommentByIdRequest) (*RetrieveCommentByIdResponse, error)
	// 根据 authorId & type 查找
	ListCommentByAuthorIdAndType(context.Context, *ListCommentByAuthorIdAndTypeRequest) (*ListCommentByAuthorIdAndTypeResponse, error)
	// 根据 replyTo & type 查找
	ListCommentByReplyToAndType(context.Context, *ListCommentByReplyToAndTypeRequest) (*ListCommentByReplyToAndTypeResponse, error)
	mustEmbedUnimplementedCommentRpcServer()
}

// UnimplementedCommentRpcServer must be embedded to have forward compatible implementations.
type UnimplementedCommentRpcServer struct {
}

func (UnimplementedCommentRpcServer) CreateComment(context.Context, *CreateCommentRequest) (*CreateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateComment not implemented")
}
func (UnimplementedCommentRpcServer) UpdateComment(context.Context, *UpdateCommentRequest) (*UpdateCommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateComment not implemented")
}
func (UnimplementedCommentRpcServer) DeleteComment(context.Context, *DeleteCommentByIdRequest) (*DeleteCommentByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteComment not implemented")
}
func (UnimplementedCommentRpcServer) ListCommentByParent(context.Context, *ListCommentByParentRequest) (*ListCommentByParentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommentByParent not implemented")
}
func (UnimplementedCommentRpcServer) RetrieveCommentById(context.Context, *RetrieveCommentByIdRequest) (*RetrieveCommentByIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveCommentById not implemented")
}
func (UnimplementedCommentRpcServer) ListCommentByAuthorIdAndType(context.Context, *ListCommentByAuthorIdAndTypeRequest) (*ListCommentByAuthorIdAndTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommentByAuthorIdAndType not implemented")
}
func (UnimplementedCommentRpcServer) ListCommentByReplyToAndType(context.Context, *ListCommentByReplyToAndTypeRequest) (*ListCommentByReplyToAndTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCommentByReplyToAndType not implemented")
}
func (UnimplementedCommentRpcServer) mustEmbedUnimplementedCommentRpcServer() {}

// UnsafeCommentRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommentRpcServer will
// result in compilation errors.
type UnsafeCommentRpcServer interface {
	mustEmbedUnimplementedCommentRpcServer()
}

func RegisterCommentRpcServer(s grpc.ServiceRegistrar, srv CommentRpcServer) {
	s.RegisterService(&CommentRpc_ServiceDesc, srv)
}

func _CommentRpc_CreateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).CreateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.comment_rpc/createComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).CreateComment(ctx, req.(*CreateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_UpdateComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).UpdateComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.comment_rpc/updateComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).UpdateComment(ctx, req.(*UpdateCommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_DeleteComment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommentByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).DeleteComment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.comment_rpc/deleteComment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).DeleteComment(ctx, req.(*DeleteCommentByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_ListCommentByParent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommentByParentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).ListCommentByParent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.comment_rpc/listCommentByParent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).ListCommentByParent(ctx, req.(*ListCommentByParentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_RetrieveCommentById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RetrieveCommentByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).RetrieveCommentById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.comment_rpc/retrieveCommentById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).RetrieveCommentById(ctx, req.(*RetrieveCommentByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_ListCommentByAuthorIdAndType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommentByAuthorIdAndTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).ListCommentByAuthorIdAndType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.comment_rpc/listCommentByAuthorIdAndType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).ListCommentByAuthorIdAndType(ctx, req.(*ListCommentByAuthorIdAndTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommentRpc_ListCommentByReplyToAndType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCommentByReplyToAndTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommentRpcServer).ListCommentByReplyToAndType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/comment.comment_rpc/listCommentByReplyToAndType",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommentRpcServer).ListCommentByReplyToAndType(ctx, req.(*ListCommentByReplyToAndTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CommentRpc_ServiceDesc is the grpc.ServiceDesc for CommentRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommentRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "comment.comment_rpc",
	HandlerType: (*CommentRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "createComment",
			Handler:    _CommentRpc_CreateComment_Handler,
		},
		{
			MethodName: "updateComment",
			Handler:    _CommentRpc_UpdateComment_Handler,
		},
		{
			MethodName: "deleteComment",
			Handler:    _CommentRpc_DeleteComment_Handler,
		},
		{
			MethodName: "listCommentByParent",
			Handler:    _CommentRpc_ListCommentByParent_Handler,
		},
		{
			MethodName: "retrieveCommentById",
			Handler:    _CommentRpc_RetrieveCommentById_Handler,
		},
		{
			MethodName: "listCommentByAuthorIdAndType",
			Handler:    _CommentRpc_ListCommentByAuthorIdAndType_Handler,
		},
		{
			MethodName: "listCommentByReplyToAndType",
			Handler:    _CommentRpc_ListCommentByReplyToAndType_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "comment.proto",
}