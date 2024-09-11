package decodeservice

import (
	"asmparser/internal/domain"
	"asmparser/internal/entities"
	"encoding/hex"
	"fmt"
)

type DecodeService struct{}

func New() *DecodeService {
	return &DecodeService{}
}

func (s *DecodeService) Decode(str string) (entities.HexString, error) {
	byteStr := str[1:]
	packageHex, err := hex.DecodeString(byteStr)

	if err != nil {
		return entities.HexString{},
			fmt.Errorf("DecodeService - Decode: %w. Reason: %w", domain.ErrInvalidHex, err)
	}

	if len(packageHex) < 5 {
		return entities.HexString{},
			fmt.Errorf("DecodeService - Decode: %w", domain.ErrInvalidStringStruct)
	}

	var hexStr entities.HexString
	hexStr.Len = packageHex[0]
	hexStr.Addr = uint16(packageHex[1])<<8 + uint16(packageHex[2])
	hexStr.FType = packageHex[3]
	hexStr.Crc = packageHex[len(packageHex)-1]

	dataBytes := packageHex[4 : len(packageHex)-1]

	if len(dataBytes) != int(hexStr.Len) || len(dataBytes)%2 != 0 {
		return entities.HexString{},
			fmt.Errorf("DecodeService - Decode: %w", domain.ErrInvalidStringStruct)
	}

	for i := 0; i < len(dataBytes); i += 2 {
		hexStr.Data = append(hexStr.Data, uint16(dataBytes[i])<<8+uint16(dataBytes[i+1]))
	}

	fmt.Println(hexStr)
	return hexStr, nil
}
