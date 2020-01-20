package server

import (
	"github.com/OhYee/rainbow/errors"
	"sync"
)

// Server Name System - SNS
// Singleton Pattern

var (
	serverMap      = make(map[string]Info)
	serverMapMutex = new(sync.Mutex)
)

// GetServerInfo get information of <serverName>
func GetServerInfo(serverName string) (info Info, err error) {
	serverMapMutex.Lock()
	defer serverMapMutex.Unlock()

	info, exist := serverMap[serverName]
	if !exist {
		err = errors.New("Can not get information of server %v", serverName)
	}

	return
}

// SetServerInfo set information of <serverName>
func SetServerInfo(info Info) {
	serverMapMutex.Lock()
	defer serverMapMutex.Unlock()

	serverMap[info.name] = info
}
