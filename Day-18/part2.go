package main

import (
	"math"
	"strconv"
	"strings"
)

func part2(lines []string) int {
	var points = [][2]int{{0, 0}}
	var b int

	var deltaMap = map[rune][2]int{
		'3': {-1, 0},
		'1': {1, 0},
		'2': {0, -1},
		'0': {0, 1},
	}

	for _, line := range lines {
		components := strings.Fields(line)
		nu, _ := strconv.ParseInt(components[2][2:7], 16, 64)
		n := int(nu)
		b += n
		lastCoord := points[len(points)-1]
		delta := deltaMap[rune(components[2][7])]
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
