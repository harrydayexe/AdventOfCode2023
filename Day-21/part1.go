package main

import (
	"strings"
)

func part1(lines []string, MaxSteps int) int {
	var startingCoords [2]int
	for i, line := range lines {
		index := strings.Index(line, "S")
		if index != -1 {
			startingCoords = [2]int{i, index}
			break
		}
	}

	var visitedLocations = map[[2]int]struct{}{
		startingCoords: {},
	}
	var answers = make(map[[2]int]struct{})
	var deltas = [4][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	var queue = [][3]int{{startingCoords[0], startingCoords[1], MaxSteps}}

	for len(queue) > 0 {
		var dqItem = queue[0]
		queue = queue[1:]
		coords := [2]int{dqItem[0], dqItem[1]}

		if dqItem[2]%2 == 0 {
			answers[coords] = struct{}{}
		}
		if dqItem[2] == 0 {
			continue
		}

		for _, delta := range deltas {
			newCoords := [2]int{coords[0] + delta[0], coords[1] + delta[1]}
			_, prs := visitedLocations[newCoords]
			if newCoords[0] < 0 || newCoords[0] >= len(lines) || newCoords[1] < 0 || newCoords[1] >= len(lines[0]) || lines[newCoords[0]][newCoords[1]] == '#' || prs {
				continue
			}
			visitedLocations[newCoords] = struct{}{}
			queue = append(queue, [3]int{newCoords[0], newCoords[1], dqItem[2] - 1})
		}
	}

	return len(answers)
}

//func printGrid(lines []string, visitedLocations map[[2]int][]int) {
//	var copyLines = make([][]rune, len(lines))
//	for i, line := range lines {
//		copyLines[i] = []rune(line)
//	}
//	for key, n := range visitedLocations {
//		if slices.Contains(n, MAX_STEPS) {
//			copyLines[key[0]][key[1]] = '0'
//		}
//	}
//
//	for _, line := range copyLines {
//		fmt.Println(string(line))
//	}
//}
