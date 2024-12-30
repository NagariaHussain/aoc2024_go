package main

import (
	"strings"
	"testing"
)

func TestDay6(t *testing.T) {
	t.Run("correct for part 1 example", func(t *testing.T) {
		input := `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

		lines := strings.Split(input, "\n")
		want := 41
		got := GetPart1(lines)

		if got != want {
			t.Errorf("incorrect answer for day 6 part 1. Want %v, got %v.", want, got)
		}
	})
}
