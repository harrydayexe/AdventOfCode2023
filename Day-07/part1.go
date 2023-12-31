package main

import (
	"slices"
)

func part1(hands []string) int {
	allHands := inputToHands(hands, false)

	slices.SortFunc(allHands, roundSortFunc)

	var returnTotal = 0
	for i, hand := range allHands {
		returnTotal += hand.bet * (i + 1)
	}

	return returnTotal
}
