package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"

	"github.com/NagariaHussain/aoc2024_go/utils"
)

func main() {
	scanner, _ := utils.GetFileAsScanner("input.txt")

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}

	fmt.Printf("part 1: %v\n", GetPart1(buf.String()))
}

func GetPart1(input string) (answer int64) {
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matchingExpressions := re.FindAllStringSubmatch(input, -1)

	for _, match := range matchingExpressions {
		num1, _ := strconv.ParseInt(match[1], 10, 0)
		num2, _ := strconv.ParseInt(match[2], 10, 0)

		mult := num1 * num2
		answer += mult
	}
	return answer
}
