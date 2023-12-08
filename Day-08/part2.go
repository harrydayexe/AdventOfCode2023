package main

func part2(lines []string) int {
	maps := parseInput(lines)
	directions := lines[0]

	var MAX_I = len(lines[0])
	var current = startingLocations(maps.left)
	var lengths []int

	for startPos := range current {
		var totalTurns = 0
		var i = 0

		for startPos[len(startPos)-1] != 'Z' {
			if i >= MAX_I {
				i = 0
			}
			if directions[i] == 'L' {
				startPos = maps.left[startPos]
			} else {
				startPos = maps.right[startPos]
			}
			i += 1
			totalTurns += 1
		}
		lengths = append(lengths, totalTurns)
	}

	var total = 1
	for _, length := range lengths {
		total = LCM(total, length)
	}
	return total
}

func isEnd(locations map[string]struct{}) bool {
	for loc, _ := range locations {
		if loc[len(loc)-1] != 'Z' {
			return false
		}
	}
	return true
}

func startingLocations(locations map[string]string) map[string]struct{} {
	var locsToReturn = make(map[string]struct{})

	for location := range locations {
		if location[len(location)-1] == 'A' {
			locsToReturn[location] = struct{}{}
		}
	}

	return locsToReturn
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
