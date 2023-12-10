package main

type PipeDirection int

const (
	Vertical PipeDirection = iota
	Horizontal
	NorthToEast
	NorthToWest
	SouthToEast
	SouthToWest
	Empty
	Start
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func convertToGrid(lines []string) [][]PipeDirection {
	var letterToPipeDirection = map[rune]PipeDirection{
		'|': Vertical,
		'-': Horizontal,
		'L': NorthToEast,
		'J': NorthToWest,
		'F': SouthToEast,
		'7': SouthToWest,
		'.': Empty,
		'S': Start,
	}

	var gridToReturn = make([][]PipeDirection, len(lines))
	for i := range gridToReturn {
		gridToReturn[i] = make([]PipeDirection, len(lines[i]))
	}

	for i, line := range lines {
		for j, r := range line {
			dir := letterToPipeDirection[r]
			gridToReturn[i][j] = dir
		}
	}

	return gridToReturn
}

func findStart(grid [][]PipeDirection) [2]int {
	for i, row := range grid {
		for j, direction := range row {
			if direction == Start {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{-1, -1}
}

func findFirstDirectionsFromStart(grid [][]PipeDirection, startCoords [2]int) [2]Direction {
	var directions [2]Direction
	var coordsFound = 0

	if startCoords[0] > 0 {
		switch grid[startCoords[0]-1][startCoords[1]] {
		case Vertical, SouthToEast, SouthToWest:
			directions[coordsFound] = North
			coordsFound += 1
			if coordsFound == 2 {
				return directions
			}
		}
	}
	if startCoords[0] < len(grid)-1 {
		switch grid[startCoords[0]+1][startCoords[1]] {
		case Vertical, NorthToEast, NorthToWest:
			directions[coordsFound] = South
			coordsFound += 1
			if coordsFound == 2 {
				return directions
			}
		}
	}
	if startCoords[1] > 0 {
		switch grid[startCoords[0]][startCoords[1]-1] {
		case Horizontal, NorthToEast, SouthToEast:
			directions[coordsFound] = West
			coordsFound += 1
			if coordsFound == 2 {
				return directions
			}
		}
	}
	if startCoords[1] < len(grid[0])-1 {
		switch grid[startCoords[0]][startCoords[1]+1] {
		case Horizontal, NorthToWest, SouthToWest:
			directions[coordsFound] = East
			coordsFound += 1
			if coordsFound == 2 {
				return directions
			}
		}
	}
	panic("Did not find two directions from start")
}

func findNextCoord(grid [][]PipeDirection, coordinate [2]int, directionTravelled Direction) ([2]int, Direction) {
	var fromNorthDeltaMap = map[PipeDirection][2]int{
		Vertical:    {1, 0},
		NorthToEast: {0, 1},
		NorthToWest: {0, -1},
	}

	var fromEastDeltaMap = map[PipeDirection][2]int{
		Horizontal:  {0, -1},
		NorthToEast: {-1, 0},
		SouthToEast: {1, 0},
	}

	var fromSouthDeltaMap = map[PipeDirection][2]int{
		Vertical:    {-1, 0},
		SouthToEast: {0, 1},
		SouthToWest: {0, -1},
	}

	var fromWestDeltaMap = map[PipeDirection][2]int{
		Horizontal:  {0, 1},
		NorthToWest: {-1, 0},
		SouthToWest: {1, 0},
	}

	var delta [2]int
	pipeType := grid[coordinate[0]][coordinate[1]]

	switch directionTravelled {
	case South:
		delta = fromNorthDeltaMap[pipeType]
	case West:
		delta = fromEastDeltaMap[pipeType]
	case North:
		delta = fromSouthDeltaMap[pipeType]
	case East:
		delta = fromWestDeltaMap[pipeType]
	}

	var deltaToDirectionMap = map[[2]int]Direction{
		[2]int{1, 0}:  South,
		[2]int{-1, 0}: North,
		[2]int{0, 1}:  East,
		[2]int{0, -1}: West,
	}

	return [2]int{delta[0] + coordinate[0], delta[1] + coordinate[1]}, deltaToDirectionMap[delta]
}

func addDeltaToCoords(coords [2]int, delta [2]int) [2]int {
	return [2]int{coords[0] + delta[0], coords[1] + delta[1]}
}
