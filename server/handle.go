package server

import (
	"github.com/OhYee/ait/message"
	"github.com/OhYee/goutils/bytes"
	"github.com/OhYee/rainbow/errors"
	"net"
)

func messageHandle(conn net.Conn) (repBytes []byte, repErr error, err error) {
	connClosed := false
	for !connClosed {
		message, err := msg.NewMessageFromBytes(conn)
		if err != nil {
			break
		}
		Log("Message: %+v %+v", message, message.Body)

		err = message.Handle(
			func(req *msg.Request) (err error) {
				return requestHandle(req, conn)
			},
			func(rep *msg.Response) (err error) {
				repBytes, repErr, err = responseHandle(rep)
				return
			},
			func(e *msg.Error) (err error) {
				Err(errors.New(e.Err))
				return nil
			},
			func(close *msg.Close) (err error) {
				connClosed = true
				return
			},
		)
	}
	return
}

func requestHandle(req *msg.Request, conn net.Conn) (err error) {
	api, err := GetAPI(req.APIName)
	if err != nil {
		return
	}
	Log("Request handle: %+v", req)

	rep, err := CallAPI(api, req.Arguments)
	Debug("Call API %s with request %s, return %+v %+v", api, req, rep, err)

	var response *msg.Message
	if err != nil {
		response = msg.NewResponse(false, bytes.FromStringWithLength32(err.Error())).ToMessage()
	} else {
		response = msg.NewResponse(true, rep).ToMessage()
	}
	Debug("Write %v", response.ToBytes())
	conn.Write(response.ToBytes())
	conn.Write(msg.NewClose().ToMessage().ToBytes())
	return
}

func responseHandle(rep *msg.Response) (reponse []byte, errMsg error, err error) {
	if rep.Success {
		reponse = rep.Arguments
		errMsg = nil
	} else {
		var errStr string
		errStr, err = bytes.ReadStringWithLength32(bytes.NewBuffer(rep.Arguments...))
		reponse = []byte{}
		errMsg = errors.New(errStr)
	}
	return
}
