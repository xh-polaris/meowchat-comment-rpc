package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-post-rpc/internal/model/mongo/commentcached"
	"github.com/xh-polaris/meowchat-post-rpc/internal/model/mongo/commenthistory"
	"github.com/xh-polaris/meowchat-post-rpc/pb"

	"github.com/xh-polaris/meowchat-post-rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCommentLogic {
	return &DeleteCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 保存
func (l *DeleteCommentLogic) saveToHistory(data *commentcached.Comment) error {
	return l.svcCtx.HistoryModel.Insert(l.ctx, &commenthistory.History{
		Type: "delete",
		Data: *data,
	})
}

// 删除
func (l *DeleteCommentLogic) DeleteComment(in *pb.DeleteCommentByIdRequest) (*pb.DeleteCommentByIdResponse, error) {
	data, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.Id)
	err = l.saveToHistory(data)
	err = l.svcCtx.CommentModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteCommentByIdResponse{}, nil
}
