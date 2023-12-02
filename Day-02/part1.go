package main

import (
	"strconv"
	"strings"
)

type hand struct {
	numOfRed   int
	numOfGreen int
	numOfBlue  int
}

type game struct {
	hands []hand
	id    int
}

func part1(games []string) int {
	const MAX_RED = 12
	const MAX_GREEN = 13
	const MAX_BLUE = 14

	cleanData(games)
	return 0
}

func cleanData(gamesI []string) []game {
	var games []game
	for gameID, gameI := range gamesI {
		handsI := strings.Split(gameI[strings.IndexByte(gameI, ':')+1:], ";")
		var hands []hand
		for _, handI := range handsI {
			colours := strings.Fields(strings.ReplaceAll(handI, ",", ""))
			var numOfBlue, numOfGreen, numOfRed int = 0, 0, 0
			for i := 1; i < len(colours); i += 2 {
				if colours[i] == "blue" {
					numOfBlue, _ = strconv.Atoi(colours[i-1])
				} else if colours[i] == "green" {
					numOfGreen, _ = strconv.Atoi(colours[i-1])
				} else {
					numOfRed, _ = strconv.Atoi(colours[i-1])
				}
			}
			hands = append(hands, hand{numOfBlue: numOfBlue, numOfGreen: numOfGreen, numOfRed: numOfRed})
		}
		games = append(games, game{id: gameID, hands: hands})
	}
	return games
}
