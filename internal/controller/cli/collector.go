package cli

import (
	"asmparser/internal/usecase/parse_usecase"
	"asmparser/internal/usecase/parsefile_usecase"
	"asmparser/internal/usecase/parsestdin_usecase"
)

type UsecaseCollector struct {
	ParseFileUsecase  *parsefile_usecase.ParseFileUsecase
	ParseStdinUsecase *parsestdin_usecase.ParseStdinUsecase
	ParseUsecase      *parse_usecase.ParseUsecase
}
