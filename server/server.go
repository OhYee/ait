package server

import (
	"github.com/OhYee/ait/message"
	"github.com/xtaci/kcp-go"
	"net"
	"time"
)

// Singleton Pattern
var (
	closeThread  = make([]chan bool, 0)
	connDeadTime = 5000
	serverAddr   = ""
	serverName   = ""
)

// Start server, blocking the process after returned
func Start(name string, addr string, threadNumber int) (err error) {
	serverAddr = addr
	serverName = name
	listener, err := kcp.Listen(addr)
	if err != nil {
		return
	}

	closeThread = make([]chan bool, threadNumber)
	for i := 0; i < threadNumber; i++ {
		closeThread[i] = make(chan bool, 1)
		go thread(i, listener)
	}
	return
}

func thread(threadID int, listener net.Listener) {
	Log("Thread %d start", threadID)
	defer Log("Thread %d stoped", threadID)

	close := false
	for !close {
		select {
		case close = <-closeThread[threadID]:
		default:
			conn, err := listener.Accept()
			if err != nil {
				Err(err)
				continue
			}
			conn.SetDeadline(time.Now().Add(time.Second * time.Duration(connDeadTime)))
			if handle(threadID, conn) != nil {
				Err(err)
				continue
			}
		}
	}
}

// Close server
func Close() {
	for i := range closeThread {
		closeThread[i] <- true
	}
}

func handle(threadID int, conn net.Conn) (err error) {
	Log("Thread %d receive a connection from %s", threadID, conn.RemoteAddr())
	// if has error, send a error message
	defer func() {
		if err != nil {
			conn.Write(msg.NewError(err).ToMessage().ToBytes())
		}
	}()
	// close connction, send close message
	defer conn.Write(msg.NewClose().ToMessage().ToBytes())
	defer Log("Thread %d close the connection with %s", threadID, conn.RemoteAddr())

	_, _, err = messageHandle(conn)
	return
}

func GetServerAddr() string {
	return serverAddr
}
func GetServerName() string {
	return serverName
}
