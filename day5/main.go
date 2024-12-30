package main

import (
	"fmt"
	"strings"

	"github.com/NagariaHussain/aoc2024_go/utils"
)

func main() {
	rules := utils.GetLines("rules.txt")
	updates := utils.GetLines("updates.txt")

	fmt.Printf("part 1: %v\n", GetPart1(rules, updates))
	fmt.Printf("part 2: %v\n", GetPart2(rules, updates))
}

func GetPart1(orderingRules, updates []string) (midSum int64) {
	rules := getRules(orderingRules)

	for _, update := range updates {
		printOrder := utils.Map(strings.Split(update, ","), utils.ToInt)
		isCorrect := isUpdateCorrect(printOrder, rules)

		if isCorrect {
			// find mid and sum it out
			mid := len(printOrder) / 2
			midSum += printOrder[mid]
		}
	}

	return
}

func isUpdateCorrect(printOrder []int64, rules map[int64]utils.Set) (isCorrect bool) {
	isCorrect = true

	for index, curNum := range printOrder {
		for i := index + 1; i < len(printOrder); i++ {
			nextNum := printOrder[i]

			// curNum comes before nextNum
			// in other words no rule for nextNum before curNum should exist AND
			afterSet := rules[nextNum]
			if afterSet.Has(curNum) {
				isCorrect = false
				return isCorrect
			}
		}

		if !isCorrect {
			return false
		}
	}

	return
}

func GetPart2(orderingRules, updates []string) (midSum int64) {
	rules := getRules(orderingRules)

	for _, update := range updates {
		printOrder := utils.Map(strings.Split(update, ","), utils.ToInt)
		isCorrect := isUpdateCorrect(printOrder, rules)

		if !isCorrect {
			// correct it
			printOrder = getCorrectedOrder(printOrder, rules)

			// then sum the mid
			// find mid and sum it out
			mid := len(printOrder) / 2
			midSum += printOrder[mid]
		}
	}

	return
}

func getCorrectedOrder(incorrectOrder []int64, rules map[int64]utils.Set) []int64 {
	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < len(incorrectOrder)-1; i++ {
			curNum := incorrectOrder[i]
			next := incorrectOrder[i+1]

			// curNum comes before nextNum
			// in other words no rule for nextNum before curNum should exist AND
			afterSet := rules[next]
			if afterSet.Has(curNum) {
				// to be swapped!
				incorrectOrder[i], incorrectOrder[i+1] = incorrectOrder[i+1], incorrectOrder[i]
				swapped = true
			}
		}
	}

	return incorrectOrder
}

func getRules(orderingRules []string) (rules map[int64]utils.Set) {
	rules = make(map[int64]utils.Set)

	for _, rule := range orderingRules {
		parts := utils.Map(strings.Split(rule, "|"), utils.ToInt)
		before, after := parts[0], parts[1]

		currentSet, ok := rules[before]

		if ok {
			currentSet.Add(after)
		} else {
			newSet := utils.NewSet()
			newSet.Add(after)
			rules[before] = newSet
		}
	}

	return
}
