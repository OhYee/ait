package server

import (
	"encoding/json"
	"github.com/OhYee/ait/message"
	"reflect"
)

func getObject(data interface{}) (valueSlice []reflect.Value) {
	dataValue := reflect.ValueOf(data)
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
type RegisterType func(apiName string, f interface{}, req interface{}, rep interface{}) (err error)

// MakeRegister make a register to register function as api
func MakeRegister(serverName string) (register RegisterType) {
	return func(apiName string, f interface{}, req interface{}, rep interface{}) (err error) {
		var reqBytes []byte
		if reqBytes, err = json.Marshal(req); err != nil {
			return
		}
		RegisterAPI(apiName, f)

		reqMessage := msg.NewRequest(serverName, apiName, reqBytes)

		var info Info
		if info, err = GetServerInfo(serverName); err != nil {
			return
		}

		var repBytes []byte
		var repErr error
		if repBytes, repErr, err = Send(info.addr, reqMessage.ToBytes()); err != nil {
			return
		}
		if repErr == nil {
			err = json.Unmarshal(repBytes, rep)
		} else {
			err = repErr
		}
		return
	}
}

// CallFunction using request object and set reponse object
func CallFunction(f interface{}, req interface{}, rep interface{}) {
	in := getObject(req)
	out := reflect.ValueOf(f).Call(in)
	setObject(out, rep)
}

// CallAPI using []byte-request, and return []byte-request with error
func CallAPI(f interface{}, req []byte) (rep []byte, err error) {
	fValue := reflect.ValueOf(f)
	fType := reflect.TypeOf(f)

	request := reflect.New(fType.In(0)).Interface()
	if err = json.Unmarshal(req, &req); err != nil {
		return
	}

	in := getObject(request)
	out := fValue.Call(in)

	response := out[0].Interface()
	rep, err = json.Marshal(response)
	return
}
