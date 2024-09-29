package commandsservice

import "fmt"

type JmpCmd struct {
	kMask uint16
}

func NewJmpCmd(kMask uint16) *JmpCmd {
	return &JmpCmd{kMask}
}

// jmp k - 1001 010k kkkk 110k kkkk kkkk kkkk kkkk
func (c *JmpCmd) GetMnemonic(cmd []uint16) string {
	w1 := cmd[0]
	w2 := cmd[1]
	k := (uint32(((w1&c.kMask)>>3)|(w1&0x1)) << 16) | uint32(w2)
	return fmt.Sprintf("jmp 0x%x", k<<1)
}
