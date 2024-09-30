package commandsservice

import "fmt"

type CbiCmd struct {
	pMask uint16
	bMask uint16
}

func NewCbiCmd(pMask, bMask uint16) *CbiCmd {
	return &CbiCmd{pMask, bMask}
}

// Cbi P,b - 1001 1000 PPPP Pbbb
func (c *CbiCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	p := (w & c.pMask) >> 3
	b := w & c.bMask
	return fmt.Sprintf("cbi 0x%x,%d", p, b)
}
