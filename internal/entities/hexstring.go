package entities

import "fmt"

type HexString struct {
	Len   byte
	Addr  uint16
	Data  []uint16
	FType byte
	Crc   byte
}

func (hstr HexString) String() string {
	return fmt.Sprint("Len: ") + fmt.Sprintf("%02x\n", hstr.Len) +
		fmt.Sprint("Addr: ") + fmt.Sprintf("%04x\n", hstr.Addr) +
		fmt.Sprint("FType: ") + fmt.Sprintf("%02x\n", hstr.FType) +
		fmt.Sprint("Data: ") + fmt.Sprintf("%04x\n", hstr.Data) +
		fmt.Sprint("Control sum: ") + fmt.Sprintf("%02x\n", hstr.Crc)
}
