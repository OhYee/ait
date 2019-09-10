package transfer

import (
	"bytes"
	"testing"
)

type testCase struct {
	name string      //stirng
	b    []byte      // bytes
	v    interface{} // value
	tf   interface{} // to function
	ff   interface{} // from function
}

var testcases = []testCase{
	{
		name: "uint8",
		b:    []byte{TypeUint8, 0x2f},
		v:    uint8(0x2f),
		ff:   fromUint8,
	},
	{
		name: "uint16",
		b:    []byte{TypeUint16, 0x2f, 0x3f},
		v:    uint16(0x2f3f),
		ff:   fromUint8,
	},
	{
		name: "uint32",
		b:    []byte{TypeUint32, 0x2f, 0x3f, 0x4f, 0x5f},
		v:    uint32(0x2f3f4f5f),
		ff:   fromUint8,
	},
	{
		name: "uint64",
		b:    []byte{TypeUint64, 0x2f, 0x3f, 0x4f, 0x5f, 0x6f, 0x7f, 0x8f, 0x9f},
		v:    uint64(0x2f3f4f5f6f7f8f9f),
		ff:   fromUint8,
	},
	{
		name: "int8 positive",
		b:    []byte{TypeInt8, 0x1f},
		v:    int8(0x1f),
		ff:   fromUint8,
	},
	{
		name: "int8 negative",
		b:    []byte{TypeInt8, 0xe1},
		v:    -int8(0x1f),
		ff:   fromUint8,
	},
	{
		name: "int16 positive",
		b:    []byte{TypeInt16, 0x1f, 0x2f},
		v:    int16(0x1f2f),
		ff:   fromUint16,
	},
	{
		name: "int16 negative",
		b:    []byte{TypeInt16, 0xE0, 0xD1},
		v:    -int16(0x1f2f),
		ff:   fromUint16,
	},
	{
		name: "int32 positive",
		b:    []byte{TypeInt32, 0x1f, 0x2f, 0x3f, 0x4f},
		v:    int32(0x1f2f3f4f),
		ff:   fromUint32,
	},
	{
		name: "int32 negative",
		b:    []byte{TypeInt32, 0xE0, 0xD0, 0xC0, 0xB1},
		v:    -int32(0x1f2f3f4f),
		ff:   fromUint32,
	},
	{
		name: "int64 positive",
		b:    []byte{TypeInt64, 0x1f, 0x2f, 0x3f, 0x4f, 0x5f, 0x6f, 0x7f, 0x8f},
		v:    int64(0x1f2f3f4f5f6f7f8f),
		ff:   fromUint64,
	},
	{
		name: "int64 negative",
		b:    []byte{TypeInt64, 0xE0, 0xD0, 0xC0, 0xB0, 0xA0, 0x90, 0x80, 0x71},
		v:    -int64(0x1f2f3f4f5f6f7f8f),
		ff:   fromUint64,
	},
}

func TestTranasferFrom(t *testing.T) {
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			b, _ := fromValue(testcase.v)
			if !bytes.Equal(b, testcase.b) {
				t.Errorf("Except %v, got %v", testcase.b, b)
			}
		})
	}
}
