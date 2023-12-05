package main

func part1(lines []string) int {
	var p = cleanData(lines)
	var lowestLocation = int(^uint(0) >> 1)

	for seed, _ := range p.seeds {
		soil, prs := p.seedToSoil[seed]
		if !prs {
			soil = seed
		}
		fertiliser, prs := p.soilToFertiliser[soil]
		if !prs {
			fertiliser = soil
		}
		water, prs := p.fertiliserToWater[fertiliser]
		if !prs {
			water = fertiliser
		}
		light, prs := p.waterToLight[water]
		if !prs {
			light = water
		}
		temperature, prs := p.lightToTemperature[light]
		if !prs {
			temperature = light
		}
		humidity, prs := p.temperatureToHumidity[temperature]
		if !prs {
			humidity = temperature
		}
		location, prs := p.humidityToLocation[humidity]
		if !prs {
			location = humidity
		}
		lowestLocation = min(lowestLocation, location)
	}

	return lowestLocation
}
