package main

import (
	"strconv"
	"strings"
)

func part1(lines []string) int {
	data := cleanData(lines)

	var runningTotal = 0
	for _, series := range data {
		runningTotal += extrapolate(series)
	}

	return runningTotal
}

func extrapolate(series []int) int {
	if isAllZero(series) {
		return 0
	}

	deltas := calcDelta(series)
	diff := extrapolate(deltas)
	return diff + series[len(series)-1]
}

func cleanData(lines []string) [][]int {
	var returnArr = make([][]int, len(lines))

	for i, line := range lines {
		for _, s := range strings.Fields(line) {
			num, _ := strconv.Atoi(s)
			returnArr[i] = append(returnArr[i], num)
		}
	}

	return returnArr
}

func isAllZero(data []int) bool {
	for _, datum := range data {
		if datum != 0 {
			return false
		}
	}
	return true
}

func calcDelta(series []int) []int {
	var deltas = make([]int, len(series)-1)

	for i := 0; i < len(series)-1; i++ {
		deltas[i] = series[i+1] - series[i]
	}

	return deltas
}
