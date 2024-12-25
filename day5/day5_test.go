package main

import (
	"strings"
	"testing"

	"github.com/NagariaHussain/aoc2024_go/utils"
)

func TestDay5(t *testing.T) {
	pageOrderRules := `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13`

	updates := `75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

	t.Run("returns correct part 1", func(t *testing.T) {
		var want int64 = 143

		got := GetPart1(
			strings.Split(pageOrderRules, "\n"),
			strings.Split(updates, "\n"),
		)

		if want != got {
			t.Errorf("incorrect input, got: %v | want: %v\n", got, want)
		}
	})

	t.Run("returns correct part 2", func(t *testing.T) {
		var want int64 = 123

		got := GetPart2(
			strings.Split(pageOrderRules, "\n"),
			strings.Split(updates, "\n"),
		)

		if want != got {
			t.Errorf("incorrect input, got: %v | want: %v\n", got, want)
		}
	})

	t.Run("set data structure", func(t *testing.T) {
		set := utils.NewSet()

		if set.Has(16) {
			t.Errorf("`Has` should return false, returns true for %v", 16)
		}

		set.Add(16)

		if !set.Has(16) {
			t.Errorf("`Has` should return true, returns false for %v", 16)
		}
	})
}
