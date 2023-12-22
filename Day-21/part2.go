package main

import (
	"strings"
)

func part2(lines []string) int {
	const dimension = 5
	var expandedGrid = make([]string, 0, len(lines)*dimension)
	for i := 0; i < dimension; i++ {
		for _, line := range lines {
			expandedGrid = append(expandedGrid, strings.Repeat(strings.ReplaceAll(line, "S", "."), dimension))
		}
	}
	sr, sc := len(expandedGrid)/2, len(expandedGrid[0])/2
	expandedGrid[sr] = expandedGrid[sr][:sc] + "S" + expandedGrid[sr][sc+1:]

	v1 := part1(expandedGrid, 65)
	v2 := part1(expandedGrid, 196)
	v3 := part1(expandedGrid, 327)

	a := (v1 - 2*v2 + v3) / 2
	b := (-3*v1 + 4*v2 - v3) / 2
	c := v1
	const n = 202300
	return a*n*n + b*n + c
}
