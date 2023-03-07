package logic

import (
	"context"

	"github.com/xh-polaris/meowchat-comment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-comment-rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CountCommentByParentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCountCommentByParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CountCommentByParentLogic {
	return &CountCommentByParentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据 parentId 统计
func (l *CountCommentByParentLogic) CountCommentByParent(in *pb.CountCommentByParentRequest) (*pb.CountCommentByParentResponse, error) {
	total, err := l.svcCtx.CommentModel.CountByParent(l.ctx, in.Type, in.ParentId)
	if err != nil {
		return nil, err
	}
	return &pb.CountCommentByParentResponse{Total: total}, nil
}
