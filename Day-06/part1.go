package main

import (
	"strconv"
	"strings"
)

func part1(lines []string) int {
	var times []int
	for _, s := range strings.Fields(lines[0][strings.Index(lines[0], ":")+1:]) {
		i, _ := strconv.Atoi(s)
		times = append(times, i)
	}
	var distances []int
	for _, s := range strings.Fields(lines[1][strings.Index(lines[1], ":")+1:]) {
		i, _ := strconv.Atoi(s)
		distances = append(distances, i)
	}

	var total = 1
	for i := 0; i < len(times); i++ {
		var ways = 0
		for j := 0; j < times[i]; j++ {
			if j*(times[i]-j) > distances[i] {
				ways += 1
			}
		}
		total *= ways
	}

	return total
}
