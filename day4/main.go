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
	fmt.Printf("part 2 answer: %v\n", GetPart2(lines))
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

func GetPart2(lines []string) (xmasCount int) {
	stringMatrix := getAsMatrix(lines)

	xMax := len(stringMatrix) - 1
	yMax := len(stringMatrix[0]) - 1

	for x := 0; x <= xMax; x++ {
		for y := 0; y <= yMax; y++ {
			char := stringMatrix[x][y]

			if char == "A" {
				howMany := searchForXMASFixed(stringMatrix, x, y, xMax, yMax)
				xmasCount += howMany
			}
		}
	}

	return xmasCount
}

func searchForXMASFixed(matrix [][]string, x, y, xMax, yMax int) int {
	BRx, BRy := x+1, y+1
	TLx, TLy := x-1, y-1

	BLx, BLy := x-1, y+1
	TRx, TRy := x+1, y-1

	foundLeft, foundRight := false, false

	if !(isOutOfBounds(BRx, BRy, xMax, yMax) || isOutOfBounds(TLx, TLy, xMax, yMax)) {
		// 1: MAS left
		foundLeft = foundLeft || (matrix[BRx][BRy] == "M" && matrix[TLx][TLy] == "S")
		// 2: SAM left
		foundLeft = foundLeft || (matrix[BRx][BRy] == "S" && matrix[TLx][TLy] == "M")
	}

	if !(isOutOfBounds(BLx, BLy, xMax, yMax) || isOutOfBounds(TRx, TRy, xMax, yMax)) {
		// 3: MAS right
		foundRight = foundRight || (matrix[BLx][BLy] == "M" && matrix[TRx][TRy] == "S")

		// 4: SAM right
		foundRight = foundRight || (matrix[BLx][BLy] == "S" && matrix[TRx][TRy] == "M")
	}

	if foundLeft && foundRight {
		return 1
	}

	return 0
}

func GetPart1(lines []string) (xmasCount int) {
	stringMatrix := getAsMatrix(lines)

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

func getAsMatrix(lines []string) [][]string {
	stringMatrix := make([][]string, 0, len(lines))

	for _, line := range lines {
		stringMatrix = append(stringMatrix, strings.Split(line, ""))
	}

	return stringMatrix
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
