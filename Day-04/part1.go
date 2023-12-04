package main

import (
	"math"
	"strconv"
	"strings"
)

type game struct {
	id      int
	winners map[int]struct{}
	actual  map[int]struct{}
}

func part1(cards []string) int {
	var games = cleanData(cards)
	var total = 0

	for _, game := range games {
		var gameTotal = 0
		for num := range game.actual {
			_, ok := game.winners[num]
			if ok {
				gameTotal += 1
			}
		}
		if gameTotal == 0 {
			continue
		}
		total += int(math.Pow(2, float64(gameTotal-1)))
	}

	return total
}

func cleanData(cards []string) []game {
	var returnValue []game
	for i, card := range cards {
		winnersString := card[strings.Index(card, ":")+1 : strings.Index(card, "|")]
		gameString := card[strings.Index(card, "|")+1 : len(card)]

		winners := strings.Fields(winnersString)
		games := strings.Fields(gameString)

		var g = game{id: i, winners: make(map[int]struct{}), actual: make(map[int]struct{})}
		for j := 0; j < len(winners); j++ {
			k, err := strconv.Atoi(winners[j])
			if err != nil {
				panic(err)
			}
			g.winners[k] = struct{}{}
		}
		for j := 0; j < len(games); j++ {
			k, err := strconv.Atoi(games[j])
			if err != nil {
				panic(err)
			}
			g.actual[k] = struct{}{}
		}
		returnValue = append(returnValue, g)
	}
	return returnValue
}
