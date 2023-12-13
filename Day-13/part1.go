package main

func part1(lines []string) int {
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
			if pattern[i] == pattern[i-1] {
				potentialReflectionsHorizontal = append(potentialReflectionsHorizontal, i-1)
			}
		}

		for _, index := range potentialReflectionsHorizontal {
			i1, i2 := index, index+1
			var isReflection = true
			for i1 >= 0 && i2 < len(pattern) && isReflection {
				if pattern[i1] != pattern[i2] {
					isReflection = false
				}
				i1 -= 1
				i2 += 1
			}
			if isReflection {
				rowsAbove += index + 1
			}
		}

		var potentialReflectionsVertical []int
		for i := 1; i < len(pattern[0]); i++ {
			var foundReflection = true
			for _, line := range pattern {
				if line[i] != line[i-1] {
					foundReflection = false
				}
			}
			if foundReflection {
				potentialReflectionsVertical = append(potentialReflectionsVertical, i-1)
			}
		}

		for _, index := range potentialReflectionsVertical {
			i1, i2 := index, index+1
			var isReflection = true
			for i1 >= 0 && i2 < len(pattern[0]) && isReflection {
				for _, line := range pattern {
					if line[i1] != line[i2] {
						isReflection = false
					}
				}
				i1 -= 1
				i2 += 1
			}
			if isReflection {
				columnsLeft += index + 1
			}
		}
	}

	return columnsLeft + 100*rowsAbove
}
