package commandsservice

import "asmparser/internal/entities"

type CommandService struct {
	cmdMap entities.CommandMap
}

func New() *CommandService {
	return &CommandService{
		cmdMap: entities.CommandMap{
			// 16-bit commands
			// eor Rd,Rr - 0010 01rd dddd rrrr
			0x2400: NewEorCmd(0x020F, 0x01F0),

			// 32-bit commands
			// jmp k - 1001 010k kkkk 110k kkkk kkkk kkkk kkkk
			0x940c: NewJmpCmd(0x01F1),
		},
	}
}

func (cs *CommandService) GetCmdMap() entities.CommandMap {
	return cs.cmdMap
}
