package main

import (
	"strings"
	"unicode"
)

func part1(lines []string) int {
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

	var low, high = 0, 0
	for i := 0; i < 1000; i++ {
		low += 1
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

			if dqItem.pulseType == Low {
				low += 1
			} else if dqItem.pulseType == High {
				high += 1
			}

			_, prs := modules[dqItem.destination]
			if !prs {
				continue
			}

			nextMod := modules[dqItem.destination]
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

	return low * high
}
