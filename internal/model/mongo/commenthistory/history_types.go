package commenthistory

import (
	"github.com/xh-polaris/meowchat-comment-rpc/internal/model/mongo/commentcached"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type History struct {
	ID       primitive.ObjectID    `bson:"_id,omitempty" json:"id,omitempty"`
	Type     string                `bson:"type,omitempty" json:"type,omitempty"`
	Data     commentcached.Comment `bson:"data,omitempty" json:"data,omitempty"`
	UpdateAt time.Time             `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt time.Time             `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
