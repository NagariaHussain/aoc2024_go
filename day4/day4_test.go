package main

import (
	"strings"
	"testing"
)

func TestDay4(t *testing.T) {
	t.Run("returns correct answer for part 1", func(t *testing.T) {
		testString := `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

		testInput := strings.Split(testString, "\n")

		want := 18
		got := GetPart1(testInput)

		if got != want {
			t.Errorf("incorrect answer, got: %v, want: %v", got, want)
		}
	})
}
