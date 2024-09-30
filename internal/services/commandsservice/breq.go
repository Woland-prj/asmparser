package commandsservice

import "fmt"

type BreqCmd struct {
	kMask uint16
}

func NewBreqCmd(kMask uint16) *BreqCmd {
	return &BreqCmd{kMask}
}

// breq k - 1111 00kk kkkk k001
func (c *BreqCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	k := (w & c.kMask) >> 3
	sign := '+'
	if k&(1<<6) != 0 {
		k ^= (1 << 7) - 1
		k += 1
		sign = '-'
	}
	return fmt.Sprintf("breq .%c%d", sign, k<<1)
}
