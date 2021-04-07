package main_test

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/rhlcoder/aoc/2020/01/01/accounting"
	"github.com/rhlcoder/aoc/2020/01/01/reading"
	"github.com/rhlcoder/aoc/2020/01/01/typecasting"
)

const (
	filename = "input.txt"
	// filename = "testdata/test_input.txt".
	objetive = 2020
)

func BenchmarkSumObjetive(b *testing.B) {
	s, _ := reading.SliceInput(filename)
	slice, err := typecasting.StringSliceToInt(s)

	if errors.Is(err, typecasting.ErrTypeConversion) {
		log.Printf("Error: %s\n", typecasting.ErrTypeConversion)
		os.Exit(1)
	}

	for i := 0; i < b.N; i++ {
		accounting.SumObjetive(slice, objetive) //nolint
	}
}
