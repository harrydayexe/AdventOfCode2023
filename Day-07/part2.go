package main

import "slices"

func part2(hands []string) int {
	allHands := inputToHands(hands, true)

	slices.SortFunc(allHands, roundSortFunc)

	var returnTotal = 0
	for i, hand := range allHands {
		returnTotal += hand.bet * (i + 1)
	}

	return returnTotal
}
