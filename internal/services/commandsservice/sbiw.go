package commandsservice

import "fmt"

type SbiwCmd struct {
	dMask uint16
	kMask uint16
}

func NewSbiwCmd(dMask, kMask uint16) *SbiwCmd {
	return &SbiwCmd{dMask, kMask}
}

// sbiw Rd,k - 1001 0111 kkdd kkkk
func (c *SbiwCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	d := (w & c.dMask) >> 4
	preK := w & c.kMask
	k := (preK >> 2) | (preK & 0xF)
	reg := 24 + d*2
	return fmt.Sprintf("sbiw R%d:R%d,0x%x", reg+1, reg, k)
}
