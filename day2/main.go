package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/NagariaHussain/aoc2024_go/utils"
)

func main() {
	scanner, _ := utils.GetFileAsScanner("input.txt")

	var report []int64
	var numSafePart1 uint
	var numSafeAfterDampening uint

	for scanner.Scan() {
		report = utils.Map(strings.Fields((scanner.Text())), func(s string) int64 {
			num, _ := strconv.ParseInt(s, 10, 64)
			return num
		})

		if isSafe(report) {
			numSafePart1 += 1
		} else if isSafeAfterDampening(report) {
			numSafeAfterDampening += 1
		}
	}

	fmt.Printf("part 1: %v\n", numSafePart1)
	fmt.Printf("part 2: %v\n", numSafePart1+numSafeAfterDampening)
}

func isSafeAfterDampening(report []int64) bool {
	for index := range report {
		if index > len(report)-1 {
			break
		}

		probablyCorrectedReport := slices.Concat(report[:index], report[index+1:])
		if isSafe(probablyCorrectedReport) {
			return true
		}
	}

	return false
}

func isSafe(report []int64) bool {
	diffs := make([]int64, len(report)-1)
	for i := 0; i < len(report)-1; i++ {
		diffs[i] = report[i+1] - report[i]
	}

	if diffs[0] == 0 {
		return false
	}

	mightBeIncreasing := (diffs[0] > 0) && (diffs[0] < 4) // between 1 and 3
	reportType := "DSC"
	if mightBeIncreasing {
		reportType = "ASC"
	}

	for _, diff := range diffs {
		if reportType == "ASC" && (diff <= 0 || diff > 3) {
			return false
		} else if reportType == "DSC" && (diff >= 0 || diff < -3) {
			return false
		}
	}

	return true
}
