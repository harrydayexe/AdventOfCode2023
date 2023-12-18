package main

import "fmt"

func part1(lines []string) int {
	var grid = createGrid(lines)
	startNode := grid[0][0]
	endNode := grid[len(grid)-1][len(grid[0])-1]
	aStar(startNode, endNode, grid)

	var heatLoss = 0
	next := endNode
	for next != nil {
		fmt.Println(next.row+1, next.col+1)
		heatLoss += next.heatLoss
		next = next.prev
	}
	return heatLoss
}
