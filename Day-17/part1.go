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
		heatLoss += next.heatLoss
		next = next.prev
	}

	printGrid(endNode, grid)

	return heatLoss
}

func printGrid(end *node, grid [][]*node) {
	var output = make([][]rune, len(grid))
	for i := range output {
		output[i] = make([]rune, len(grid[0]))
	}

	for i, nodes := range grid {
		for j, n := range nodes {
			output[i][j] = rune(n.heatLoss + '0')
		}
	}

	var dirToRune = map[Direction]rune{
		North: '^',
		East:  '>',
		South: 'V',
		West:  '<',
	}

	next := end
	for next != nil {
		output[next.row][next.col] = dirToRune[next.inDirection]
		next = next.prev
	}

	for _, runes := range output {
		fmt.Println(string(runes))
	}
}
