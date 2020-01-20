package server

import (
	"github.com/xtaci/kcp-go"
)

// Send data to addr
func Send(addr string, data []byte) (repBytes []byte, repErr error, err error) {
	conn, err := kcp.Dial(addr)
	if err != nil {
		return
	}

	conn.Write(data)
	repBytes, repErr, err = messageHandle(conn)
	return
}
