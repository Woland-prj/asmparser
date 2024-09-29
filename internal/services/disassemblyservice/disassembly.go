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
		for _, mask := range domain.MasksX32 {
			opcode := cmdHex[0] & mask
			cmd, ok := cmdMap[opcode]
			if !ok {
				continue
			}
			asmStr := fmt.Sprintf("%2x: %s", addr, cmd.GetMnemonic(cmdHex))
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
