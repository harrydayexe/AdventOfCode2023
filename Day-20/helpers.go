package main

import (
	"strings"
	"unicode"
)

type moduleType int
type pulse int

const (
	FlipFlop moduleType = iota
	Conjunction
)
const (
	None pulse = iota
	Low
	High
)

type module struct {
	name              string
	modType           moduleType
	destinations      []string
	conjunctionMemory map[string]pulse
	flipFlopMemory    bool
}

func newModule(line string) *module {
	var moduleToReturn = module{}
	if line[0] == '%' {
		moduleToReturn.modType = FlipFlop
		moduleToReturn.flipFlopMemory = false
	} else {
		moduleToReturn.modType = Conjunction
		moduleToReturn.conjunctionMemory = make(map[string]pulse)
	}
	moduleToReturn.name = line[1:strings.Index(line, " ->")]

	moduleToReturn.destinations = strings.FieldsFunc(line[strings.Index(line, ">")+1:], func(r rune) bool {
		return !unicode.IsLetter(r)
	})

	return &moduleToReturn
}

func (m *module) isAllInputHigh() bool {
	for _, p := range m.conjunctionMemory {
		if p == Low {
			return false
		}
	}
	return true
}

type queueItem struct {
	origin      string
	destination string
	pulseType   pulse
}
