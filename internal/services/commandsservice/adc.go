package commandsservice

import "fmt"

type AdcCmd struct {
	dMask uint16
	rMask uint16
}

func NewAdcCmd(dMask, rMask uint16) *AddCmd {
	return &AddCmd{dMask, rMask}
}

// adc Rd,Rr - 0001 11rd dddd rrrr
func (c *AddCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	d := (w & c.dMask) >> 4
	preR := w & c.rMask
	r := (preR >> 5) | (preR & 0xF)
	return fmt.Sprintf("adc R%d,R%d", d, r)
}
