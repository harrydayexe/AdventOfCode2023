package main

func part2(cards []string) int {
	var games = cleanData(cards)
	var gameResultMap = make(map[int]int)

	for _, game := range games {
		var gameTotal = 0
		for num := range game.actual {
			_, ok := game.winners[num]
			if ok {
				gameTotal += 1
			}
		}
		gameResultMap[game.id] = gameTotal
	}

	var gameTotalMap = make(map[int]int)
	for i := len(games); i > 0; i-- {
		res := gameResultMap[i]
		var totalToAdd = 1
		for j := i; (j < i+res) && (j < len(games)); j++ {
			totalToAdd += gameTotalMap[j+1]
		}
		gameTotalMap[i] = totalToAdd
	}

	var total = 0
	for _, i := range gameTotalMap {
		total += i
	}

	return total
}
