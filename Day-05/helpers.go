package main

import (
	"strconv"
	"strings"
)

type puzzleInput struct {
	seeds                 map[int]struct{}
	seedToSoil            map[int]int
	soilToFertiliser      map[int]int
	fertiliserToWater     map[int]int
	waterToLight          map[int]int
	lightToTemperature    map[int]int
	temperatureToHumidity map[int]int
	humidityToLocation    map[int]int
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

	p.seedToSoil = makeMap(sections[0])
	p.soilToFertiliser = makeMap(sections[1])
	p.fertiliserToWater = makeMap(sections[2])
	p.waterToLight = makeMap(sections[3])
	p.lightToTemperature = makeMap(sections[4])
	p.temperatureToHumidity = makeMap(sections[5])
	p.humidityToLocation = makeMap(sections[6])

	return &p
}

func splitSections(lines []string) [7][][]int {
	var returnValue [7][][]int
	var sectionNum = 0
	for i, line := range lines {
		if i <= 2 || line == "" {
			continue
		}
		if strings.Contains(line, ":") {
			sectionNum += 1
			continue
		}

		var arrOfNums []int
		for _, num := range strings.Fields(line) {
			n, err := strconv.Atoi(num)
			if err != nil {
				panic(err)
			}
			arrOfNums = append(arrOfNums, n)
		}
		returnValue[sectionNum] = append(returnValue[sectionNum], arrOfNums)
	}
	return returnValue
}

func makeMap(params [][]int) map[int]int {
	var returnMap = make(map[int]int)

	for _, mapEntry := range params {
		for i := 0; i < mapEntry[2]; i++ {
			returnMap[mapEntry[1]+i] = mapEntry[0] + i
		}
	}

	return returnMap
}
