package server

import (
	"github.com/xtaci/kcp-go"
	"time"
)

// Send data to addr
func Send(addr string, data []byte) (repBytes []byte, repErr error, err error) {
	conn, err := kcp.Dial(addr)
	if err != nil {
		return
	}
	conn.SetDeadline(time.Now().Add(time.Second * time.Duration(connDeadTime)))

	conn.Write(data)
	Log("Write %v", data)
	repBytes, repErr, err = messageHandle(conn)
	return
}
