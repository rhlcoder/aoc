package main

import (
	"log"

	"github.com/rhlcoder/aoc/2020/01/01/accounting"
	"github.com/rhlcoder/aoc/2020/01/01/errorhandling"
	"github.com/rhlcoder/aoc/2020/01/01/reading"
	"github.com/rhlcoder/aoc/2020/01/01/typecasting"
)

const (
	filename = "input.txt"
	objetive = 2020
)

func main() {
	s, err := reading.SliceInput(filename)
	errorhandling.CheckError(err, reading.ErrFileNotFound)

	i, err := typecasting.StringSliceToInt(s)
	errorhandling.CheckError(err, typecasting.ErrTypeConversion)

	x, y, it, err := accounting.SumObjetive(i, objetive)
	errorhandling.CheckError(err, accounting.ErrNoPairFound)

	log.Printf("Took %d iterations: %d + %d = %d. And its product is: %d", it, x, y, objetive, x*y)
}
