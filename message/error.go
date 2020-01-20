package msg

import (
	"github.com/OhYee/goutils/bytes"
	"github.com/OhYee/rainbow/errors"
	"io"
)

// Error message body
type Error struct {
	Err string
}

// NewError initial a Error
func NewError(err error) *Error {
	return &Error{
		Err: err.Error(),
	}
}

// NewErrorFromBytes initial a Error from []byte
func NewErrorFromBytes(r io.Reader) (e *Error, err error) {
	s, err := bytes.ReadStringWithLength32(r)
	if err != nil {
		return
	}
	e = NewError(errors.New(s))
	return
}

// ToBytes transfer Error to []byte
func (e *Error) ToBytes() []byte {
	return bytes.FromStringWithLength32(e.Err)
}

// ToMessage initial a Error message
func (e *Error) ToMessage() *Message {
	return NewMessage(MessageTypeError, e)
}
