package main

func part1(lines []string) int {
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

	var numOfSteps = 1
	for pointer1 != pointer2 {
		pointer1, direction1 = findNextCoord(grid, pointer1, direction1)
		pointer2, direction2 = findNextCoord(grid, pointer2, direction2)
		numOfSteps += 1
	}

	return numOfSteps
}
