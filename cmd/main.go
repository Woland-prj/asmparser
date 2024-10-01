package main

import (
	"asmparser/internal/domain"
	"asmparser/internal/entities"
	"asmparser/internal/services/composeservice"
	"asmparser/internal/services/decodeservice"
	"asmparser/internal/services/disassemblyservice"
	"asmparser/internal/services/readservice"
	"errors"
	"fmt"
	"log"
)

func main() {
	rs := readservice.New()
	ds := decodeservice.New()
	cs := composeservice.New()
	dasm := disassemblyservice.New()
	rs.ReadFile("test.txt")
	strs, err := rs.GetFileData()
	if err != nil {
		log.Fatal(err)
	}

	var hstrs []entities.HexString
	for _, str := range strs {
		hstr, err := ds.Decode(str)
		if err != nil {
			log.Fatal(err)
		}
		hstrs = append(hstrs, hstr)
	}

	var maps []entities.AddressMap
	for _, hstr := range hstrs {
		mp, err := cs.Structurize(hstr)
		if err != nil {
			if errors.Is(err, domain.ErrEOF) {
				break
			}
			log.Fatal(err)
		}
		maps = append(maps, mp)
	}
	mp := cs.Compose(maps)

	progStrs, err := dasm.Disassemble(mp)
	if err != nil {
		log.Fatal(err)
	}
	for _, progStr := range progStrs {
		fmt.Println(progStr)
	}
}
