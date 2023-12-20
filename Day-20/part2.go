package main

import (
	"slices"
	"strings"
	"unicode"
)

func part2(lines []string) int {
	var modules = make(map[string]*module)
	var broadcastTargets []string
	for _, line := range lines {
		if unicode.IsLetter(rune(line[0])) {
			broadcastTargets = strings.FieldsFunc(line[strings.Index(line, ">")+1:], func(r rune) bool {
				return !unicode.IsLetter(r)
			})
		} else {
			mod := newModule(line)
			modules[mod.name] = mod
		}
	}

	for name, mod := range modules {
		for _, output := range mod.destinations {
			_, prs := modules[output]
			if prs && modules[output].modType == Conjunction {
				modules[output].conjunctionMemory[name] = Low
			}
		}
	}

	var feed string
	for s, m := range modules {
		if slices.Contains(m.destinations, "rx") {
			feed = s
			break
		}
	}

	var cycleLengths = make(map[string]int)
	var seen = make(map[string]int)
	for s, m := range modules {
		if slices.Contains(m.destinations, feed) {
			seen[s] = 0
		}
	}

	var presses = 0
	for {
		presses += 1
		q := make([]*queueItem, len(broadcastTargets))
		for i2, target := range broadcastTargets {
			q[i2] = &queueItem{
				origin:      "broadcaster",
				destination: target,
				pulseType:   Low,
			}
		}

		for len(q) > 0 {
			dqItem := q[0]
			q = q[1:]

			nextMod, prs := modules[dqItem.destination]
			if !prs {
				continue
			}

			if nextMod.name == feed && dqItem.pulseType == High {
				seen[dqItem.origin] += 1
				_, prs := cycleLengths[dqItem.origin]
				if !prs {
					cycleLengths[dqItem.origin] = presses
				}

				allSet := true
				for _, i := range seen {
					if i == 0 {
						allSet = false
						break
					}
				}
				if allSet {
					x := 1
					for _, i := range cycleLengths {
						x = LCM(x, i)
					}
					return x
				}
			}

			if nextMod.modType == FlipFlop {
				if dqItem.pulseType == Low {
					nextMod.flipFlopMemory = !nextMod.flipFlopMemory
					outgoing := Low
					if nextMod.flipFlopMemory {
						outgoing = High
					}
					for _, destination := range nextMod.destinations {
						q = append(q, &queueItem{nextMod.name, destination, outgoing})
					}
				}
			} else {
				nextMod.conjunctionMemory[dqItem.origin] = dqItem.pulseType
				outgoing := High
				if nextMod.isAllInputHigh() {
					outgoing = Low
				}
				for _, destination := range nextMod.destinations {
					q = append(q, &queueItem{nextMod.name, destination, outgoing})
				}
			}
		}
	}

	return -1
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
