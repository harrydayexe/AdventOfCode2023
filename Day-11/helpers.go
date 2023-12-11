package main

import "math"

func cleanData(lines []string) [][]bool {
	var returnGrid = make([][]bool, len(lines))
	for i := range returnGrid {
		returnGrid[i] = make([]bool, len(lines[0]))
	}

	for i, line := range lines {
		for j, cell := range line {
			if cell == '#' {
				returnGrid[i][j] = true
			}
		}
	}
	return returnGrid
}

func expandSpace(grid [][]bool) [][]bool {
	var notEmptyRows = make(map[int]struct{})
	var notEmptyCols = make(map[int]struct{})

	for i, row := range grid {
		for j, cell := range row {
			if cell {
				notEmptyRows[i] = struct{}{}
				notEmptyCols[j] = struct{}{}
			}
		}
	}

	var newGrid = make([][]bool, len(grid))
	for i := range newGrid {
		newGrid[i] = make([]bool, len(grid[0]))
		copy(newGrid[i], grid[i])
	}

	var offset = 0
	for i := 0; i < len(grid); i++ {
		_, prs := notEmptyRows[i]
		if !prs {
			newGrid = append(newGrid[:i+offset+1], newGrid[i+offset:]...)
			newGrid[i+offset] = make([]bool, len(grid[0]))
			offset += 1
		}
	}

	offset = 0
	for i := 0; i < len(grid[0]); i++ {
		_, prs := notEmptyCols[i]
		if !prs {
			for j := range newGrid {
				newGrid[j] = append(newGrid[j][:i+offset+1], newGrid[j][i+offset:]...)
				newGrid[j][i+offset] = false
			}
			offset += 1
		}
	}

	return newGrid
}

func findShortestPathLength(startCoord [2]int, endCoord [2]int) int {
	return int(math.Abs(float64(startCoord[0]-endCoord[0]))) + int(math.Abs(float64(startCoord[1]-endCoord[1])))
}
