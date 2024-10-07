package commandsservice

import "fmt"

type SbcCmd struct {
	dMask uint16
	rMask uint16
}

func NewSbcCmd(dMask, rMask uint16) *SbcCmd {
	return &SbcCmd{dMask, rMask}
}

// sbc Rd,Rr - 0000 10rd dddd rrrr
func (c *SbcCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	d := (w & c.dMask) >> 4
	preR := w & c.rMask
	r := (preR >> 5) | (preR & 0xF)
	return fmt.Sprintf("sbc R%d,R%d", d, r)
}
