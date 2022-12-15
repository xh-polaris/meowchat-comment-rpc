package common

import (
	"github.com/xh-polaris/meowchat-post-rpc/internal/model/mongo/commentcached"
	"github.com/xh-polaris/meowchat-post-rpc/pb"
)

func CommentConvert(in commentcached.Comment) *pb.Comment {
	return &pb.Comment{
		Id:       in.ID.Hex(),
		Text:     in.Text,
		AuthorId: in.AuthorId,
		ReplyTo:  in.ReplyTo,
		Type:     in.Type,
		ParentId: in.ParentId,
		UpdateAt: in.UpdateAt.Unix(),
		CreateAt: in.CreateAt.Unix(),
	}
}
