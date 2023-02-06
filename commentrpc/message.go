package commentrpc

import (
	"bytes"
	"encoding/gob"
)

const CreateCommentTopic = "CREATE_COMMENT_MSG"
const DeleteCommentTopic = "DELETE_COMMENT_MSG"

type CommentMsg struct {
	Id       string
	AuthorId string
	Type     string
	ParentId string
	Time     int64
}

func (m CommentMsg) Encode() ([]byte, error) {
	b := new(bytes.Buffer)
	e := gob.NewEncoder(b)
	if err := e.Encode(m); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

func (m *CommentMsg) Decode(data []byte) error {
	b := bytes.NewBuffer(data)
	d := gob.NewDecoder(b)
	if err := d.Decode(m); err != nil {
		return err
	}
	return nil
}
