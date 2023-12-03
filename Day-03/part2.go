package main

import (
	"strconv"
	"unicode"
)

func part2(grid []string) int {
	var total = 0

	for rowIndex, row := range grid {
		for colIndex, ch := range row {
			if ch != '*' {
				continue
			}
			var coordinates = make(map[pair]struct{})
			for _, r := range [3]int{rowIndex - 1, rowIndex, rowIndex + 1} {
				if r >= 0 && r < len(grid) {
					for _, c := range [3]int{colIndex - 1, colIndex, colIndex + 1} {
						if c >= 0 && c < len(grid[r]) && unicode.IsDigit(rune(grid[r][c])) {
							for c > 0 && unicode.IsDigit(rune(grid[r][c-1])) {
								c -= 1
							}
							coordinates[pair{x: c, y: r}] = struct{}{}
						}
					}
				}
			}
			if len(coordinates) != 2 {
				continue
			}

			var numbers []int
			for pair, _ := range coordinates {
				s := ""
				for pair.x < len(grid[pair.y]) && unicode.IsDigit(rune(grid[pair.y][pair.x])) {
					s = s + string(grid[pair.y][pair.x])
					pair.x += 1
				}
				i, _ := strconv.Atoi(s)
				numbers = append(numbers, i)
			}

			total += numbers[0] * numbers[1]
		}
	}

	return total
}
