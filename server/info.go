package server

import (
	"github.com/OhYee/goutils/bytes"
	"io"
)

// Info information of server
type Info struct {
	name     string
	addr     string
	sendTime int64
	recvTime int64
}

// NewInfo initial a Info object
func NewInfo(name string, addr string, sendTime int64, recvTime int64) Info {
	return Info{
		name:     name,
		addr:     addr,
		sendTime: sendTime,
		recvTime: recvTime,
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
	buf.Write(bytes.FromStringWithLength32(info.name))
	buf.Write(bytes.FromStringWithLength32(info.addr))
	buf.Write(bytes.FromInt64(info.sendTime))
	buf.Write(bytes.FromInt64(info.recvTime))
	return buf.Bytes()
}
