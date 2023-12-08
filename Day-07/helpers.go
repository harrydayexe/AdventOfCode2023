package main

import (
	"bytes"
	"slices"
	"strconv"
	"strings"
)

type HandType int

const (
	HighCard HandType = 1 + iota
	OnePair
	TwoPair
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

type Hand struct {
	handType HandType
	hand     []byte
	bet      int
}

func roundSortFunc(a, b Hand) int {
	// first sort by hand type
	if a.handType != b.handType {
		return int(a.handType) - int(b.handType)
	}

	// then sort by the card values, left->right (we can rely on ASCII compare)
	return bytes.Compare(a.hand, b.hand)
}

func inputToHands(lines []string, wildcard bool) []Hand {
	var allHands []Hand

	jsub := byte('B')
	if wildcard {
		jsub = '1'
	}

	for _, hand := range lines {
		ht := handColouring(hand[:strings.Index(hand, " ")], wildcard)
		i, _ := strconv.Atoi(hand[strings.Index(hand, " ")+1:])
		handBytes := []byte(hand[:strings.Index(hand, " ")])
		for j, c := range handBytes {
			switch c {
			case 'T':
				handBytes[j] = 'A'
			case 'J':
				handBytes[j] = jsub
			case 'Q':
				handBytes[j] = 'C'
			case 'K':
				handBytes[j] = 'D'
			case 'A':
				handBytes[j] = 'E'
			default:
				// leave as-is
			}
		}

		allHands = append(allHands, Hand{
			handType: ht,
			hand:     handBytes,
			bet:      i,
		})
	}

	return allHands
}

func handColouring(hand string, wildcard bool) HandType {
	if hand == "JJJJJ" {
		return FiveOfAKind
	}
	var handCopy = strings.Clone(hand)

	if wildcard {
		handCopy = strings.ReplaceAll(handCopy, "J", "")
	}

	var seen = make(map[rune]int)
	for _, ch := range handCopy {
		i, hasBeenSeen := seen[ch]
		if hasBeenSeen {
			seen[ch] = i + 1
		} else {
			seen[ch] = 1
		}
	}
	var counts []int
	for _, i := range seen {
		counts = append(counts, i)
	}
	slices.Sort(counts)
	slices.Reverse(counts)

	if len(handCopy) < len(hand) {
		var numOfJoker = len(hand) - len(handCopy)
		counts[0] += numOfJoker
	}

	if len(counts) == 1 {
		return FiveOfAKind
	} else if len(counts) == 2 {
		if counts[0] == 4 {
			return FourOfAKind
		} else {
			return FullHouse
		}
	} else if len(counts) == 3 {
		if counts[0] == 3 {
			return ThreeOfAKind
		} else {
			return TwoPair
		}
	} else if len(counts) == 5 {
		return HighCard
	} else {
		return OnePair
	}
}
