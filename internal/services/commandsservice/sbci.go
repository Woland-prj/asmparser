package commandsservice

import "fmt"

type SbciCmd struct {
	dMask uint16
	kMask uint16
}

func NewSbciCmd(dMask, kMask uint16) *SbciCmd {
	return &SbciCmd{dMask, kMask}
}

// sbci Rd,K - 0100 KKKK dddd KKKK
func (c *SbciCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	k := ((w & c.kMask) >> 4) | (w & 0xF)
	d := ((w & c.dMask) >> 4) + 0x10
	return fmt.Sprintf("sbci R%d,0x%x", d, k)
}
