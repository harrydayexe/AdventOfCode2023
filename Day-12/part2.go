package main

import "fmt"

func part2(lines []string) int {
	rows := cleanData(lines)
	for i := range rows {
		rows[i].unfoldRow()
	}

	var total = 0
	fmt.Println("Calculating part 2...")
	for i, row := range rows {
		fmt.Println(i)
		count := countPossibleCombinations(row)
		total += count
	}
	return total
}

func (row *HotSpringRow) unfoldRow() {
	var groupLength = len(row.HotSpringGroups)
	var repeatedGroups = make([]int, groupLength*5)
	for i := 0; i < 5; i++ {
		for j := 0; j < groupLength; j++ {
			repeatedGroups[i*groupLength+j] = row.HotSpringGroups[j]
		}
	}
	row.HotSpringGroups = repeatedGroups

	var recordLength = len(row.HotSpringRecords)
	var repeatedRecords = make([]HotSpring, (recordLength+1)*5-1)
	for i := 0; i < 5; i++ {
		for j := 0; j < recordLength; j++ {
			repeatedRecords[i+i*recordLength+j] = row.HotSpringRecords[j]
		}
	}
	row.HotSpringRecords = repeatedRecords
}
