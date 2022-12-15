package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-post-rpc/internal/model/mongo/commentcached"
	"github.com/xh-polaris/meowchat-post-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-post-rpc/pb"
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
	err := l.svcCtx.CommentModel.Insert(l.ctx, &data)
	if err != nil {
		return nil, err
	}
	return &pb.CreateCommentResponse{
		Id: data.ID.Hex(),
	}, nil
}
