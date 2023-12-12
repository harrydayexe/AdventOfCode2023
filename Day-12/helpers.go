package main

import (
	"strconv"
	"strings"
)

type HotSpring int

const (
	Unknown HotSpring = iota
	Working
	Damaged
)

var runeToHotSpringMap = map[rune]HotSpring{
	'?': Unknown,
	'.': Working,
	'#': Damaged,
}

type HotSpringRow struct {
	HotSpringRecords []HotSpring
	HotSpringGroups  []int
}

func (row *HotSpringRow) countHotSprings(ofType HotSpring) int {
	var count = 0
	for _, record := range row.HotSpringRecords {
		if record == ofType {
			count += 1
		}
	}
	return count
}

func (row *HotSpringRow) countTargetTotalDamaged() int {
	var count = 0
	for _, group := range row.HotSpringGroups {
		count += group
	}
	return count
}

func (row *HotSpringRow) isValid() bool {
	if row.countHotSprings(Unknown) != 0 {
		return false
	}

	groups := groupDamagedHotSprings(row.HotSpringRecords)

	if len(groups) != len(row.HotSpringGroups) {
		return false
	}

	for i := 0; i < len(groups); i++ {
		if groups[i] != row.HotSpringGroups[i] {
			return false
		}
	}
	return true
}

func (row *HotSpringRow) setNextUnknown(toValue HotSpring) HotSpringRow {
	var newRow = HotSpringRow{
		HotSpringGroups:  make([]int, len(row.HotSpringGroups)),
		HotSpringRecords: make([]HotSpring, len(row.HotSpringRecords)),
	}
	copy(newRow.HotSpringRecords, row.HotSpringRecords)
	copy(newRow.HotSpringGroups, row.HotSpringGroups)

	for i, record := range newRow.HotSpringRecords {
		if record == Unknown {
			newRow.HotSpringRecords[i] = toValue
			return newRow
		}
	}
	return newRow
}

func groupDamagedHotSprings(springs []HotSpring) []int {
	var returnArray []int

	var groupRunningCount = 0
	for _, spring := range springs {
		if spring == Damaged {
			groupRunningCount += 1
		} else if groupRunningCount != 0 {
			returnArray = append(returnArray, groupRunningCount)
			groupRunningCount = 0
		}
	}
	if groupRunningCount != 0 {
		returnArray = append(returnArray, groupRunningCount)
	}

	return returnArray
}

func cleanData(lines []string) []HotSpringRow {
	var returnArray = make([]HotSpringRow, len(lines))

	for i, line := range lines {
		spaceIndex := strings.Index(line, " ")
		records, groups := toHotSpringSlice(line[:spaceIndex]), toIntSlice(strings.FieldsFunc(line[spaceIndex+1:], func(r rune) bool {
			return r == ','
		}))
		returnArray[i] = HotSpringRow{
			HotSpringRecords: records,
			HotSpringGroups:  groups,
		}
	}

	return returnArray
}

func toIntSlice(in []string) []int {
	var returnArray []int

	for _, s := range in {
		num, _ := strconv.Atoi(s)
		returnArray = append(returnArray, num)
	}

	return returnArray
}

func toHotSpringSlice(in string) []HotSpring {
	var returnArray = make([]HotSpring, len(in))

	for i, ch := range in {
		returnArray[i] = runeToHotSpringMap[ch]
	}

	return returnArray
}
