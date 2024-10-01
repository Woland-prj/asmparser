package main

import (
	"asmparser/internal/controller/cli"
	"asmparser/internal/services/composeservice"
	"asmparser/internal/services/decodeservice"
	"asmparser/internal/services/disassemblyservice"
	"asmparser/internal/services/readservice"
	"asmparser/internal/usecase/parse_usecase"
	"asmparser/internal/usecase/parsefile_usecase"
	"asmparser/internal/usecase/parsestdin_usecase"
)

func main() {
	// Init services
	rs := readservice.New()
	ds := decodeservice.New()
	cs := composeservice.New()
	dasm := disassemblyservice.New()

	// Init usecases
	cl := initCollector(rs, ds, cs, dasm)

	// Init CLI
	c := cli.New()
	c.Configure(cl)
	c.Serve()
}

func initCollector(
	rs *readservice.ReadService,
	ds *decodeservice.DecodeService,
	cs *composeservice.ComposeService,
	dasm *disassemblyservice.DisassemblyService,
) *cli.UsecaseCollector {
	pu := parse_usecase.New(rs, ds, cs, dasm)
	return &cli.UsecaseCollector{
		ParseFileUsecase:  parsefile_usecase.New(rs, pu),
		ParseStdinUsecase: parsestdin_usecase.New(rs, pu),
		ParseUsecase:      pu,
	}
}
