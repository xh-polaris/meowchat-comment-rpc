package common

import (
	"context"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/apache/rocketmq-client-go/v2/primitive"
	"github.com/apache/rocketmq-client-go/v2/producer"
	"github.com/xh-polaris/meowchat-comment-rpc/commentrpc"
	"github.com/xh-polaris/meowchat-comment-rpc/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

var _ MsgQ = (*MsgQImpl)(nil)

type (
	MsgQ interface {
		SendCreateAsync(msg commentrpc.CommentMsg) error
		SendDeleteAsync(msg commentrpc.CommentMsg) error
	}
	MsgQImpl struct {
		producer rocketmq.Producer
	}
)

func callback(ctx context.Context, result *primitive.SendResult, err error) {
	if err != nil {
		logx.Error(err)
	}
}

func (m *MsgQImpl) SendDeleteAsync(msg commentrpc.CommentMsg) error {
	data, err := msg.Encode()
	if err != nil {
		return errorx.ErrMsgEncoder
	}
	err = m.producer.SendAsync(context.Background(), callback, &primitive.Message{Body: data, Topic: commentrpc.DeleteCommentTopic})
	if err != nil {
		return errorx.ErrMsgQ
	}
	return nil
}

func (m *MsgQImpl) SendCreateAsync(msg commentrpc.CommentMsg) error {
	data, err := msg.Encode()
	if err != nil {
		return err
	}
	err = m.producer.SendAsync(context.Background(), callback, &primitive.Message{Body: data, Topic: commentrpc.CreateCommentTopic})
	if err != nil {
		return err
	}
	return nil

}

func NewMsgQImpl(nameServer []string, retry int, groupName string) *MsgQImpl {
	mq, err := rocketmq.NewProducer(
		producer.WithNameServer(nameServer),
		producer.WithRetry(retry),
		producer.WithGroupName(groupName),
	)
	if err != nil {
		logx.Error(errorx.ErrMsgQ)
	}
	err = mq.Start()
	if err != nil {
		logx.Error(errorx.ErrMsgQ)
	}
	return &MsgQImpl{producer: mq}
}
