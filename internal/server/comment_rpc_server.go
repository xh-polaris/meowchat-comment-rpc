// Code generated by goctl. DO NOT EDIT!
// Source: comment.proto

package server

import (
	"context"
	pb2 "github.com/xh-polaris/meowchat-post-rpc/pb"

	"github.com/xh-polaris/meowchat-post-rpc/internal/logic"
	"github.com/xh-polaris/meowchat-post-rpc/internal/svc"
)

type CommentRpcServer struct {
	svcCtx *svc.ServiceContext
	pb2.UnimplementedCommentRpcServer
}

func NewCommentRpcServer(svcCtx *svc.ServiceContext) *CommentRpcServer {
	return &CommentRpcServer{
		svcCtx: svcCtx,
	}
}

// 创建
func (s *CommentRpcServer) CreateComment(ctx context.Context, in *pb2.CreateCommentRequest) (*pb2.CreateCommentResponse, error) {
	l := logic.NewCreateCommentLogic(ctx, s.svcCtx)
	return l.CreateComment(in)
}

// 修改
func (s *CommentRpcServer) UpdateComment(ctx context.Context, in *pb2.UpdateCommentRequest) (*pb2.UpdateCommentResponse, error) {
	l := logic.NewUpdateCommentLogic(ctx, s.svcCtx)
	return l.UpdateComment(in)
}

// 删除
func (s *CommentRpcServer) DeleteComment(ctx context.Context, in *pb2.DeleteCommentByIdRequest) (*pb2.DeleteCommentByIdResponse, error) {
	l := logic.NewDeleteCommentLogic(ctx, s.svcCtx)
	return l.DeleteComment(in)
}

// 根据 parentId 查找
func (s *CommentRpcServer) ListCommentByParent(ctx context.Context, in *pb2.ListCommentByParentRequest) (*pb2.ListCommentByParentResponse, error) {
	l := logic.NewListCommentByParentLogic(ctx, s.svcCtx)
	return l.ListCommentByParent(in)
}

// 根据 id 查找
func (s *CommentRpcServer) RetrieveCommentById(ctx context.Context, in *pb2.RetrieveCommentByIdRequest) (*pb2.RetrieveCommentByIdResponse, error) {
	l := logic.NewRetrieveCommentByIdLogic(ctx, s.svcCtx)
	return l.RetrieveCommentById(in)
}

// 根据 authorId & type 查找
func (s *CommentRpcServer) ListCommentByAuthorIdAndType(ctx context.Context, in *pb2.ListCommentByAuthorIdAndTypeRequest) (*pb2.ListCommentByAuthorIdAndTypeResponse, error) {
	l := logic.NewListCommentByAuthorIdAndTypeLogic(ctx, s.svcCtx)
	return l.ListCommentByAuthorIdAndType(in)
}

// 根据 replyTo & type 查找
func (s *CommentRpcServer) ListCommentByReplyToAndType(ctx context.Context, in *pb2.ListCommentByReplyToAndTypeRequest) (*pb2.ListCommentByReplyToAndTypeResponse, error) {
	l := logic.NewListCommentByReplyToAndTypeLogic(ctx, s.svcCtx)
	return l.ListCommentByReplyToAndType(in)
}