package main

import (
	"strings"
)

func part1(line string) int {
	var steps = strings.Split(line, ",")
	var total = 0
	for _, step := range steps {
		total += hashStep(step)
	}
	return total
}

func hashStep(step string) int {
	var total = 0
	for _, i := range step {
		total += int(i)
		total = (total * 17) % 256
	}
	return total
}
