package parse_usecase

import (
	"asmparser/internal/domain"
	"asmparser/internal/entities"
	"asmparser/internal/services/composeservice"
	"asmparser/internal/services/decodeservice"
	"asmparser/internal/services/disassemblyservice"
	"asmparser/internal/services/readservice"
	"errors"
)

type ParseUsecase struct {
	rs   *readservice.ReadService
	ds   *decodeservice.DecodeService
	cs   *composeservice.ComposeService
	dasm *disassemblyservice.DisassemblyService
}

func New(
	rs *readservice.ReadService,
	ds *decodeservice.DecodeService,
	cs *composeservice.ComposeService,
	dasm *disassemblyservice.DisassemblyService,
) *ParseUsecase {
	return &ParseUsecase{rs, ds, cs, dasm}
}

func (u *ParseUsecase) Do(strs []string) ([]string, error) {
	var hstrs []entities.HexString
	for _, str := range strs {
		hstr, err := u.ds.Decode(str)
		if err != nil {
			return nil, err
		}
		hstrs = append(hstrs, hstr)
	}

	var addrmaps []entities.AddressMap
	for _, hstr := range hstrs {
		addrmp, err := u.cs.Structurize(hstr)
		if err != nil {
			if errors.Is(err, domain.ErrEOF) {
				break
			}
			return nil, err
		}
		addrmaps = append(addrmaps, addrmp)
	}
	progmp := u.cs.Compose(addrmaps)

	progStrs, err := u.dasm.Disassemble(progmp)
	if err != nil {
		return nil, err
	}

	return progStrs, nil
}
