package files

import (
	"bufio"
	"errors"
	"os"
)

const ErrorInOpenFile string = "one error occurred trying to open file: "

func OpenFile(fileName string) (*os.File, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New(ErrorInOpenFile + err.Error())
	}

	return file, nil
}

func ScannerFile(file *os.File) *bufio.Scanner {
	scanner := bufio.NewScanner(file)
	return scanner
}
