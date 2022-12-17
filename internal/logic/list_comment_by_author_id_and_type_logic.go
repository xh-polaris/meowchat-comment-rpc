package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-comment-rpc/internal/common"
	"github.com/xh-polaris/meowchat-comment-rpc/pb"

	"github.com/xh-polaris/meowchat-comment-rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentByAuthorIdAndTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCommentByAuthorIdAndTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentByAuthorIdAndTypeLogic {
	return &ListCommentByAuthorIdAndTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据 authorId & type 查找
func (l *ListCommentByAuthorIdAndTypeLogic) ListCommentByAuthorIdAndType(in *pb.ListCommentByAuthorIdAndTypeRequest) (*pb.ListCommentByAuthorIdAndTypeResponse, error) {
	data, err := l.svcCtx.CommentModel.FindByAuthorIdAndType(l.ctx, in.AuthorId, in.Type, in.Skip, in.Limit)
	if err != nil {
		return nil, err
	}
	res := make([]*pb.Comment, 0, len(data))
	for _, val := range data {
		res = append(res, common.CommentConvert(val))
	}
	return &pb.ListCommentByAuthorIdAndTypeResponse{
		Comments: res,
	}, nil
}
