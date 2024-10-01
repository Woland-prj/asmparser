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

func (s *ReadService) ReadStdin() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		s.fileData = append(s.fileData, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func (s *ReadService) GetData() ([]string, error) {
	if len(s.fileData) == 0 {
		return nil, fmt.Errorf("ReadService - GetFileData: %w", domain.ErrFileNotRead)
	}
	return s.fileData, nil
}
