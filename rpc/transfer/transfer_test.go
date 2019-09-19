package transfer

import (
	"bytes"
	"github.com/OhYee/goutil"
	"reflect"
	"testing"
)

type testCase struct {
	name    string      //stirng
	b       []byte      // bytes
	v       interface{} // value
	tf      interface{} // to function
	wantErr bool
}

var testcases = []testCase{
	{
		name:    "uint8",
		b:       []byte{TypeUint8, 0x2f},
		v:       uint8(0x2f),
		tf:      ToUint8,
		wantErr: false,
	},
	{
		name:    "uint16",
		b:       []byte{TypeUint16, 0x2f, 0x3f},
		v:       uint16(0x2f3f),
		tf:      ToUint16,
		wantErr: false,
	},
	{
		name:    "uint32",
		b:       []byte{TypeUint32, 0x2f, 0x3f, 0x4f, 0x5f},
		v:       uint32(0x2f3f4f5f),
		tf:      ToUint32,
		wantErr: false,
	},
	{
		name:    "uint64",
		b:       []byte{TypeUint64, 0x2f, 0x3f, 0x4f, 0x5f, 0x6f, 0x7f, 0x8f, 0x9f},
		v:       uint64(0x2f3f4f5f6f7f8f9f),
		tf:      ToUint64,
		wantErr: false,
	},
	{
		name:    "int8 positive",
		b:       []byte{TypeInt8, 0x1f},
		v:       int8(0x1f),
		tf:      ToInt8,
		wantErr: false,
	},
	{
		name:    "int8 negative",
		b:       []byte{TypeInt8, 0xe1},
		v:       -int8(0x1f),
		tf:      ToInt8,
		wantErr: false,
	},
	{
		name:    "int16 positive",
		b:       []byte{TypeInt16, 0x1f, 0x2f},
		v:       int16(0x1f2f),
		tf:      ToInt16,
		wantErr: false,
	},
	{
		name:    "int16 negative",
		b:       []byte{TypeInt16, 0xE0, 0xD1},
		v:       -int16(0x1f2f),
		tf:      ToInt16,
		wantErr: false,
	},
	{
		name:    "int32 positive",
		b:       []byte{TypeInt32, 0x1f, 0x2f, 0x3f, 0x4f},
		v:       int32(0x1f2f3f4f),
		tf:      ToInt32,
		wantErr: false,
	},
	{
		name:    "int32 negative",
		b:       []byte{TypeInt32, 0xE0, 0xD0, 0xC0, 0xB1},
		v:       -int32(0x1f2f3f4f),
		tf:      ToInt32,
		wantErr: false,
	},
	{
		name:    "int64 positive",
		b:       []byte{TypeInt64, 0x1f, 0x2f, 0x3f, 0x4f, 0x5f, 0x6f, 0x7f, 0x8f},
		v:       int64(0x1f2f3f4f5f6f7f8f),
		tf:      ToInt64,
		wantErr: false,
	},
	{
		name:    "int64 negative",
		b:       []byte{TypeInt64, 0xE0, 0xD0, 0xC0, 0xB0, 0xA0, 0x90, 0x80, 0x71},
		v:       -int64(0x1f2f3f4f5f6f7f8f),
		tf:      ToInt64,
		wantErr: false,
	},
	{
		name:    "string",
		b:       []byte{TypeString, 0x00, 0x00, 0x00, 0x08, 'A', 'b', 'C', 'd', '1', '2', '3', '!'},
		v:       "AbCd123!",
		tf:      ToString,
		wantErr: false,
	},
	{
		name:    "string error",
		b:       []byte{TypeString, 0x00, 0x00, 0x00, 0x08, 'A', 'b', 'C', 'd', '1', '2', '3'},
		v:       "AbCd123!",
		tf:      ToString,
		wantErr: true,
	},
	{
		name:    "bytes",
		b:       []byte{TypeBytes, 0x00, 0x00, 0x00, 0x02, 0x05, 0x04},
		v:       []byte{0x05, 0x04},
		tf:      ToBytes,
		wantErr: false,
	},
	{
		name:    "bytes",
		b:       []byte{TypeBytes, 0x00, 0x00, 0x00, 0x02, 0x05},
		v:       []byte{0x05, 0x04},
		tf:      ToBytes,
		wantErr: true,
	},
}

func TestTranasferFrom(t *testing.T) {
	t.Run("unknown type", func(t *testing.T) {
		b, err := FromValue(testCase{})
		if !bytes.Equal(b, []byte{}) || err == nil {
			t.Errorf("Except %v, got %v", []byte{}, b)
		}
	})
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			b, err := FromValue(testcase.v)
			if testcase.wantErr == false && (!bytes.Equal(b, testcase.b) || err != nil) {
				t.Errorf("Except %v, got %v", testcase.b, b)
			}
		})
	}
}

func TestTranasferTo(t *testing.T) {
	t.Run("unknown type", func(t *testing.T) {
		_, err := ToValue(bytes.NewBuffer([]byte{0xff}))
		if err == nil {
			t.Errorf("Except err, got nil")
		}
	})
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			value, err := ToValue(bytes.NewBuffer(testcase.b))
			if testcase.wantErr == false && (!goutil.Equal(value, testcase.v) || err != nil) {
				t.Errorf("Except %v, got %v", testcase.v, value)
			}
		})
		t.Run("To"+testcase.name, func(t *testing.T) {
			output := reflect.ValueOf(testcase.tf).Call([]reflect.Value{reflect.ValueOf(bytes.NewBuffer(testcase.b))})
			value := output[0].Interface()
			err := output[1].Interface()
			if testcase.wantErr == false && (!goutil.Equal(value, testcase.v) || err != nil) {
				t.Errorf("Except %v nil, got %v %v", testcase.v, value, err)
			}
		})
		t.Run("To"+testcase.name+" error", func(t *testing.T) {
			output := reflect.ValueOf(testcase.tf).Call([]reflect.Value{reflect.ValueOf(bytes.NewBuffer([]byte{0xff}))})
			err := output[1].Interface()
			if err == nil {
				t.Errorf("Except err, got nil")
			}
		})
	}
}
