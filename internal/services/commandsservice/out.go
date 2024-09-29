package commandsservice

import "fmt"

type OutCmd struct {
	pMask uint16
	rMask uint16
}

func NewOutCmd(pMask, rMask uint16) *OutCmd {
	return &OutCmd{pMask, rMask}
}

// out P,R - 1011 1PPr rrrr PPPP
func (c *OutCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	p := ((w & c.pMask) >> 5) | (w & 0xF)
	r := (w & c.rMask) >> 4
	return fmt.Sprintf("out 0x%x,R%d", p, r)
}
