package server

import (
	"github.com/OhYee/goutils/bytes"
	"io"
)

// Info information of server
type Info struct {
	Name     string
	Addr     string
	SendTime int64
	RecvTime int64
}

// NewInfo initial a Info object
func NewInfo(name string, addr string, sendTime int64, recvTime int64) Info {
	return Info{
		Name:     name,
		Addr:     addr,
		SendTime: sendTime,
		RecvTime: recvTime,
	}
}

// NewInfoFromBytes initial a Info object from []bytes
func NewInfoFromBytes(r io.Reader) (info Info, err error) {
	var name, addr string
	var sendTime, recvTime int64

	if name, err = bytes.ReadStringWithLength32(r); err != nil {
		return
	}
	if addr, err = bytes.ReadStringWithLength32(r); err != nil {
		return
	}
	if sendTime, err = bytes.ReadInt64(r); err != nil {
		return
	}
	if recvTime, err = bytes.ReadInt64(r); err != nil {
		return
	}

	info = NewInfo(name, addr, sendTime, recvTime)
	return
}

// ToBytes transfer Info to []byte
func (info Info) ToBytes() []byte {
	buf := bytes.NewBuffer()
	buf.Write(bytes.FromStringWithLength32(info.Name))
	buf.Write(bytes.FromStringWithLength32(info.Addr))
	buf.Write(bytes.FromInt64(info.SendTime))
	buf.Write(bytes.FromInt64(info.RecvTime))
	return buf.Bytes()
}
