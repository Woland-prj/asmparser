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
			// out P,R - 1011 1PPr rrrr PPPP
			0xb800: NewOutCmd(0x060F, 0x01F0),
			// ldi Rd,K - 1110 KKKK dddd KKKK
			0xe000: NewLdiCmd(0x0F0F, 0x00F0),

			// 32-bit commands
			// jmp k - 1001 010k kkkk 110k kkkk kkkk kkkk kkkk
			0x940c: NewJmpCmd(0x01F1),
		},
	}
}

func (cs *CommandService) GetCmdMap() entities.CommandMap {
	return cs.cmdMap
}

// 0 0000
// 1 0001
// 2 0010
// 3 0011
// 4 0100
// 5 0101
// 6 0110
// 7 0111
// 8 1000
// 9 1001
// a 1010
// b 1011
// c 1100
// d 1101
// e 1110
// f 1111
