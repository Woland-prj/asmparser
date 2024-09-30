package commandsservice

import "fmt"

type RjmpCmd struct {
	kMask uint16
}

func NewRjmpCmd(kMask uint16) *RjmpCmd {
	return &RjmpCmd{kMask}
}

// rjmp k - 1100 kkkk kkkk kkkk
func (c *RjmpCmd) GetMnemonic(cmd []uint16) string {
	w := cmd[0]
	k := w & c.kMask
	sign := '+'
	if k&(1<<6) != 0 {
		k ^= (1 << 12) - 1
		k += 1
		sign = '-'
	}
	return fmt.Sprintf("rjmp .%c%d", sign, k<<1)
}
