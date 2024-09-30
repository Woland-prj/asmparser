package commandsservice

import "fmt"

type SubiCmd struct {
	dMask uint16
	kMask uint16
}

func NewSubiCmd(dMask, kMask uint16) *SubiCmd {
	return &SubiCmd{dMask, kMask}
}

// subi Rd,K - 0101 KKKK dddd KKKK
func (c *SubiCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	k := ((w & c.kMask) >> 4) | (w & 0xF)
	d := ((w & c.dMask) >> 4) + 0x10
	return fmt.Sprintf("subi R%d,0x%x", d, k)
}
