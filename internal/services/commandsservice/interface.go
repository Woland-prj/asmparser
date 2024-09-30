package commandsservice

import "asmparser/internal/entities"

type CommandService struct {
	cmdMap entities.CommandMap
}

func New() *CommandService {
	return &CommandService{
		cmdMap: entities.CommandMap{
			// 16-bit commands
			// nop - 0000 0000 0000 0000
			0x0000: NewNopCmd(),
			// cli - 1001 0100 1111 1000
			0x94f8: NewCliCmd(),
			// eor Rd,Rr - 0010 01rd dddd rrrr
			0x2400: NewEorCmd(0x01F0, 0x020F),
			// out P,R - 1011 1PPr rrrr PPPP
			0xb800: NewOutCmd(0x060F, 0x01F0),
			// ldi Rd,K - 1110 KKKK dddd KKKK
			0xe000: NewLdiCmd(0x00F0, 0x0F0F),
			// sbi P,b - 1001 1010 PPPP Pbbb
			0x9a00: NewSbiCmd(0x00F8, 0x0007),
			// cbi P,b - 1001 1000 PPPP Pbbb
			0x9800: NewCbiCmd(0x00F8, 0x0007),
			// breq k - 1111 00kk kkkk k001
			0xf001: NewBreqCmd(0x03F8),
			// brne k - 1111 01kk kkkk k001
			0xf401: NewBrneCmd(0x03F8),
			// rjmp k - 1100 kkkk kkkk kkkk
			0xc000: NewRjmpCmd(0x0FFF),
			// subi Rd,K - 0101 KKKK dddd KKKK
			0x5000: NewSubiCmd(0x00F0, 0x0F0F),
			// sbci Rd,K - 0100 KKKK dddd KKKK
			0x4000: NewSbciCmd(0x00F0, 0x0F0F),

			// 32-bit commands
			// jmp k - 1001 010k kkkk 110k kkkk kkkk kkkk kkkk
			0x940c: NewJmpCmd(0x01F1),
			// call k - 1001 010k kkkk 111k kkkk kkkk kkkk kkkk
			0x940e: NewCallCmd(0x01F1),
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
