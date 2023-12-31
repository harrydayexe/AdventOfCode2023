package main

func part1(lines []string) int {
	data := cleanData(lines)

	var runningTotal = 0
	for _, series := range data {
		runningTotal += extrapolate1(series)
	}

	return runningTotal
}

func extrapolate1(series []int) int {
	if isAllZero(series) {
		return 0
	}

	deltas := calcDelta(series)
	diff := extrapolate1(deltas)
	return diff + series[len(series)-1]
}
