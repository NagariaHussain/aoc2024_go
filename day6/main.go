package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/NagariaHussain/aoc2024_go/utils"
)

func main() {
	lines := utils.GetLines("input.txt")

	fmt.Printf("part 1 answer: %v\n", GetPart1(lines))
}

type Location [2]int
type DirectionVector [2]int

var Top = DirectionVector{-1, 0}
var Left = DirectionVector{0, -1}
var Right = DirectionVector{0, 1}
var Bottom = DirectionVector{1, 0}

const (
	Obstacle      = "#"
	VisitedTop    = "^"
	VisitedLeft   = "<"
	VisitedRight  = ">"
	VisitedBottom = "v"
)

var visitedMap = map[DirectionVector]string{
	Top:    VisitedTop,
	Left:   VisitedLeft,
	Right:  VisitedRight,
	Bottom: VisitedBottom,
}

func alreadyVisited(spot string) bool {
	visitedSymbols := []string{"^", ">", "<", "v"}
	return slices.Contains(visitedSymbols, spot)
}

func GetPart1(lines []string) (answer int) {
	answer = 1
	curDirection := Top

	grid, curPos := getGridAndStart(lines)

	yMax := len(grid)
	xMax := len(grid[0])

	for {
		nextPos := Location{curPos[0] + curDirection[0], curPos[1] + curDirection[1]}

		// if nextPos out of bounds, its over!
		if isOutOfBounds(nextPos, xMax, yMax) {
			break
		}

		nextSpot := grid[nextPos[0]][nextPos[1]]

		if nextSpot == Obstacle {
			// change direction
			if curDirection == Top {
				curDirection = Right
			} else if curDirection == Left {
				curDirection = Top
			} else if curDirection == Right {
				curDirection = Bottom
			} else {
				curDirection = Left
			}
		} else if alreadyVisited(nextSpot) {
			curPos = nextPos
			grid[curPos[0]][curPos[1]] = visitedMap[curDirection]
		} else {
			// we can go here
			answer += 1
			curPos = nextPos
			grid[curPos[0]][curPos[1]] = visitedMap[curDirection]
		}
	}

	for _, line := range grid {
		fmt.Println(line)
	}

	return
}

func GetPart2(lines []string) (answer int) {
	return
}

func isOutOfBounds(coord Location, xMax, yMax int) bool {
	if coord[0] < 0 || coord[0] >= yMax {
		return true
	}

	if coord[1] < 0 || coord[1] >= xMax {
		return true
	}

	return false
}

func Walk() (result string) {
	return
}

func getGridAndStart(lines []string) ([][]string, Location) {
	grid := make([][]string, len(lines))
	curPos := *new(Location)
	startPosFound := false

	for index, line := range lines {
		grid[index] = strings.Split(line, "")

		if !startPosFound {
			guardCurrentLocation := slices.Index(grid[index], "^")
			if guardCurrentLocation != -1 {
				curPos[0] = index
				curPos[1] = guardCurrentLocation
				startPosFound = true
				grid[curPos[0]][curPos[1]] = visitedMap[Top]
			}
		}
	}

	return grid, curPos
}
