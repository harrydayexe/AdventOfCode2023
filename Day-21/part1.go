package main

import (
	"fmt"
	"strings"
)

func part1(lines []string) int {
	var startingCoords [2]int
	for i, line := range lines {
		index := strings.Index(line, "S")
		if index != -1 {
			startingCoords = [2]int{i, index}
			break
		}
	}

	var visitedLocations = make(map[[2]int]int)
	var deltas = [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	var queue = [][3]int{{startingCoords[0], startingCoords[1], -1}}

	for len(queue) > 0 {
		var dqItem = queue[0]
		coords := [2]int{dqItem[0], dqItem[1]}
		queue = queue[1:]

		visitedLocations[coords] = dqItem[2] + 1

		for _, delta := range deltas {
			newCoords := [2]int{coords[0] + delta[0], coords[1] + delta[1]}
			if 0 <= newCoords[0] && newCoords[0] < len(lines) && 0 <= newCoords[1] && newCoords[1] < len(lines[0]) {
				if lines[newCoords[0]][newCoords[1]] != '#' && dqItem[2]+1 < 6 {
					n, prs := visitedLocations[[2]int{newCoords[0], newCoords[1]}]
					if !prs || dqItem[2]+1 < n {
						queue = append(queue, [3]int{newCoords[0], newCoords[1], dqItem[2] + 1})
					}
				}
			}
		}
	}

	printGrid(lines, visitedLocations)
	return len(visitedLocations)
}

func printGrid(lines []string, visitedLocations map[[2]int]int) {
	var copyLines = make([][]rune, len(lines))
	for i, line := range lines {
		copyLines[i] = []rune(line)
	}
	for key := range visitedLocations {
		copyLines[key[0]][key[1]] = rune(visitedLocations[key] + '0')
	}

	for _, line := range copyLines {
		fmt.Println(string(line))
	}
}
