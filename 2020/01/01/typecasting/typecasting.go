package typecasting

import (
	"errors"
	"strconv"
)

var ErrTypeConversion = errors.New("could not convert string into integer")

// StringSliceToInt converts a slice of strings to int.
func StringSliceToInt(s []string) ([]int, error) {
	var intSlice = make([]int, len(s))

	var err error

	for i, v := range s {
		intSlice[i], err = strconv.Atoi(v)
		if err != nil {
			err = ErrTypeConversion

			return intSlice, err
		}
	}

	return intSlice, nil
}
