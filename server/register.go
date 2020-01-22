package server

import (
	"encoding/json"
	"github.com/OhYee/ait/message"
	"github.com/OhYee/rainbow/errors"
	"reflect"
)

func getObject(data interface{}) (valueSlice []reflect.Value) {
	Debug("%T %+v", data, data)
	dataValue := reflect.ValueOf(data).Elem()
	inputNum := dataValue.NumField()
	valueSlice = make([]reflect.Value, inputNum)
	for i := 0; i < inputNum; i++ {
		valueSlice[i] = dataValue.Field(i)
	}
	return
}

func setObject(valueSlice []reflect.Value, data interface{}) {
	dataValue := reflect.ValueOf(data).Elem()
	outputNum := dataValue.NumField()
	for i := 0; i < outputNum; i++ {
		dataValue.Field(i).Set(valueSlice[i])
	}
	return
}

// RegisterType register type
type RegisterType func(apiName string, f interface{}, reqType interface{}, repType interface{}) (caller CallerType)

// CallerType api caller function
type CallerType func(req interface{}, rep interface{}) (err error)

// MakeRegister make a register to register function as api
func MakeRegister(_serverName string) (register RegisterType) {
	serverName = _serverName
	register = func(apiName string, f interface{}, reqType interface{}, repType interface{}) (caller CallerType) {
		RegisterAPI(apiName, f, reqType, repType)
		caller = func(req interface{}, rep interface{}) (err error) {
			var reqBytes []byte
			if reqBytes, err = json.Marshal(req); err != nil {
				return
			}

			reqMessage := msg.NewRequest(_serverName, apiName, reqBytes).ToMessage()

			var info Info
			if info, err = GetServerInfo(_serverName); err != nil {
				return
			}

			var repBytes []byte
			var repErr error
			if repBytes, repErr, err = Send(info.Addr, reqMessage.ToBytes()); err != nil {
				return
			}
			if repErr == nil {
				err = json.Unmarshal(repBytes, rep)
			} else {
				err = repErr
			}
			return
		}
		return
	}
	return
}

// CallFunction using request object and set reponse object
func CallFunction(f interface{}, req interface{}, rep interface{}) {
	in := getObject(req)
	out := reflect.ValueOf(f).Call(in)
	setObject(out, rep)
}

// CallAPI using []byte-request, and return []byte-request with error
func CallAPI(api API, req []byte) (rep []byte, err error) {
	Debug("Call API %+v with request %s", api, req)

	request := reflect.New(reflect.TypeOf(api.request)).Interface()

	if err = json.Unmarshal(req, &request); err != nil {
		err = errors.NewErr(err)
		return
	}
	Debug("%+v %T %s", request, request, reflect.ValueOf(request))

	in := getObject(request)
	out := reflect.ValueOf(api.function).Call(in)

	Debug("%+v", out[0])

	response := reflect.New(reflect.TypeOf(api.response)).Interface()
	setObject(out, response)

	Debug("%+v", response)

	rep, err = json.Marshal(response)
	Debug("%+v %+v", rep, err)
	if err != nil {
		err = errors.NewErr(err)
	}

	return
}
