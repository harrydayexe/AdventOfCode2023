package main

import (
	"fmt"
	"strconv"
	"strings"
)

type puzzleInput1 struct {
	seeds                 map[int]struct{}
	seedToSoil            [][3]int
	soilToFertiliser      [][3]int
	fertiliserToWater     [][3]int
	waterToLight          [][3]int
	lightToTemperature    [][3]int
	temperatureToHumidity [][3]int
	humidityToLocation    [][3]int
}

type puzzleInput2 struct {
	seeds  [][2]int
	blocks [7][][3]int
}

func cleanData(lines []string) *puzzleInput1 {
	var seeds = make(map[int]struct{})
	for _, seed := range strings.Fields(lines[0][strings.Index(lines[0], ":")+1:]) {
		var i, err = strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		seeds[i] = struct{}{}
	}

	var p = puzzleInput1{seeds: seeds}
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

func cleanData2(lines []string) *puzzleInput2 {
	var seeds [][2]int
	inputsSeeds := strings.Fields(lines[0][strings.Index(lines[0], ":")+1:])
	for i := 0; i < len(inputsSeeds); i += 2 {
		i1, _ := strconv.Atoi(inputsSeeds[i])
		i2, _ := strconv.Atoi(inputsSeeds[i+1])
		seeds = append(seeds, [2]int{i1, i1 + i2})
	}

	var p = puzzleInput2{seeds: seeds}
	var sections = splitSections(lines)

	p.blocks[0] = sections[0]
	p.blocks[1] = sections[1]
	p.blocks[2] = sections[2]
	p.blocks[3] = sections[3]
	p.blocks[4] = sections[4]
	p.blocks[5] = sections[5]
	p.blocks[6] = sections[6]

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

func convertMap2(seeds [][2]int, ranges [][3]int) [][2]int {
	fmt.Println(seeds)
	var new [][2]int
	for len(seeds) > 0 {
		var start, end int
		start, end, seeds = seeds[0][0], seeds[0][1], seeds[:len(seeds)-1]
		var found = false
		for _, r := range ranges {
			overlapStart := max(start, r[1])
			overlapEnd := min(end, r[1]+r[2])
			if overlapStart < overlapEnd {
				found = true
				new = append(new, [2]int{overlapStart - r[1] + r[0], overlapEnd - r[1] + r[0]})
				if overlapStart > start {
					seeds = append(seeds, [2]int{start, overlapStart})
				}
				if end > overlapEnd {
					seeds = append(seeds, [2]int{end, overlapEnd})
				}
				break
			}
		}
		if !found {
			new = append(new, [2]int{start, end})
		}
	}
	return new
}
