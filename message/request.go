package msg

import (
	"github.com/OhYee/goutils/bytes"
	"io"
)

// Request message body
type Request struct {
	ServerName string
	APIName    string
	Arguments  []byte
}

// NewRequest initial a Request
func NewRequest(serverName string, apiName string, args []byte) *Request {
	return &Request{
		ServerName: serverName,
		APIName:    apiName,
		Arguments:  args,
	}
}

// NewRequestFromBytes initial a Request from []byte
func NewRequestFromBytes(r io.Reader) (req *Request, err error) {
	var serverName, apiName string
	var args []byte
	if serverName, err = bytes.ReadStringWithLength32(r); err != nil {
		return
	}
	if apiName, err = bytes.ReadStringWithLength32(r); err != nil {
		return
	}
	if args, err = bytes.ReadBytesWithLength32(r); err != nil {
		return
	}
	req = NewRequest(serverName, apiName, args)
	return
}

// ToBytes transfer Request to []byte
func (req *Request) ToBytes() []byte {
	buf := bytes.NewBuffer()
	buf.Write(bytes.FromStringWithLength32(req.ServerName))
	buf.Write(bytes.FromStringWithLength32(req.APIName))
	buf.Write(bytes.FromBytesWithLength32(req.Arguments))
	return buf.Bytes()
}

// ToMessage initial a Request message
func (req *Request) ToMessage() *Message {
	return NewMessage(MessageTypeRequest, req)
}
