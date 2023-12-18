package main

import (
	"container/heap"
	"math"
)

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

type node struct {
	row, col    int
	heatLoss    int
	cost        int // Heat loss so far + own heatloss
	heuristic   int
	index       int
	n           int // How many in same direction
	inDirection Direction
	prev        *node
}

var deltaMap = map[Direction][2]int{
	North: {1, 0},
	South: {-1, 0},
	East:  {0, 1},
	West:  {0, -1},
}

func createGrid(lines []string) [][]*node {
	var grid = make([][]*node, len(lines))
	for i := range grid {
		grid[i] = make([]*node, len(lines[0]))
	}

	for row, line := range lines {
		for col, ch := range line {
			grid[row][col] = &node{
				row:         row,
				col:         col,
				heatLoss:    int(ch - '0'),
				cost:        0,
				heuristic:   int(math.Abs(float64(row-(len(grid)-1))) + math.Abs(float64(col-(len(grid[0])-1)))),
				index:       0,
				n:           0,
				inDirection: 0,
				prev:        nil,
			}
		}
	}

	return grid
}

func aStar(start, goal *node, grid [][]*node) {
	var d = &priorityQueue{}
	var dMap = make(map[[2]int]struct{})
	var f = make(map[[2]int]struct{})
	heap.Init(d)
	heap.Push(d, start)
	dMap[[2]int{start.row, start.col}] = struct{}{}

	for d.Len() > 0 {
		current := heap.Pop(d).(*node)
		delete(dMap, [2]int{current.row, current.col})
		if current == goal {
			return
		}
		f[[2]int{current.row, current.col}] = struct{}{}
		// For each neighbour
		if current.n < 3 {
			row, col := current.row+deltaMap[current.inDirection][0], current.col+deltaMap[current.inDirection][1]
			if 0 <= row && row < len(grid) && 0 <= col && col < len(grid[0]) {
				coord := [2]int{row, col}
				_, prs := f[coord]
				_, prs2 := dMap[coord]
				if !prs && !prs2 {
					// Case where neighbour is undiscovered
					// Add to d and set cost and prev to using current
					newNode := grid[row][col]
					newNode.n = current.n + 1
					newNode.inDirection = current.inDirection
					newNode.prev = current
					newNode.cost = newNode.heatLoss + current.cost
					heap.Push(d, newNode)
					dMap[coord] = struct{}{}
				} else if !prs {
					// case where node is discovered but not finished
					// Update cost if lower
					newNode := grid[row][col]
					if current.cost+newNode.heatLoss < newNode.cost {
						d.update(newNode, current.cost+newNode.heatLoss)
						newNode.prev = current
						newNode.n = current.n + 1
						newNode.inDirection = current.inDirection
					}
				}
			}
		}
		leftDir := (current.inDirection - 1) % 4
		rightDir := (current.inDirection + 1) % 4
		coords := [2][3]int{
			{current.row + deltaMap[leftDir][0], current.col + deltaMap[leftDir][1], int(leftDir)},
			{current.row + deltaMap[rightDir][0], current.col + deltaMap[rightDir][1], int(rightDir)},
		}
		for _, coord := range coords {
			if 0 <= coord[0] && coord[0] < len(grid) && 0 <= coord[1] && coord[1] < len(grid[0]) {
				// if coord is valid on grid
				_, prs := f[[2]int{coord[0], coord[1]}]
				_, prs2 := dMap[[2]int{coord[0], coord[1]}]
				if !prs && !prs2 {
					// Case where neighbour is undiscovered
					// Add to d and set cost and prev to using current
					newNode := grid[coord[0]][coord[1]]
					newNode.n = 1
					newNode.inDirection = Direction(coord[2])
					newNode.prev = current
					heap.Push(d, newNode)
					dMap[[2]int{coord[0], coord[1]}] = struct{}{}
				} else if !prs {
					// case where node is discovered but not finished
					// Update cost if lower
					newNode := grid[coord[0]][coord[1]]
					if current.cost+newNode.heatLoss < newNode.cost {
						d.update(newNode, current.cost+newNode.heatLoss)
						newNode.n = 1
						newNode.inDirection = Direction(coord[2])
						newNode.prev = current
					}
				}
			}
		}
	}
	//make an openlist containing only the starting node
	//make an empty closed list
	//while (the destination node has not been reached):
	//consider the node with the lowest f score in the open list
	//if (this node is our destination node) :
	//we are finished
	//if not:
	//put the current node in the closed list and look at all of its neighbors
	//for (each neighbor of the current node):
	//if (neighbor has lower g value than current and is in the closed list) :
	//replace the neighbor with the new, lower, g value
	//current node is now the neighbor's parent
	//else if (current g value is lower and this neighbor is in the open list ) :
	//replace the neighbor with the new, lower, g value
	//change the neighbor's parent to our current node
	//
	//else if this neighbor is not in both lists:
	//add it to the open list and set its g
}
