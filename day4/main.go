package main

import (
	"fmt"
	"strings"

	"github.com/NagariaHussain/aoc2024_go/utils"
)

func main() {
	lines := make([]string, 0)

	scanner, _ := utils.GetFileAsScanner("input.txt")

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("part 1 answer: %v\n", GetPart1(lines))
}

var directions = [][]int{
	{0, 1},
	{0, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{-1, -1},
	{1, 1},
	{1, -1},
}

func GetPart1(lines []string) (xmasCount int) {
	stringMatrix := make([][]string, 0, len(lines))

	for _, line := range lines {
		stringMatrix = append(stringMatrix, strings.Split(line, ""))
	}

	xMax := len(stringMatrix) - 1
	yMax := len(stringMatrix[0]) - 1

	for x := 0; x <= xMax; x++ {
		for y := 0; y <= yMax; y++ {
			char := stringMatrix[x][y]

			if char == "X" {
				howMany := searchForXMAS(stringMatrix, x, y, xMax, yMax)
				xmasCount += howMany
			}
		}
	}

	return xmasCount
}

func isOutOfBounds(x, y, xMax, yMax int) bool {
	if x > xMax || x < 0 {
		return true
	}

	if y > yMax || y < 0 {
		return true
	}

	return false
}

func searchForXMAS(matrix [][]string, x, y, xMax, yMax int) (foundCount int) {
	for _, dir := range directions {
		found := true

		dx := dir[0]
		dy := dir[1]

		nextX := x + dx
		nextY := y + dy

		if isOutOfBounds(nextX, nextY, xMax, yMax) {
			continue
		}
		found = found && (matrix[nextX][nextY] == "M")
		if !found {
			continue
		}

		nextX += dx
		nextY += dy
		if isOutOfBounds(nextX, nextY, xMax, yMax) {
			continue
		}
		found = found && (matrix[nextX][nextY] == "A")
		if !found {
			continue
		}

		nextX += dx
		nextY += dy
		if isOutOfBounds(nextX, nextY, xMax, yMax) {
			continue
		}
		found = found && (matrix[nextX][nextY] == "S")

		if found {
			foundCount += 1
		}
	}
	return
}
