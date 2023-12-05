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

func NewPuzzleInput(seeds []int) *puzzleInput {
	var seedsMap = make(map[int]struct{})
	for _, seed := range seeds {
		seedsMap[seed] = struct{}{}
	}
	var o = puzzleInput{
		seeds:                 seedsMap,
		seedToSoil:            make(map[int]int),
		soilToFertiliser:      make(map[int]int),
		fertiliserToWater:     make(map[int]int),
		waterToLight:          make(map[int]int),
		lightToTemperature:    make(map[int]int),
		temperatureToHumidity: make(map[int]int),
		humidityToLocation:    make(map[int]int)}
	return &o
}

func cleanData(lines []string) *puzzleInput {
	var seeds []int
	for _, seed := range strings.Fields(lines[0][strings.Index(lines[0], ":")+1:]) {
		var i, err = strconv.Atoi(seed)
		if err != nil {
			panic(err)
		}
		seeds = append(seeds, i)
	}

	var p = NewPuzzleInput(seeds)

	return p
}
