package main

import (
	"math"
	"strconv"
	"strings"
)

func part1(lines []string) int {
	var points = [][2]int{{0, 0}}
	var b int

	var deltaMap = map[string][2]int{
		"U": {-1, 0},
		"D": {1, 0},
		"L": {0, -1},
		"R": {0, 1},
	}

	for _, line := range lines {
		components := strings.Fields(line)
		n, _ := strconv.Atoi(components[1])
		b += n
		lastCoord := points[len(points)-1]
		delta := deltaMap[components[0]]
		newCoord := [2]int{lastCoord[0] + (delta[0] * n), lastCoord[1] + (delta[1] * n)}
		points = append(points, newCoord)
	}

	var a int
	for i := 0; i < len(points)-1; i++ {
		a += (points[i+1][0] + points[i][0]) * (points[i+1][1] - points[i][1])
	}
	a = int(math.Abs(float64(a))) / 2

	i := a - b/2 + 1
	return i + b
}
