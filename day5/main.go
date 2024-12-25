package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/NagariaHussain/aoc2024_go/utils"
)

func main() {

	rules := utils.GetLines("rules.txt")
	updates := utils.GetLines("updates.txt")

	fmt.Printf("part 1: %v\n", GetPart1(rules, updates))
}

func GetPart1(orderingRules, updates []string) (midSum int64) {
	rules := make(map[int64][]int64)

	for _, rule := range orderingRules {
		parts := utils.Map(strings.Split(rule, "|"), utils.ToInt)
		before, after := parts[0], parts[1]

		rules[before] = append(rules[before], after)
	}

	for _, update := range updates {
		printOrder := utils.Map(strings.Split(update, ","), utils.ToInt)
		isCorrect := true

		for index, curNum := range printOrder {
			for i := index + 1; i < len(printOrder); i++ {
				nextNum := printOrder[i]

				// curNum comes before nextNum
				// in other words no rule for nextNum before curNum should exist AND
				if slices.Contains(rules[nextNum], curNum) { // TODO: implement set for faster membership checks
					isCorrect = false
					break
				}
			}

			if !isCorrect {
				break
			}
		}

		if isCorrect {
			// find mid and sum it out
			mid := len(printOrder) / 2
			midSum += printOrder[mid]
		}
	}

	return
}

func GetPart2(orderingRules, updates []string) (midSum int64) {
	return
}
