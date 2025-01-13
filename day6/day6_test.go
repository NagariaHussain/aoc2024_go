package main

import (
	"strings"
	"testing"
)

func TestDay6(t *testing.T) {
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

	t.Run("correct for part 1 example", func(t *testing.T) {
		want := 41
		got := GetPart1(lines)

		if got != want {
			t.Errorf("incorrect answer for day 6 part 1. Want %v, got %v.", want, got)
		}
	})

	t.Run("correct for part 2 example", func(t *testing.T) {
		want := 6
		got := GetPart2(lines)

		if got != want {
			t.Errorf("incorrect answer for day 6 part 2. Want %v, got %v.", want, got)
		}
	})
}
