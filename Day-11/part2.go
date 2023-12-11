package main

func part2(lines []string) int {
	grid := cleanData(lines)
	expandedRows, expandedCols := getExpandedRowsAndCols(grid)

	var runningTotal = 0
	var galaxies = make(map[[2]int]struct{})

	// Find all galaxies
	for i, row := range grid {
		for j, cell := range row {
			if cell {
				galaxies[[2]int{i, j}] = struct{}{}
			}
		}
	}

	var galaxyCoords [][2]int
	for galaxyCoord := range galaxies {
		galaxyCoords = append(galaxyCoords, galaxyCoord)
	}

	for i, coord := range galaxyCoords {
		for j := i + 1; j < len(galaxyCoords); j++ {
			runningTotal += findShortestPathLengthWithExpansion(coord, galaxyCoords[j], expandedRows, expandedCols)
		}
	}

	return runningTotal
}
