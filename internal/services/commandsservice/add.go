package commandsservice

import "fmt"

type AddCmd struct {
	dMask uint16
	rMask uint16
}

func NewAddCmd(dMask, rMask uint16) *AddCmd {
	return &AddCmd{dMask, rMask}
}

// add Rd,Rr - 0000 11rd dddd rrrr
func (c *AddCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	d := (w & c.dMask) >> 4
	preR := w & c.rMask
	r := (preR >> 5) | (preR & 0xF)
	return fmt.Sprintf("add R%d,R%d", d, r)
}
