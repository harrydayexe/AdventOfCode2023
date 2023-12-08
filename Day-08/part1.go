package main

func part1(lines []string) int {
	maps := parseInput(lines)
	directions := lines[0]

	var totalTurns = 0
	var i = 0
	var MAX_I = len(lines[0])
	var current = "AAA"
	for current != "ZZZ" {
		if i >= MAX_I {
			i = 0
		}
		if directions[i] == 'L' {
			current = maps.left[current]
		} else {
			current = maps.right[current]
		}
		i += 1
		totalTurns += 1
	}

	return totalTurns
}
