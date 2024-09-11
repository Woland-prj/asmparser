package domain

import "errors"

var (
	ErrInvalidHex          = errors.New("invalid hex string")
	ErrInvalidStringStruct = errors.New("invalid string struct")
	ErrFileNotRead         = errors.New("file not read")
	ErrFileNotOpen         = errors.New("file not open")
)
