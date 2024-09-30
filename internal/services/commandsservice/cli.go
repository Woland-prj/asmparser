package commandsservice

import "fmt"

type CliCmd struct{}

func NewCliCmd() *CliCmd {
	return &CliCmd{}
}

// cli - 1001 0100 1111 1000
func (c *CliCmd) GetMnemonic(cmd []uint16) string {
	return fmt.Sprintf("cli")
}
