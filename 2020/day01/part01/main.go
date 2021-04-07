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

	x, y, it, err := accounting.SumPair(i, objetive)
	errorhandling.CheckError(err, accounting.ErrNoPairFound)

	tx, ty, tz, itt, err := accounting.SumTern(i, objetive)
	errorhandling.CheckError(err, accounting.ErrNoPairFound)

	log.Printf("Pair %d iterations: %d + %d = %d. And its product is: %d", it, x, y, objetive, x*y)
	log.Printf("Tern %d iterations: %d + %d + %d = %d. And its product is: %d", itt, tx, ty, tz, objetive, tx*ty*tz)
}
