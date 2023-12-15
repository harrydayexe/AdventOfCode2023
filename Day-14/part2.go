package main

import (
	"fmt"
	"slices"
	"strings"
)

func part2(lines []string) int {
	var seen = make(map[string]struct{})
	var array []string
	var iter = 0

	var grid = make([]string, len(lines))
	copy(grid, lines)

	for !contains(seen, grid) {
		iter += 1
		hashed := toString(grid)
		seen[hashed] = struct{}{}
		array = append(array, hashed)
		grid = cycle(grid)
	}

	for _, s := range grid {
		fmt.Println(s)
	}
	var first = slices.Index(array, toString(grid))

	grid = strings.Split(array[(1000000000-first)%(iter-first)+first], ",")
	return part1(grid)
}

func contains(m map[string]struct{}, grid []string) bool {
	_, prs := m[toString(grid)]
	return prs
}

func cycle(grid []string) []string {
	var byteSlice = make([][]rune, len(grid))
	for i, s := range grid {
		byteSlice[i] = []rune(s)
	}
	rollNorth(byteSlice)
	rollWest(byteSlice)
	rollSouth(byteSlice)
	rollEast(byteSlice)

	var returnVal = make([]string, len(byteSlice))
	for i, runes := range byteSlice {
		returnVal[i] = string(runes)
	}
	return returnVal
}

func toString(grid []string) string {
	return strings.Join(grid, ",")
}

func printRunes(runes [][]rune) {
	fmt.Println()
	for _, line := range runes {
		fmt.Println(string(line))
	}
}

func rollNorth(grid [][]rune) {
	for row := 0; row < len(grid); row++ {
		for col, val := range grid[row] {
			if val == 'O' {
				for i := row - 1; i >= 0 && grid[i][col] == '.'; i-- {
					grid[i+1][col], grid[i][col] = grid[i][col], grid[i+1][col]
				}
			}
		}
	}
}

func rollWest(grid [][]rune) {
	for col := 0; col < len(grid[0]); col++ {
		for row, line := range grid {
			if line[col] == 'O' {
				for i := col - 1; i >= 0 && grid[row][i] == '.'; i-- {
					grid[row][i+1], grid[row][i] = grid[row][i], grid[row][i+1]
				}
			}
		}
	}
}

func rollSouth(grid [][]rune) {
	for row := len(grid) - 1; row >= 0; row-- {
		for col, val := range grid[row] {
			if val == 'O' {
				for i := row + 1; i < len(grid) && grid[i][col] == '.'; i++ {
					grid[i-1][col], grid[i][col] = grid[i][col], grid[i-1][col]
				}
			}
		}
	}
}

func rollEast(grid [][]rune) {
	for col := len(grid[0]) - 1; col >= 0; col-- {
		for row, line := range grid {
			if line[col] == 'O' {
				for i := col + 1; i < len(grid[0]) && grid[row][i] == '.'; i++ {
					grid[row][i-1], grid[row][i] = grid[row][i], grid[row][i-1]
				}
			}
		}
	}
}
