package commandsservice

import "fmt"

type BrneCmd struct {
	kMask uint16
}

func NewBrneCmd(kMask uint16) *BrneCmd {
	return &BrneCmd{kMask}
}

// brne k - 1111 01kk kkkk k001
func (c *BrneCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	k := (w & c.kMask) >> 3
	sign := '+'
	if k&(1<<6) != 0 {
		k ^= (1 << 7) - 1
		k += 1
		sign = '-'
	}
	return fmt.Sprintf("brne .%c%d", sign, k<<1)
}
