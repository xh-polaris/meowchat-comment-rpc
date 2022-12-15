package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-post-rpc/internal/model/mongo/commentcached"
	"github.com/xh-polaris/meowchat-post-rpc/internal/model/mongo/commenthistory"
	"github.com/xh-polaris/meowchat-post-rpc/pb"

	"github.com/xh-polaris/meowchat-post-rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentLogic {
	return &UpdateCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateCommentLogic) saveToHistory(data *commentcached.Comment) error {
	return l.svcCtx.HistoryModel.Insert(l.ctx, &commenthistory.History{
		Type: "update",
		Data: *data,
	})
}

// 修改
func (l *UpdateCommentLogic) UpdateComment(in *pb.UpdateCommentRequest) (*pb.UpdateCommentResponse, error) {
	old, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	err = l.saveToHistory(old)
	old.Text = in.Text
	err = l.svcCtx.CommentModel.Update(l.ctx, old)
	if err != nil {
		return nil, err
	}
	return &pb.UpdateCommentResponse{}, nil
}