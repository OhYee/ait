package transfer

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/OhYee/goutil"
	"io"
)

//go:generate gcg int.json int.go
//go:generate gcg from.json from.go
//go:generate gcg to.json to.go

type any = interface{}

// Object the struct could be transfer
type Object interface {
	ToBytes() (b []byte)
}

var byteOrder = binary.BigEndian

const (
	_ = iota
	// TypeUint8 uint8 type
	TypeUint8 uint8 = iota
	// TypeUint16 uint16 type
	TypeUint16
	// TypeUint32 uint32 type
	TypeUint32
	// TypeUint64 uint64 type
	TypeUint64
	// TypeInt8 int8 type
	TypeInt8
	// TypeInt16 int16 type
	TypeInt16
	// TypeInt32 int32 type
	TypeInt32
	// TypeInt64 int64 type
	TypeInt64
	// TypeFloat32 float32 type
	TypeFloat32
	// TypeFloat64 float64 type
	TypeFloat64
	// TypeString string type
	TypeString
	// TypeBytes string type
	TypeBytes
	// TypeObject object type
	TypeObject
)

// FromString transfer from `uint8` to `[]byte`
func FromString(value string) (b []byte) {
	var temp []byte
	buf := bytes.NewBuffer([]byte{TypeString})
	temp = FromUint32(uint32(len(value)))
	buf.Write(temp[1:])
	buf.WriteString(value)
	b = buf.Bytes()
	return
}

// ToString transfer from `[]byte` to `string`
func ToString(r io.Reader) (value string, err error) {
	var t []byte
	t, err = goutil.ReadNBytes(r, 1)
	if t[0] != TypeString {
		err = fmt.Errorf("Transfer error: want %d(%T), got %d", TypeString, value, t[0])
	}
	value, err = toString(r)
	return
}

func toString(r io.Reader) (value string, err error) {
	var data []byte
	length, err := toUint32(r)
	if err != nil {
		return
	}
	data, err = goutil.ReadNBytes(r, int(length))
	if err != nil {
		return
	}
	value = string(data)
	return
}

// FromBytes transfer from `uint8` to `[]byte`
func FromBytes(value []byte) (b []byte) {
	var temp []byte
	buf := bytes.NewBuffer([]byte{TypeBytes})
	temp = FromUint32(uint32(len(value)))
	buf.Write(temp[1:])
	buf.Write(value)
	b = buf.Bytes()
	return
}

// ToBytes transfer from `[]byte` to `[]byte`
func ToBytes(r io.Reader) (value []byte, err error) {
	var t []byte
	t, err = goutil.ReadNBytes(r, 1)
	if t[0] != TypeBytes {
		err = fmt.Errorf("Transfer error: want %d(%T), got %d", TypeBytes, value, t[0])
	}
	value, err = toBytes(r)
	return
}

func toBytes(r io.Reader) (value []byte, err error) {
	length, err := toUint32(r)
	if err != nil {
		return
	}
	value, err = goutil.ReadNBytes(r, int(length))
	if err != nil {
		return
	}
	return
}

// FromObject transfer from `uint8` to `[]byte`
func FromObject(value Object) (b []byte) {
	var content, lengthBytes []byte
	buf := bytes.NewBuffer([]byte{TypeObject})
	content = value.ToBytes()

	lengthBytes = FromUint32(uint32(len(content)))
	buf.Write(lengthBytes[1:])
	buf.Write(content)
	b = buf.Bytes()
	return
}

// ToObject transfer from `[]byte` to `[]byte`
func ToObject(r io.Reader) (value []byte, err error) {
	var t []byte
	t, err = goutil.ReadNBytes(r, 1)
	if t[0] != TypeObject {
		err = fmt.Errorf("Transfer error: want %d(%T), got %d", TypeObject, value, t[0])
	}
	value, err = toObject(r)
	return
}

func toObject(r io.Reader) (value []byte, err error) {
	length, err := toUint32(r)
	if err != nil {
		return
	}
	value, err = goutil.ReadNBytes(r, int(length))
	if err != nil {
		return
	}
	return
}
