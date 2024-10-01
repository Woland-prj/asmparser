package parsestdin_usecase

import (
	"asmparser/internal/services/readservice"
	"asmparser/internal/usecase/parse_usecase"
	"fmt"
)

type ParseStdinUsecase struct {
	rs *readservice.ReadService
	pu *parse_usecase.ParseUsecase
}

func New(
	rs *readservice.ReadService,
	pu *parse_usecase.ParseUsecase,
) *ParseStdinUsecase {
	return &ParseStdinUsecase{rs, pu}
}

func (u *ParseStdinUsecase) Do() ([]string, error) {
	fmt.Println("Enter Intel HEX strings separated by newlines:")
	u.rs.ReadStdin()
	strs, err := u.rs.GetData()
	if err != nil {
		return nil, err
	}

	return u.pu.Do(strs)
}
