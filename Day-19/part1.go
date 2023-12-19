package main

import (
	"slices"
	"strconv"
	"strings"
	"unicode"
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

func part1(lines []string) int {
	var extractedWorkflows = extractWorkflows(lines)
	var parts = extractParts(lines)

	var total = 0
	for _, p := range parts {
		if isPartAccepted(p, extractedWorkflows) {
			total += p.x + p.m + p.a + p.s
		}
	}

	return total
}

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

func extractParts(lines []string) []part {
	divIndex := slices.Index(lines, "")
	var parts = make([]part, len(lines)-divIndex-1)
	for i, s := range lines[divIndex+1:] {
		components := strings.FieldsFunc(s, func(r rune) bool {
			return !unicode.IsDigit(r)
		})
		var nums [4]int
		for i, component := range components {
			nums[i], _ = strconv.Atoi(component)
		}
		parts[i] = part{
			x: nums[0],
			m: nums[1],
			a: nums[2],
			s: nums[3],
		}
	}

	return parts
}

func isPartAccepted(p part, workflowsIn workflows) bool {
	var nextWorkflow = "in"
	for nextWorkflow != "A" && nextWorkflow != "R" {
		workflow := workflowsIn[nextWorkflow]
		for _, r := range workflow {
			var varToComp int
			switch r.variable {
			case 'Q':
				nextWorkflow = r.result
				break
			case 'x':
				varToComp = p.x
			case 'm':
				varToComp = p.m
			case 'a':
				varToComp = p.a
			case 's':
				varToComp = p.s
			}
			if r.isGreaterThan {
				if varToComp > r.value {
					nextWorkflow = r.result
					break
				}
			} else {
				if varToComp < r.value {
					nextWorkflow = r.result
					break
				}
			}
		}
	}
	return nextWorkflow == "A"
}
