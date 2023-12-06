package main

import (
	"strconv"
	"strings"
)

func part2(lines []string) int {
	time, _ := strconv.Atoi(strings.ReplaceAll(lines[0], " ", "")[strings.Index(lines[0], ":")+1:])
	distance, _ := strconv.Atoi(strings.ReplaceAll(lines[1], " ", "")[strings.Index(lines[1], ":")+1:])

	var total = 0

	for j := 0; j < time; j++ {
		if j*(time-j) > distance {
			total += 1
		}
	}

	return total
}
