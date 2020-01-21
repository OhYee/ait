package server

import (
	"github.com/OhYee/rainbow/errors"
	"sync"
)

// Singleton Pattern
var (
	mapMutex = new(sync.Mutex)
	apiMap   = make(map[string]API)
)

// API for micro-server
type API struct {
	apiName  string
	function interface{}
	request  interface{}
	response interface{}
}

// RegisterAPI register api as server api
func RegisterAPI(apiName string, f interface{}, reqType interface{}, repType interface{}) {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	apiMap[apiName] = API{apiName, f, reqType, repType}
}

// GetAPI get api (is not exist, return nil)
func GetAPI(apiName string) (api API, err error) {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	api, exist := apiMap[apiName]
	if !exist {
		err = errors.New("No api named %s", apiName)
	}
	return
}
