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

type GridState struct {
	grid           [][]string
	startLocation  Location
	startDirection DirectionVector
	walkHistory    map[Location][]DirectionVector // e.g. (1, 0): ["^", ">"]
}

func (gs *GridState) hasAlreadyVisitedWithDirection(loc Location, dir DirectionVector) bool {
	dirsVisited, exists := gs.walkHistory[loc]

	if !exists {
		return false
	}

	for _, dirVisit := range dirsVisited {
		if dirVisit == dir {
			return true
		}
	}

	return false
}

func (gs *GridState) trackVisit(loc Location, dir DirectionVector) {
	_, exists := gs.walkHistory[loc]

	if exists {
		gs.walkHistory[loc] = append(gs.walkHistory[loc], dir)
	} else {
		gs.walkHistory[loc] = []DirectionVector{dir}
	}
}

func GetPart1(lines []string) (answer int) {
	gridState := getGridState(lines)
	grid := gridState.grid
	curPos := gridState.startLocation
	curDirection := gridState.startDirection

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
		} else {
			curPos = nextPos
			gridState.trackVisit(curPos, curDirection)
		}
	}

	return len(gridState.walkHistory)
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

func getGridState(lines []string) (gridState GridState) {
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
			}
		}
	}

	gridState.grid = grid
	gridState.startLocation = curPos
	gridState.startDirection = Top
	gridState.walkHistory = make(map[Location][]DirectionVector)

	gridState.walkHistory[curPos] = []DirectionVector{gridState.startDirection}

	return gridState
}
