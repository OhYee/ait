package server

import (
	"github.com/OhYee/rainbow/errors"
	"sync"
)

// Singleton Pattern
var (
	mapMutex = new(sync.Mutex)
	funcMap  = make(map[string]interface{})
	apiMap   = make(map[string]interface{})
)

// RegisterAPI register api as server api
func RegisterAPI(apiName string, f interface{}) {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	apiMap[apiName] = f
}

// GetAPI get api (is not exist, return nil)
func GetAPI(apiName string) (api interface{}, err error) {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	api, exist := apiMap[apiName]
	if !exist {
		err = errors.New("No api named %s", apiName)
	}
	return
}
