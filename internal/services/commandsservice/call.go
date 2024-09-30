package commandsservice

import "fmt"

type CallCmd struct {
	kMask uint16
}

func NewCallCmd(kMask uint16) *CallCmd {
	return &CallCmd{kMask}
}

// call k - 1001 010k kkkk 111k kkkk kkkk kkkk kkkk
func (c *CallCmd) GetMnemonic(cmd []uint16) string {
	w1 := cmd[0]
	w2 := cmd[1]
	k := (uint32(((w1&c.kMask)>>3)|(w1&0x1)) << 16) | uint32(w2)
	return fmt.Sprintf("call 0x%x", k<<1)
}
