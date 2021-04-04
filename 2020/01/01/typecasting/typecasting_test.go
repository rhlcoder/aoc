package typecasting_test

import (
	"errors"
	"log"
	"os"
	"testing"

	. "github.com/rhlcoder/aoc/2020/01/01/reading"
	. "github.com/rhlcoder/aoc/2020/01/01/typecasting"
)

const input = "../testdata/test_input.txt"

func TestStringSliceToInt(t *testing.T) {
	t.Parallel()
	t.Run("Should get sum of integers", func(t *testing.T) {
		t.Parallel()
		s, _ := SliceInput(input)
		intSlice, err := StringSliceToInt(s)
		if errors.Is(err, ErrTypeConversion) {
			log.Printf("Error: %s\n", ErrTypeConversion)
			os.Exit(1)
		}

		want := 5496
		got := 0
		for _, v := range intSlice {
			got += v
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
