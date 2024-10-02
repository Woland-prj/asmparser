package commandsservice

import "fmt"

type AdiwCmd struct {
	dMask uint16
	kMask uint16
}

func NewAdiwCmd(dMask, kMask uint16) *AdiwCmd {
	return &AdiwCmd{dMask, kMask}
}

// adiw Rd,k - 1001 0110 kkdd kkkk
func (c *AdiwCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	d := (w & c.dMask) >> 4
	preK := w & c.kMask
	k := (preK >> 2) | (preK & 0xF)
	return fmt.Sprintf("adiw R%d,0x%x", d, k)
}
