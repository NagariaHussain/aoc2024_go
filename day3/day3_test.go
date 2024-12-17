package main

import "testing"

func TestDay2Part1(t *testing.T) {
	t.Run("returns correct answer of part 1", func(t *testing.T) {
		testInput := `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
		var want int64 = 161

		got := GetPart1(testInput)

		if got != want {
			t.Errorf("does not return correct result, got: %v, want: %v\n", got, want)
		}
	})

	t.Run("returns correct answer of part 2", func(t *testing.T) {
		testInput := `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`
		var want int64 = 48

		got := GetPart2(testInput)

		if got != want {
			t.Errorf("does not return correct result, got: %v, want: %v\n", got, want)
		}
	})
}
