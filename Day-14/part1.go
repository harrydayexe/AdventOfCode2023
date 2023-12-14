package main

func part1(lines []string) int {
	var totalWeights = 0
	for col := 0; col < len(lines[0]); col++ {
		// For each column
		var runningCount = 0
		for i := len(lines) - 1; i >= 0; i-- {
			if lines[i][col] == 'O' {
				runningCount += 1
			} else if lines[i][col] == '#' {
				for j := 1; j <= runningCount; j++ {
					totalWeights += len(lines) - i - j
				}
				runningCount = 0
			}
		}
		if runningCount > 0 {
			for j := 0; j < runningCount; j++ {
				totalWeights += len(lines) - j
			}
		}
	}
	return totalWeights
}
