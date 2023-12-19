package main

func part2(lines []string) int {
	var extractedWorkflows = extractWorkflows(lines)
	var inRanges = map[rune][2]int{
		'x': {1, 4000},
		'm': {1, 4000},
		'a': {1, 4000},
		's': {1, 4000},
	}
	return count(inRanges, "in", extractedWorkflows)
}

func count(ranges map[rune][2]int, name string, wfs workflows) int {
	var rangesCopy = make(map[rune][2]int)
	for r, ints := range ranges {
		rangesCopy[r] = ints
	}
	if name == "R" {
		return 0
	}
	if name == "A" {
		product := 1
		for _, loHi := range rangesCopy {
			product *= loHi[1] - loHi[0] + 1
		}
		return product
	}

	rules := wfs[name]
	total := 0

	var allFound = false
	for _, r := range rules[:len(rules)-1] {
		lo, hi := rangesCopy[r.variable][0], rangesCopy[r.variable][1]
		var T, F [2]int
		if !r.isGreaterThan {
			T = [2]int{lo, r.value - 1}
			F = [2]int{r.value, hi}
		} else {
			T = [2]int{r.value + 1, hi}
			F = [2]int{lo, r.value}
		}

		if T[0] <= T[1] {
			mapCopy1 := make(map[rune][2]int)
			for r2, ints := range rangesCopy {
				mapCopy1[r2] = ints
			}
			mapCopy1[r.variable] = T
			total += count(mapCopy1, r.result, wfs)
		}
		if F[0] <= F[1] {
			for r2, ints := range rangesCopy {
				rangesCopy[r2] = ints
			}
			rangesCopy[r.variable] = F
		} else {
			allFound = true
			break
		}
	}
	if !allFound {
		total += count(rangesCopy, rules[len(rules)-1].result, wfs)
	}

	return total
}
