package main

func part1(games []string) int {
	const MAX_RED = 12
	const MAX_GREEN = 13
	const MAX_BLUE = 14

	var result = 0

	var cleanedGames = cleanData(games)
	for _, cleanedGame := range cleanedGames {
		var gameMaxRed, gameMaxGreen, gameMaxBlue = 0, 0, 0
		for _, h := range cleanedGame.hands {
			gameMaxRed = max(gameMaxRed, h.numOfRed)
			gameMaxGreen = max(gameMaxGreen, h.numOfGreen)
			gameMaxBlue = max(gameMaxBlue, h.numOfBlue)
		}
		if gameMaxRed <= MAX_RED && gameMaxGreen <= MAX_GREEN && gameMaxBlue <= MAX_BLUE {
			result += cleanedGame.id
		}
	}
	return result
}
