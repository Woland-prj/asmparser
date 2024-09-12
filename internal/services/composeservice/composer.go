package composeservice

import "asmparser/internal/entities"

type ComposeService struct{}

func NewComposeService() *ComposeService {
	return &ComposeService{}
}

func (s *ComposeService) Structurize(str entities.HexString) (entities.Addresmap, error) {
	return nil, nil
}

func (s *ComposeService) Compose([]entities.Addresmap) (entities.Addresmap, error) {
	return nil, nil
}
