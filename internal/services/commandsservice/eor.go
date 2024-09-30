package commandsservice

import "fmt"

type EorCmd struct {
	dMask uint16
	rMask uint16
}

func NewEorCmd(dMask, rMask uint16) *EorCmd {
	return &EorCmd{dMask, rMask}
}

// eor Rd,Rr - 0010 01rd dddd rrrr
func (c *EorCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	d := (w & c.dMask) >> 4
	preR := w & c.rMask
	r := (preR >> 5) | (preR & 0xF)
	return fmt.Sprintf("eor R%d,R%d", d, r)
}
