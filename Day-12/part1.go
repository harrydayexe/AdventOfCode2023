package main

func part1(lines []string) int {
	rows := cleanData(lines)

	var total = 0
	for _, row := range rows {
		count := countPossibleCombinations(row)
		total += count
		//fmt.Println(count)
	}
	return total
}
