package main

import (
	"slices"
	"strconv"
	"strings"
)

type part struct {
	x, m, a, s int
}

type rule struct {
	variable      rune
	isGreaterThan bool
	value         int
	result        string
}

type workflows map[string][]rule

func extractWorkflows(lines []string) workflows {
	var workflowsToReturn = make(map[string][]rule)
	divIndex := slices.Index(lines, "")

	for _, workflow := range lines[:divIndex] {
		name := workflow[:strings.Index(workflow, "{")]
		ruleStrings := strings.FieldsFunc(workflow[strings.Index(workflow, "{")+1:len(workflow)-1], func(r rune) bool {
			return r == ','
		})
		rules := make([]rule, len(ruleStrings))
		for i, ruleString := range ruleStrings {
			if strings.Contains(ruleString, ":") {
				val, _ := strconv.Atoi(ruleString[2:strings.Index(ruleString, ":")])
				rules[i] = rule{
					variable:      rune(ruleString[0]),
					isGreaterThan: strings.Contains(ruleString, ">"),
					value:         val,
					result:        ruleString[strings.Index(ruleString, ":")+1:],
				}
			} else {
				rules[i] = rule{
					variable: 'Q',
					result:   ruleString,
				}
			}
		}
		workflowsToReturn[name] = rules
	}

	return workflowsToReturn
}
