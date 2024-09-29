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

func (s *ComposeService) Structurize(str entities.HexString) (entities.AdrressMap, error) {
	if int(str.Len) != len(str.Data)*2 {
		return nil, domain.ErrWithAddr(domain.ErrLenNotMatch, str.Addr)
	}

	if str.FType != 0x00 && str.FType != 0x01 {
		return nil, domain.ErrWithAddr(domain.ErrInconsistenType, str.Addr)
	}

	if str.FType == 0x01 {
		return nil, domain.ErrEOF
	}

	var addrMap entities.AdrressMap
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
			nextAddr += 4
		} else {
			nextAddr += 2
		}

		addrMap[nextAddr] = cmd
	}

	return addrMap, nil
}

func (s *ComposeService) Compose([]entities.CommandMap) (entities.CommandMap, error) {
	return nil, nil
}
