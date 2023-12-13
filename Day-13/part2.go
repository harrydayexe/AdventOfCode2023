package main

func part2(lines []string) int {
	var patterns [][]string
	var strings []string
	for _, line := range lines {
		if line == "" {
			patterns = append(patterns, strings)
			strings = make([]string, 0)
		} else {
			strings = append(strings, line)
		}
	}
	patterns = append(patterns, strings)

	var rowsAbove, columnsLeft = 0, 0
	for _, pattern := range patterns {
		var potentialReflectionsHorizontal []int
		for i := 1; i < len(pattern); i++ {
			var charsDifferent = 0
			for j := 0; j < len(pattern[i]); j++ {
				if pattern[i][j] != pattern[i-1][j] {
					charsDifferent += 1
				}
			}
			if charsDifferent <= 1 {
				potentialReflectionsHorizontal = append(potentialReflectionsHorizontal, i-1)
			}
		}

		for _, index := range potentialReflectionsHorizontal {
			i1, i2 := index, index+1
			var charsDifferent = 0
			for i1 >= 0 && i2 < len(pattern) && charsDifferent <= 1 {
				for j := 0; j < len(pattern[i1]); j++ {
					if pattern[i1][j] != pattern[i2][j] {
						charsDifferent += 1
					}
				}
				i1 -= 1
				i2 += 1
			}
			if charsDifferent == 1 {
				rowsAbove += index + 1
			}
		}

		var potentialReflectionsVertical []int
		for i := 1; i < len(pattern[0]); i++ {
			var charsDifferent = 0
			for _, line := range pattern {
				if line[i] != line[i-1] {
					charsDifferent += 1
				}
			}
			if charsDifferent <= 1 {
				potentialReflectionsVertical = append(potentialReflectionsVertical, i-1)
			}
		}

		for _, index := range potentialReflectionsVertical {
			i1, i2 := index, index+1
			var charsDifferent = 0
			for i1 >= 0 && i2 < len(pattern[0]) && charsDifferent <= 1 {
				for _, line := range pattern {
					if line[i1] != line[i2] {
						charsDifferent += 1
					}
				}
				i1 -= 1
				i2 += 1
			}
			if charsDifferent == 1 {
				columnsLeft += index + 1
			}
		}
	}

	return columnsLeft + 100*rowsAbove
}
