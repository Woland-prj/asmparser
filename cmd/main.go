package main

import (
	"asmparser/internal/entities"
	"asmparser/internal/services/decodeservice"
	"asmparser/internal/services/readservice"
	"fmt"
	"log"
)

func main() {
	rs := readservice.New()
	ds := decodeservice.New()
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

	for _, hstr := range hstrs {
		fmt.Println(hstr)
	}
}
