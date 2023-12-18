package main

func part1(lines []string) int {
	var grid = createGrid(lines)
	heatLoss := aStar(grid)

	return heatLoss
}
