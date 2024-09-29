package domain

// 1111 1110 0000 1110 (jmp, call)
// 1111 1110 0000 1111 (sts, lds)
var MasksX32 = []uint16{0xfe0e, 0xfe0f}

var MasksX16 = []uint16{
	0xfc00, //eor
	0xf000, //ldi, rjmp, subi, subci
	0xf800, //out, in
	0xff00, //sbi, cbi
	0xfc07, //breq, brne, br....
	0xffff, //nop
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
