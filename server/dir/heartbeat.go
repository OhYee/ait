package dir

import (
	"github.com/OhYee/ait/server"
	"time"
)

var register = server.MakeRegister("directory")

// HeartBeatRequest request of heartbeat
type HeartBeatRequest struct {
	ServerInfo server.Info
}

// HeartBeatResponse response of heartbeat
type HeartBeatResponse struct {
	ServerNum  int
	ServerList []server.Info
}

func heartBeatHandle(serverInfo server.Info) (serverNum int, serverList []server.Info) {
	serverInfo.RecvTime = time.Now().Unix()
	server.SetServerInfo(serverInfo)
	serverList = server.GetServerInfoList()
	serverNum = len(serverList)
	return
}

// HeartBeat send heartbeat and recvive servers info
var HeartBeat = func() func(req HeartBeatRequest) (rep HeartBeatResponse, err error) {
	caller := register("heartbeat", heartBeatHandle, HeartBeatRequest{}, HeartBeatResponse{})
	return func(req HeartBeatRequest) (rep HeartBeatResponse, err error) {
		err = caller(req, &rep)
		return
	}
}()

// StartHeartBeatThread send heartbeat and get server list
func StartHeartBeatThread() {
	for {
		info := server.NewInfo(server.GetServerName(), server.GetServerAddr(), time.Now().Unix(), 0)
		server.Debug("Send HeartBeat %+v", info)
		req, err := HeartBeat(HeartBeatRequest{info})
		if err != nil {
			server.Err(err)
		} else {
			for _, serverInfo := range req.ServerList {
				server.SetServerInfo(serverInfo)
			}
		}
		<-time.After(time.Second * 10)
	}
}
