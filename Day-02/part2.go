package main

func part2(games []string) int {
	var result = 0

	var cleanedGames = cleanData(games)
	for _, cleanedGame := range cleanedGames {
		var gameMaxRed, gameMaxGreen, gameMaxBlue = 0, 0, 0
		for _, h := range cleanedGame.hands {
			gameMaxRed = max(gameMaxRed, h.numOfRed)
			gameMaxGreen = max(gameMaxGreen, h.numOfGreen)
			gameMaxBlue = max(gameMaxBlue, h.numOfBlue)
		}
		result += gameMaxRed * gameMaxGreen * gameMaxBlue
	}
	return result
}
