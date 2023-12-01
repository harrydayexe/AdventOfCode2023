package main

import (
	"fmt"
	"strings"
)

func part2(lines []string) int {
	for i := range lines {
		translateValue(&lines[i])
	}

	return part1(lines)
}

func translateValue(val *string) {
	writtenNumbers := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	for _, number := range writtenNumbers {
		if strings.Contains(*val, number) {
			*val = strings.ReplaceAll(*val, number, fmt.Sprintf("%s%s%s", number, writtenNumberToNumeral(number), number))
		}
	}
	for _, number := range writtenNumbers {
		if strings.Contains(*val, number) {
			*val = strings.ReplaceAll(*val, number, "")
		}
	}
}

func writtenNumberToNumeral(number string) string {
	translation := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	return translation[number]
}
