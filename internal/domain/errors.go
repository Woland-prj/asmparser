package domain

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidHex          = errors.New("invalid hex string")
	ErrInvalidStringStruct = errors.New("invalid string struct")
	ErrLenNotMatch         = errors.New("count of package data bytes and len byte do not match")
	ErrInconsistenType     = errors.New("inconsistent package type")
	ErrFileNotRead         = errors.New("file not read")
	ErrFileNotOpen         = errors.New("file not open")

	ErrEOF = errors.New("end of file")
)

func ErrWithAddr(err error, addr uint16) error {
	return fmt.Errorf("%w. Adress: %04x", err, addr)
}
