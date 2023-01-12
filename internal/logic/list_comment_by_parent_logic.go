package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-comment-rpc/internal/common"
	"github.com/xh-polaris/meowchat-comment-rpc/pb"

	"github.com/xh-polaris/meowchat-comment-rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentByParentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCommentByParentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentByParentLogic {
	return &ListCommentByParentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据 parentId 查找
func (l *ListCommentByParentLogic) ListCommentByParent(in *pb.ListCommentByParentRequest) (*pb.ListCommentByParentResponse, error) {
	data, count, err := l.svcCtx.CommentModel.FindByParent(l.ctx, in.Type, in.ParentId, in.Skip, in.Limit)
	if err != nil {
		return nil, err
	}

	res := pb.ListCommentByParentResponse{
		Comments: make([]*pb.Comment, 0, len(data)),
		Total:    count,
	}
	for _, val := range data {
		res.Comments = append(res.Comments,
			common.CommentConvert(val),
		)
	}
	return &res, nil
}
