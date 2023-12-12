package main

import (
	"strconv"
	"strings"
)

func part2(lines []string) int {
	var total = 0
	for _, line := range lines {
		var cfg, nums = strings.Fields(line)[0], toIntSlice(strings.FieldsFunc(strings.Fields(line)[1], func(r rune) bool {
			return r == ','
		}))

		cfg = strings.Join([]string{cfg, cfg, cfg, cfg, cfg}, "?")
		var groupLength = len(nums)
		var repeatedGroups = make([]int, groupLength*5)
		for i := 0; i < 5; i++ {
			for j := 0; j < groupLength; j++ {
				repeatedGroups[i*groupLength+j] = nums[j]
			}
		}
		total += count(cfg, repeatedGroups)
	}
	return total
}

type CacheKey struct {
	cfg  string
	nums string
}

var cache = make(map[CacheKey]int)

func count(cfg string, nums []int) int {
	if cfg == "" {
		if len(nums) == 0 {
			return 1
		} else {
			return 0
		}
	}

	if len(nums) == 0 {
		if strings.ContainsRune(cfg, '#') {
			return 0
		} else {
			return 1
		}
	}

	var key = CacheKey{
		cfg:  cfg,
		nums: convertNumsToString(nums),
	}

	res, prs := cache[key]
	if prs {
		return res
	}

	var result = 0

	if cfg[0] == '.' || cfg[0] == '?' {
		result += count(cfg[1:], nums)
	}

	if cfg[0] == '#' || cfg[0] == '?' {
		if nums[0] <= len(cfg) && !strings.ContainsRune(cfg[:nums[0]], '.') && (nums[0] == len(cfg) || cfg[nums[0]] != '#') {
			if nums[0] > len(nums) {
				result += count(cfg[nums[0]:], nums[1:])
			} else {
				result += count(cfg[nums[0]+1:], nums[1:])
			}
		}
	}

	cache[key] = result
	return result
}

func convertNumsToString(nums []int) string {
	var outputString string
	for _, num := range nums {
		outputString += "," + strconv.Itoa(num)
	}
	return outputString[1:]
}
