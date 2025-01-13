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
	fmt.Printf("part 2 answer: %v\n", GetPart2(lines))
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
	gs.grid[loc[0]][loc[1]] = "$"

	if exists {
		gs.walkHistory[loc] = append(gs.walkHistory[loc], dir)
	} else {
		gs.walkHistory[loc] = []DirectionVector{dir}
	}
}

func GetPart1(lines []string) (answer int) {
	gridState := getGridState(lines)
	gridState.Walk()

	return len(gridState.walkHistory)
}

func GetPart2(lines []string) (answer int) {
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			gs := getGridState(lines)

			if gs.grid[i][j] == Obstacle {
				continue
			}

			// add obstacle
			gs.grid[i][j] = Obstacle
			goesOut := gs.Walk()

			if !goesOut {
				answer += 1
			}
		}
	}

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

func (gridState *GridState) Walk() (goesOut bool) { // true if goes out, false if loops
	grid := gridState.grid
	curPos := gridState.startLocation
	curDirection := gridState.startDirection

	yMax := len(grid)
	xMax := len(grid[0])

	for {
		nextPos := Location{curPos[0] + curDirection[0], curPos[1] + curDirection[1]}

		// if nextPos out of bounds, its over!
		if isOutOfBounds(nextPos, xMax, yMax) {
			return true
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
		} else if gridState.hasAlreadyVisitedWithDirection(nextPos, curDirection) {
			// IN A LOOP!
			return false
		} else {
			curPos = nextPos
			gridState.trackVisit(curPos, curDirection)
		}
	}
}

func (gs *GridState) printGrid() {
	for _, row := range gs.grid {
		fmt.Println(row)
	}
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

	gridState.trackVisit(gridState.startLocation, gridState.startDirection)

	return gridState
}
