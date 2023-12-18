package main

import (
	"container/heap"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type node struct {
	row, col   int
	rowD, colD int
	heatLoss   int
	n          int // How many in same direction
	index      int
}

func createGrid(lines []string) [][]int {
	var grid = make([][]int, len(lines))
	for i := range grid {
		grid[i] = make([]int, len(lines[0]))
	}

	for row, line := range lines {
		for col, ch := range line {
			grid[row][col] = int(ch - '0')
		}
	}

	return grid
}

func aStar(grid [][]int) int {
	var d = &priorityQueue{}
	// row, col, rowDelta, colDelta, n
	var seen = make(map[[5]int]struct{})
	heap.Init(d)
	heap.Push(d, &node{
		row:      0,
		col:      0,
		rowD:     0,
		colD:     0,
		heatLoss: 0,
		n:        0,
		index:    0,
	})

	for d.Len() > 0 {
		current := heap.Pop(d).(*node)
		if current.row == len(grid)-1 && current.col == len(grid[0])-1 {
			return current.heatLoss
		}

		_, prs := seen[[5]int{current.row, current.col, current.rowD, current.colD, current.n}]
		if prs {
			continue
		}
		seen[[5]int{current.row, current.col, current.rowD, current.colD, current.n}] = struct{}{}

		if current.n < 3 && (current.rowD != 0 || current.colD != 0) {
			nr := current.row + current.rowD
			nc := current.col + current.colD
			if 0 <= nr && nr < len(grid) && 0 <= nc && nc < len(grid[0]) {
				heap.Push(d, &node{
					row:      nr,
					col:      nc,
					rowD:     current.rowD,
					colD:     current.colD,
					heatLoss: current.heatLoss + grid[nr][nc],
					n:        current.n + 1,
					index:    0,
				})
			}
		}

		for _, deltas := range [4][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			if deltas != [2]int{current.rowD, current.colD} && deltas != [2]int{-current.rowD, -current.colD} {
				nr := current.row + deltas[0]
				nc := current.col + deltas[1]
				if 0 <= nr && nr < len(grid) && 0 <= nc && nc < len(grid[0]) {
					heap.Push(d, &node{
						row:      nr,
						col:      nc,
						rowD:     deltas[0],
						colD:     deltas[1],
						heatLoss: current.heatLoss + grid[nr][nc],
						n:        1,
					})
				}
			}
		}
	}

	return -1
}
