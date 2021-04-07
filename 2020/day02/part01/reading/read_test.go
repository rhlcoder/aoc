package reading_test

import (
	"testing"

	. "github.com/rhlcoder/aoc/2020/day02/part01/reading"
)

const input = "../testdata/test_input.txt"

func TestSliceInput(t *testing.T) {
	t.Parallel()

	t.Run("length test", func(t *testing.T) {
		t.Parallel()

		got, _ := SliceInput(input)
		want := 6

		if len(got) != want {
			t.Errorf("got %d, want %d", len(got), want)
		}
	})

	t.Run("First line shoud be '1721'", func(t *testing.T) {
		t.Parallel()
		got, _ := SliceInput(input)
		want := "1721"

		if got[0] != want {
			t.Errorf("got %q, want %q", got[0], want)
		}
	})
}
