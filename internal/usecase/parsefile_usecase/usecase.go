package parsefile_usecase

import (
	"asmparser/internal/services/readservice"
	"asmparser/internal/usecase/parse_usecase"
)

type ParseFileUsecase struct {
	rs *readservice.ReadService
	pu *parse_usecase.ParseUsecase
}

func New(
	rs *readservice.ReadService,
	pu *parse_usecase.ParseUsecase,
) *ParseFileUsecase {
	return &ParseFileUsecase{rs, pu}
}

func (u *ParseFileUsecase) Do(filePath string) ([]string, error) {
	u.rs.ReadFile(filePath)
	strs, err := u.rs.GetData()
	if err != nil {
		return nil, err
	}

	return u.pu.Do(strs)
}
