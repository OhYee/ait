package msg

import (
	"github.com/OhYee/goutils/bytes"
	"io"
)

// Response message body
type Response struct {
	Success   bool
	Arguments []byte
}

// NewResponse initial a Response
func NewResponse(success bool, args []byte) *Response {
	return &Response{
		Arguments: args,
	}
}

// NewResponseFromBytes initial a Response from []byte
func NewResponseFromBytes(r io.Reader) (rep *Response, err error) {
	var success, args []byte
	if success, err = bytes.ReadNBytes(r, 1); err != nil {
		return
	}
	if args, err = bytes.ReadBytesWithLength32(r); err != nil {
		return
	}
	rep = NewResponse(success[0] == 1, args)
	return
}

// ToBytes transfer Response to []byte
func (rep *Response) ToBytes() []byte {
	buf := bytes.NewBuffer()
	if rep.Success {
		buf.WriteByte(1)
	} else {
		buf.WriteByte(0)
	}
	buf.Write(bytes.FromBytesWithLength32(rep.Arguments))
	return buf.Bytes()
}

// ToMessage initial a Response message
func (rep *Response) ToMessage() *Message {
	return NewMessage(MessageTypeResponse, rep)
}
