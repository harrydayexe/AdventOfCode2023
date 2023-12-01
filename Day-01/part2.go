package main

import (
	"regexp"
)

func part2(lines []string) int {
	convertToInt := func(match string) int {
		digitMap := map[string]int{
			"one":   1,
			"two":   2,
			"three": 3,
			"four":  4,
			"five":  5,
			"six":   6,
			"seven": 7,
			"eight": 8,
			"nine":  9,
		}
		val, ok := digitMap[match]
		if ok {
			return val
		}

		return int(match[0] - '0')
	}

	count := 0

	r, _ := regexp.Compile("(\\d)|(one)|(two)|(three)|(four)|(five)|(six)|(seven)|(eight)|(nine)")

	for _, line := range lines {
		match := r.FindAllString(line, -1)
		if len(match) != 0 {
			count += convertToInt(match[0]) * 10
			count += convertToInt(match[len(match)-1])
		}
	}

	return count
}
