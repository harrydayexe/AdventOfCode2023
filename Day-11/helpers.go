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

func getExpandedRowsAndCols(grid [][]bool) ([]int, []int) {
	var notEmptyRows, notEmptyCols = make(map[int]struct{}), make(map[int]struct{})

	for i, row := range grid {
		for j, cell := range row {
			if cell {
				notEmptyRows[i] = struct{}{}
				notEmptyCols[j] = struct{}{}
			}
		}
	}

	var fullRows, fullCols = make(map[int]struct{}), make(map[int]struct{})
	for i := 0; i < len(grid); i++ {
		fullRows[i] = struct{}{}
	}
	for i := 0; i < len(grid[0]); i++ {
		fullCols[i] = struct{}{}
	}

	var resultCols []int
	for v := range fullCols {
		_, prs := notEmptyCols[v]
		if prs {
			continue
		}
		resultCols = append(resultCols, v)
	}

	var resultRows []int
	for v := range fullRows {
		_, prs := notEmptyRows[v]
		if prs {
			continue
		}
		resultRows = append(resultRows, v)
	}

	return resultRows, resultCols
}

func findShortestPathLengthWithExpansion(startCoord [2]int, endCoord [2]int, expandedRows []int, expandedCols []int) int {
	var deltaX, deltaY = int(math.Abs(float64(startCoord[0] - endCoord[0]))), int(math.Abs(float64(startCoord[1] - endCoord[1])))

	for _, rowNum := range expandedRows {
		if min(startCoord[0], endCoord[0]) < rowNum && rowNum < max(startCoord[0], endCoord[0]) {
			deltaY += 999999
		}
	}
	for _, colNum := range expandedCols {
		if min(startCoord[1], endCoord[1]) < colNum && colNum < max(startCoord[1], endCoord[1]) {
			deltaX += 999999
		}
	}

	return deltaX + deltaY
}
