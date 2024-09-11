package readservice

import (
	"asmparser/internal/domain"
	"bufio"
	"fmt"
	"os"
)

type ReadService struct {
	fileData []string
}

func New() *ReadService {
	return &ReadService{
		fileData: nil,
	}
}

func (s *ReadService) ReadFile(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return fmt.Errorf("ReadService - NewReader: %w. Reason: %w",
			domain.ErrFileNotOpen, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		s.fileData = append(s.fileData, scanner.Text())
	}

	return nil
}

func (s *ReadService) GetFileData() ([]string, error) {
	if len(s.fileData) == 0 {
		return nil, fmt.Errorf("ReadService - GetFileData: %w", domain.ErrFileNotRead)
	}
	return s.fileData, nil
}
