package main

func part1(lines []string) int {
	var p = cleanData(lines)
	var lowestLocation = int(^uint(0) >> 1)

	for seed, _ := range p.seeds {
		soil := convertMap(seed, p.seedToSoil)
		fertiliser := convertMap(soil, p.soilToFertiliser)
		water := convertMap(fertiliser, p.fertiliserToWater)
		light := convertMap(water, p.waterToLight)
		temperature := convertMap(light, p.lightToTemperature)
		humidity := convertMap(temperature, p.temperatureToHumidity)
		location := convertMap(humidity, p.humidityToLocation)

		lowestLocation = min(lowestLocation, location)
	}

	return lowestLocation
}
