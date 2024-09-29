package entities

type Command interface {
	GetMnemonic(cmd []uint16) string
}

type CommandMap = map[uint16]Command
