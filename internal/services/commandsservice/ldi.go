package commandsservice

import "fmt"

type LdiCmd struct {
	dMask uint16
	kMask uint16
}

func NewLdiCmd(dMask, kMask uint16) *LdiCmd {
	return &LdiCmd{dMask, kMask}
}

// ldi Rd,K - 1110 KKKK dddd KKKK
func (c *LdiCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	k := ((w & c.kMask) >> 4) | (w & 0xF)
	d := (w & c.dMask) >> 4
	return fmt.Sprintf("ldi R%d,0x%x", d, k)
}
