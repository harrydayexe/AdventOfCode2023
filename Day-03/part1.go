package main

import (
	"strconv"
	"unicode"
)

type pair struct {
	x int
	y int
}

func part1(grid []string) int {
	var coordinates = make(map[pair]struct{})

	for rowIndex, row := range grid {
		for colIndex, ch := range row {
			if unicode.IsDigit(ch) || ch == '.' {
				continue
			}
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
		}
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

	var returnValue int = 0
	for _, number := range numbers {
		returnValue += number
	}
	return returnValue
}
