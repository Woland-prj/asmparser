package commandsservice

import "fmt"

type SbisCmd struct {
	pMask uint16
	bMask uint16
}

func NewSbisCmd(pMask, bMask uint16) *SbicCmd {
	return &SbicCmd{pMask, bMask}
}

// sbis P,b - 1001 1011 PPPP Pbbb
func (c *SbisCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	p := (w & c.pMask) >> 3
	b := w & c.bMask
	return fmt.Sprintf("sbis 0x%x,%d", p, b)
}
