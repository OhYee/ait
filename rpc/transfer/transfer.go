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

var byteOrder = binary.BigEndian

const (
	// TypeUint8 uint8 type
	TypeUint8 uint8 = 1
	// TypeUint16 uint16 type
	TypeUint16 uint8 = 2
	// TypeUint32 uint32 type
	TypeUint32 uint8 = 3
	// TypeUint64 uint64 type
	TypeUint64 uint8 = 4
	// TypeInt8 int8 type
	TypeInt8 uint8 = 5
	// TypeInt16 int16 type
	TypeInt16 uint8 = 6
	// TypeInt32 int32 type
	TypeInt32 uint8 = 7
	// TypeInt64 int64 type
	TypeInt64 uint8 = 8
	// TypeString string type
	TypeString uint8 = 9
	// TypeBytes string type
	TypeBytes uint8 = 10
)

// FromString transfer from `uint8` to `[]byte`
func FromString(value string) (b []byte, err error) {
	var temp []byte
	buf := bytes.NewBuffer([]byte{TypeString})
	temp, err = FromUint32(uint32(len(value)))
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
func FromBytes(value []byte) (b []byte, err error) {
	var temp []byte
	buf := bytes.NewBuffer([]byte{TypeBytes})
	temp, err = FromUint32(uint32(len(value)))
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
