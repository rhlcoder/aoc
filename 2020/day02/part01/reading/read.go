package reading

import (
	"bufio"
	"errors"
	"os"
)

var ErrFileNotFound = errors.New("input file not found")

func SliceInput(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, ErrFileNotFound
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var slice = make([]string, 0)

	for scanner.Scan() {
		slice = append(slice, scanner.Text())
	}

	return slice, nil
}
