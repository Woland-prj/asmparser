package commandsservice

import "fmt"

type SbicCmd struct {
	pMask uint16
	bMask uint16
}

func NewSbicCmd(pMask, bMask uint16) *SbicCmd {
	return &SbicCmd{pMask, bMask}
}

// sbic P,b - 1001 1001 PPPP Pbbb
func (c *SbicCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	p := (w & c.pMask) >> 3
	b := w & c.bMask
	return fmt.Sprintf("sbic 0x%x,%d", p, b)
}
