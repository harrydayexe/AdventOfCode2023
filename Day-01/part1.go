package main

import (
	"strconv"
	"strings"
)

func part1(lines []string) int {
	var cleanedData [][]int
	for _, line := range lines {
		var digits []int
		for _, char := range strings.Split(line, "") {
			if err, d := isInt(char); err {
				digits = append(digits, d)
			}
		}
		if len(digits) > 0 {
			cleanedData = append(cleanedData, digits)
		}
	}

	count := 0
	for _, digits := range cleanedData {
		count += (digits[0] * 10) + digits[len(digits)-1]
	}

	return count
}

func isInt(char string) (bool, int) {
	i, err := strconv.Atoi(char)
	return err == nil, i
}
