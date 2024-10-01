package disassemblyservice

import (
	"asmparser/internal/domain"
	"asmparser/internal/entities"
	"asmparser/internal/services/commandsservice"
	"fmt"
	"sort"
)

type DisassemblyService struct {
	cs *commandsservice.CommandService
}

func New() *DisassemblyService {
	return &DisassemblyService{
		cs: commandsservice.New(),
	}
}

func (ds *DisassemblyService) Disassemble(progAddrMap entities.AddressMap) ([]string, error) {
	cmdMap := ds.cs.GetCmdMap()
	var resProg []string
	keys := make([]uint16, 0)
	for k := range progAddrMap {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, addr := range keys {
		cmdHex := progAddrMap[addr]
		found := false
		for _, mask := range domain.Masks {
			opcode := cmdHex[0] & mask
			cmd, ok := cmdMap[opcode]
			if !ok {
				continue
			}
			asmStr := fmt.Sprintf("%02x:  %s  %s", addr, ds.splitCmd(cmdHex), cmd.GetMnemonic(cmdHex))
			resProg = append(resProg, asmStr)
			found = true
			break
		}

		if !found {
			return nil, domain.ErrWithAddr(domain.ErrUnknownOpcode, addr)
		}
	}

	return resProg, nil
}

func (ds *DisassemblyService) splitCmd(cmd []uint16) string {
	res := fmt.Sprintf("%02x %02x", cmd[0]&0x00ff, (cmd[0]&0xff00)>>8)
	if (len(cmd) == 2) && (cmd[1] != 0) {
		res += fmt.Sprintf(" %02x %02x", cmd[1]&0x00ff, (cmd[1]&0xff00)>>8)
	}
	return res
}
