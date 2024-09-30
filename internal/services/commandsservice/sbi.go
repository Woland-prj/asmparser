package commandsservice

import "fmt"

type SbiCmd struct {
	pMask uint16
	bMask uint16
}

func NewSbiCmd(pMask, bMask uint16) *SbiCmd {
	return &SbiCmd{pMask, bMask}
}

// sbi P,b - 1001 1010 PPPP Pbbb
func (c *SbiCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	p := (w & c.pMask) >> 3
	b := w & c.bMask
	return fmt.Sprintf("sbi 0x%x,%d", p, b)
}
