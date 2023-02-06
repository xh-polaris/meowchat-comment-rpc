package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-comment-rpc/commentrpc"
	"github.com/xh-polaris/meowchat-comment-rpc/errorx"
	"github.com/xh-polaris/meowchat-comment-rpc/internal/model/mongo/commentcached"
	"github.com/xh-polaris/meowchat-comment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-comment-rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCommentLogic {
	return &CreateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建
func (l *CreateCommentLogic) CreateComment(in *pb.CreateCommentRequest) (*pb.CreateCommentResponse, error) {
	data := commentcached.Comment{
		Text:     in.Text,
		AuthorId: in.AuthorId,
		ReplyTo:  in.ReplyTo,
		Type:     in.Type,
		ParentId: in.ParentId,
	}
	if err := l.svcCtx.CommentModel.Insert(l.ctx, &data); err != nil {
		return nil, errorx.ErrDataBase
	}

	msg := commentrpc.CommentMsg{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorId,
		Type:     data.Type,
		ParentId: data.ParentId,
		Time:     data.UpdateAt.Unix(),
	}

	err := l.svcCtx.MsgQ.SendCreateAsync(msg)
	if err != nil {
		logx.Error(err)
	}

	return &pb.CreateCommentResponse{
		Id: data.ID.Hex(),
	}, nil
}
