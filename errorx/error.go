package errorx

import (
	"google.golang.org/grpc/status"
)

var (
	ErrDataBase   = status.Error(16002, "database error")
	ErrMsgEncoder = status.Error(16003, "message encoding error")
	ErrMsgDecoder = status.Error(16004, "message decoding error")
	ErrMsgQ       = status.Error(16005, "message queue request failed")
)
