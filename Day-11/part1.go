package main

func part1(lines []string) int {
	grid := cleanData(lines)
	expandedGrid := expandSpace(grid)

	var runningTotal = 0
	var galaxies = make(map[[2]int]struct{})

	// Find all galaxies
	for i, row := range expandedGrid {
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
			runningTotal += findShortestPathLength(coord, galaxyCoords[j])
		}
	}

	return runningTotal
}
