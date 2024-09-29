package composeservice

import (
	"asmparser/internal/domain"
	"asmparser/internal/entities"
	"slices"
)

type ComposeService struct{}

func New() *ComposeService {
	return &ComposeService{}
}

func (s *ComposeService) Structurize(str entities.HexString) (entities.AddressMap, error) {
	if int(str.Len) != len(str.Data)*2 {
		return nil, domain.ErrWithAddr(domain.ErrLenNotMatch, str.Addr)
	}

	if str.FType != 0x00 && str.FType != 0x01 {
		return nil, domain.ErrWithAddr(domain.ErrInconsistenType, str.Addr)
	}

	if str.FType == 0x01 {
		return nil, domain.ErrEOF
	}

	addrMap := make(entities.AddressMap)
	isX32 := false
	nextAddr := str.Addr
	for i := 0; i < len(str.Data); i++ {
		if isX32 {
			isX32 = false
			continue
		}

		for _, mask := range domain.MasksX32 {
			if slices.Contains(domain.OpX32, str.Data[i]&mask) {
				isX32 = true
				break
			}
		}

		cmd := make([]uint16, 2)
		cmd[0] = str.Data[i]

		if isX32 {
			cmd[1] = str.Data[i+1]
			addrMap[nextAddr] = cmd
			nextAddr += 4
		} else {
			addrMap[nextAddr] = cmd
			nextAddr += 2
		}
	}

	return addrMap, nil
}

func (s *ComposeService) Compose(ams []entities.AddressMap) entities.AddressMap {
	progMap := make(entities.AddressMap)
	for _, am := range ams {
		for addr, cmds := range am {
			progMap[addr] = cmds
		}
	}
	return progMap
}
