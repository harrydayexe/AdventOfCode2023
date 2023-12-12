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

func (row *HotSpringRow) countHotSprings() (int, int, int) {
	var countUnknown, countWorking, countDamaged int
	for _, record := range row.HotSpringRecords {
		if record == Unknown {
			countUnknown += 1
		} else if record == Working {
			countWorking += 1
		} else {
			countDamaged += 1
		}
	}
	return countUnknown, countWorking, countDamaged
}

func (row *HotSpringRow) countTargetTotalDamaged() int {
	var count = 0
	for _, group := range row.HotSpringGroups {
		count += group
	}
	return count
}

func (row *HotSpringRow) countTotalNeeded() int {
	var count = 0
	for _, group := range row.HotSpringGroups {
		count += group
	}
	count += len(row.HotSpringGroups) - 1
	return count
}

func (row *HotSpringRow) isValid() bool {
	var groupIndex = 0
	var groupRunningCount = 0
	for _, spring := range row.HotSpringRecords {
		if spring == Damaged {
			groupRunningCount += 1
		} else if groupRunningCount != 0 {
			if groupIndex >= len(row.HotSpringGroups) || row.HotSpringGroups[groupIndex] != groupRunningCount {
				return false
			}
			groupRunningCount = 0
			groupIndex += 1
		}
	}
	if groupRunningCount != 0 {
		if groupIndex >= len(row.HotSpringGroups) || row.HotSpringGroups[groupIndex] != groupRunningCount {
			return false
		}
	}
	return true
}

func (row *HotSpringRow) shouldBePruned() bool {
	remainingNeeded := row.countTotalNeeded()
	totalLength := len(row.HotSpringRecords) - 1

	var currentGroupLength = 0
	var groupIndex = 0
	for i, record := range row.HotSpringRecords {
		if record == Damaged {
			currentGroupLength += 1
			remainingNeeded -= 1
			if groupIndex == len(row.HotSpringGroups) || currentGroupLength > row.HotSpringGroups[groupIndex] {
				return true
			}
		} else if record == Working {
			if currentGroupLength > 0 {
				// First working spring after (contiguous group of) damaged
				if currentGroupLength != row.HotSpringGroups[groupIndex] {
					// Group does not match
					return true
				} else {
					// Group matches, ignore and move on
					currentGroupLength = 0
					groupIndex += 1
					remainingNeeded -= 1
				}
			}
			// Previous was also working
			if totalLength-i < remainingNeeded {
				// If this Working spring needed to be damaged to be able to fit the remainder of the pattern in
				return true
			}
		} else {
			return false
		}
		// keep track of length of contiguous damaged and compare against group
		// keep track of how far from the end and what groups are still needed
	}
	return false
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

func countPossibleCombinations(row HotSpringRow) int {
	countUnknown, _, countDamaged := row.countHotSprings()

	if countUnknown+countDamaged < row.countTargetTotalDamaged() {
		// If the number of unknowns is not enough to make the total required then prune
		return 0
	}

	if countUnknown == 0 {
		// If all the unknowns have been set check if the row is valid
		if row.isValid() {
			return 1
		} else {
			return 0
		}
	}

	// At this point, there are enough unknowns to fulfill the criteria
	// Now check if the shape of the row will fit the format or can be pruned early
	if row.shouldBePruned() {
		return 0
	}

	return countPossibleCombinations(row.setNextUnknown(Damaged)) + countPossibleCombinations(row.setNextUnknown(Working))
}
