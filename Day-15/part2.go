package main

import (
	orderedmap "github.com/wk8/go-ordered-map"
	"strconv"
	"strings"
)

func part2(line string) int {
	var steps = strings.Split(line, ",")
	var boxes [256]*orderedmap.OrderedMap
	for i := 0; i < 256; i++ {
		boxes[i] = orderedmap.New()
	}

	for _, step := range steps {
		opIndex := strings.IndexAny(step, "=-")
		hash := hashStep(step[:opIndex])
		if step[opIndex] == '-' {
			boxes[hash].Delete(step[:opIndex])
		} else {
			focalLength, _ := strconv.Atoi(step[opIndex+1:])
			boxes[hash].Set(step[:opIndex], focalLength)
		}
	}

	var totalFocusingPower = 0
	for i, box := range boxes {
		slot := 1
		for pair := box.Oldest(); pair != nil; pair = pair.Next() {
			totalFocusingPower += (1 + i) * slot * pair.Value.(int)
			slot += 1
		}
	}

	return totalFocusingPower
}
