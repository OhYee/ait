package transfer

import (
	"encoding/binary"
)

//go:generate gcg int.json int.go
//go:generate gcg from.json from.go

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
)


