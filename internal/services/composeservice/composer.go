package composeservice

import "asmparser/internal/entities"

type ComposeService struct{}

func New() *ComposeService {
	return &ComposeService{}
}

func (s *ComposeService) Structurize(hstr entities.HexString) (string, error) {
	return "", nil
}
