package main

func part1(lines []string) int {
	rows := cleanData(lines)

	var total = 0
	for _, row := range rows {
		count := countPossibleCombinations(row)
		total += count
	}
	return total
}

func countPossibleCombinations(row HotSpringRow) int {
	if row.isValid() {
		return 1
	}
	if row.countHotSprings(Unknown)+row.countHotSprings(Damaged) < row.countTargetTotalDamaged() {
		// If the number of unknowns is not enough to make the total required then prune
		return 0
	}
	if row.countHotSprings(Unknown) == 0 && !row.isValid() {
		// If all the unknowns have been set and the row is invalid, then this is not a valid solution
		return 0
	}

	return countPossibleCombinations(row.setNextUnknown(Damaged)) + countPossibleCombinations(row.setNextUnknown(Working))
}
