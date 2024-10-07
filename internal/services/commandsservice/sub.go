package commandsservice

import "fmt"

type SubCmd struct {
	dMask uint16
	rMask uint16
}

func NewSubCmd(dMask, rMask uint16) *SubCmd {
	return &SubCmd{dMask, rMask}
}

// sub Rd,Rr - 0001 10rd dddd rrrr
func (c *SubCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	d := (w & c.dMask) >> 4
	preR := w & c.rMask
	r := (preR >> 5) | (preR & 0xF)
	return fmt.Sprintf("sub R%d,R%d", d, r)
}
