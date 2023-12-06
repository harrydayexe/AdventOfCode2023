package main

import (
	"strconv"
	"strings"
)

type puzzleInput struct {
	seeds                 map[int]struct{}
	seedToSoil            [][3]int
	soilToFertiliser      [][3]int
	fertiliserToWater     [][3]int
	waterToLight          [][3]int
	lightToTemperature    [][3]int
	temperatureToHumidity [][3]int
	humidityToLocation    [][3]int
}

func cleanData(lines []string) *puzzleInput {
	var seeds = make(map[int]struct{})
	for _, seed := range strings.Fields(lines[0][strings.Index(lines[0], ":")+1:]) {
		var i, err = strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		seeds[i] = struct{}{}
	}

	var p = puzzleInput{seeds: seeds}
	var sections = splitSections(lines)

	p.seedToSoil = sections[0]
	p.soilToFertiliser = sections[1]
	p.fertiliserToWater = sections[2]
	p.waterToLight = sections[3]
	p.lightToTemperature = sections[4]
	p.temperatureToHumidity = sections[5]
	p.humidityToLocation = sections[6]

	return &p
}

func splitSections(lines []string) [7][][3]int {
	var returnValue [7][][3]int
	var sectionNum = 0
	for i, line := range lines {
		if i <= 2 || line == "" {
			continue
		}
		if strings.Contains(line, ":") {
			sectionNum += 1
			continue
		}

		var stringNums = strings.Fields(line)
		var i1, i2, i3 int
		i1, _ = strconv.Atoi(stringNums[0])
		i2, _ = strconv.Atoi(stringNums[1])
		i3, _ = strconv.Atoi(stringNums[2])
		var arrOfNums = [3]int{i1, i2, i3}
		returnValue[sectionNum] = append(returnValue[sectionNum], arrOfNums)
	}
	return returnValue
}

func convertMap(num int, ranges [][3]int) int {
	for _, r := range ranges {
		if r[1] <= num && num < r[1]+r[2] {
			return num - r[1] + r[0]
		}
	}
	return num
}
