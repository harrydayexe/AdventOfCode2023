package main

func part2(lines []string) int {
	data := cleanData(lines)

	var runningTotal = 0
	for _, series := range data {
		runningTotal += extrapolate2(series)
	}

	return runningTotal
}

func extrapolate2(series []int) int {
	if isAllZero(series) {
		return 0
	}

	deltas := calcDelta(series)
	diff := extrapolate2(deltas)
	return series[0] - diff
}
