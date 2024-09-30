package commandsservice

import "fmt"

type NopCmd struct{}

func NewNopCmd() *NopCmd {
	return &NopCmd{}
}

// nop - 0000 0000 0000 0000
func (c *NopCmd) GetMnemonic(cmd []uint16) string {
	return fmt.Sprintf("nop")
}
