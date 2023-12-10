package main

func part2(lines []string) int {
	grid := convertToGrid(lines)
	startCoords := findStart(grid)
	directionsFromStart := findFirstDirectionsFromStart(grid, startCoords)

	var directionDelta = map[Direction][2]int{
		North: {-1, 0},
		South: {1, 0},
		East:  {0, 1},
		West:  {0, -1},
	}

	var pointer1, pointer2 = addDeltaToCoords(startCoords, directionDelta[directionsFromStart[0]]), addDeltaToCoords(startCoords, directionDelta[directionsFromStart[1]])
	var direction1, direction2 = directionsFromStart[0], directionsFromStart[1]

	var pipeGrid = make([][]PipeDirection, len(grid))
	for i := range pipeGrid {
		pipeGrid[i] = make([]PipeDirection, len(grid[i]))
	}
	pipeGrid[startCoords[0]][startCoords[1]] = calculateStartPipe(direction1, direction2)
	for pointer1 != pointer2 {
		pipeGrid[pointer1[0]][pointer1[1]] = grid[pointer1[0]][pointer1[1]]
		pipeGrid[pointer2[0]][pointer2[1]] = grid[pointer2[0]][pointer2[1]]
		pointer1, direction1 = findNextCoord(grid, pointer1, direction1)
		pointer2, direction2 = findNextCoord(grid, pointer2, direction2)
	}

	var nestTileCount = 0
	for i, row := range pipeGrid {
		for j, cell := range row {
			if cell == 0 && isNestTile(pipeGrid, [2]int{i, j}) {
				nestTileCount += 1
			}
		}
	}

	return nestTileCount
}

func calculateStartPipe(direction1 Direction, direction2 Direction) PipeDirection {
	switch direction1 {
	case North:
		switch direction2 {
		case East:
			return NorthToEast
		case South:
			return Vertical
		case West:
			return NorthToWest
		}
	case East:
		switch direction2 {
		case North:
			return NorthToEast
		case South:
			return SouthToEast
		case West:
			return Horizontal
		}
	case South:
		switch direction2 {
		case North:
			return Vertical
		case East:
			return SouthToEast
		case West:
			return SouthToWest
		}
	case West:
		switch direction2 {
		case North:
			return NorthToWest
		case East:
			return Horizontal
		case South:
			return SouthToWest
		}
	}
	panic("Same directions given to calculate start pipe direction")
}

func countPipeCrossingInDirection(grid [][]PipeDirection, startCoord [2]int, direction Direction) int {
	var delta [2]int
	switch direction {
	case North:
		delta = [2]int{-1, 0}
	case East:
		delta = [2]int{0, 1}
	case South:
		delta = [2]int{1, 0}
	case West:
		delta = [2]int{0, -1}
	}

	var pointer = addDeltaToCoords(startCoord, delta)
	var totalPipeCrossings = 0
	var lastCornerCrossing PipeDirection
	for 0 <= pointer[0] && pointer[0] < len(grid) && 0 <= pointer[1] && pointer[1] < len(grid[0]) {
		switch grid[pointer[0]][pointer[1]] {
		case Vertical:
			if direction == West || direction == East {
				totalPipeCrossings += 1
				lastCornerCrossing = Empty
			}
		case Horizontal:
			if direction == North || direction == South {
				totalPipeCrossings += 1
				lastCornerCrossing = Empty
			}
		case NorthToEast:
			if lastCornerCrossing == Empty {
				lastCornerCrossing = NorthToEast
			} else if lastCornerCrossing == SouthToWest {
				totalPipeCrossings += 1
				lastCornerCrossing = Empty
			} else {
				lastCornerCrossing = Empty
			}
		case NorthToWest:
			if lastCornerCrossing == Empty {
				lastCornerCrossing = NorthToWest
			} else if lastCornerCrossing == SouthToEast {
				totalPipeCrossings += 1
				lastCornerCrossing = Empty
			} else {
				lastCornerCrossing = Empty
			}
		case SouthToEast:
			if lastCornerCrossing == Empty {
				lastCornerCrossing = SouthToEast
			} else if lastCornerCrossing == NorthToWest {
				totalPipeCrossings += 1
				lastCornerCrossing = Empty
			} else {
				lastCornerCrossing = Empty
			}
		case SouthToWest:
			if lastCornerCrossing == Empty {
				lastCornerCrossing = SouthToWest
			} else if lastCornerCrossing == NorthToEast {
				totalPipeCrossings += 1
				lastCornerCrossing = Empty
			} else {
				lastCornerCrossing = Empty
			}
		}
		pointer = addDeltaToCoords(pointer, delta)
	}

	return totalPipeCrossings
}

func isNestTile(grid [][]PipeDirection, coord [2]int) bool {
	for i := 0; i < 4; i++ {
		if countPipeCrossingInDirection(grid, coord, Direction(i))%2 == 0 {
			return false
		}
	}
	return true
}
