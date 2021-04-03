package read

import (
	"bufio"
	"errors"
	"os"
	"strconv"
)

var ErrFileNotFound = errors.New("no esta el archivo")

// SliceInput opens a text file and returns a slice containing its lines
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

// StringSliceToInt converts a slice of strings to int.
func StringSliceToInt(s []string) []int {
	var intSlice = make([]int, len(s))

	for i, v := range s {
		intSlice[i], _ = strconv.Atoi(v)
	}

	return intSlice
}
