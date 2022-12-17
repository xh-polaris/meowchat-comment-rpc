package logic

import (
	"context"
	"github.com/xh-polaris/meowchat-comment-rpc/internal/common"
	"github.com/xh-polaris/meowchat-comment-rpc/internal/svc"
	"github.com/xh-polaris/meowchat-comment-rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type RetrieveCommentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRetrieveCommentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RetrieveCommentByIdLogic {
	return &RetrieveCommentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据 id 查找
func (l *RetrieveCommentByIdLogic) RetrieveCommentById(in *pb.RetrieveCommentByIdRequest) (*pb.RetrieveCommentByIdResponse, error) {
	ret, err := l.svcCtx.CommentModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.RetrieveCommentByIdResponse{
		Comment: common.CommentConvert(*ret),
	}, nil
}
