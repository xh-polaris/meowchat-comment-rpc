package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-post-rpc/internal/common"
	"github.com/xh-polaris/meowchat-post-rpc/pb"

	"github.com/xh-polaris/meowchat-post-rpc/internal/svc"
	"github.com/zeromicro/go-zero/core/logx"
)

type ListCommentByReplyToAndTypeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListCommentByReplyToAndTypeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListCommentByReplyToAndTypeLogic {
	return &ListCommentByReplyToAndTypeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据 replyTo & type 查找
func (l *ListCommentByReplyToAndTypeLogic) ListCommentByReplyToAndType(in *pb.ListCommentByReplyToAndTypeRequest) (*pb.ListCommentByReplyToAndTypeResponse, error) {
	data, err := l.svcCtx.CommentModel.FindByReplyToAndType(l.ctx, in.Type, in.ReplyTo, in.Skip, in.Limit)
	if err != nil {
		return nil, err
	}
	res := pb.ListCommentByReplyToAndTypeResponse{
		Comments: make([]*pb.Comment, 0, len(data)),
	}
	for _, val := range data {
		res.Comments = append(res.Comments,
			common.CommentConvert(val),
		)
	}
	return &res, nil
}
