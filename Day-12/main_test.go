package main

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func Test_On_Input(t *testing.T) {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		t.Error("An error occurred while reading the input. Exiting...")
		return
	}

	games := strings.Split(string(content), "\n")

	if part1(games) != 6852 {
		t.Error("Incorrect!")
	}
}

func TestPrunePasses1(t *testing.T) {
	var row = HotSpringRow{
		HotSpringRecords: []HotSpring{Damaged, Working, Unknown, Working, Damaged, Damaged, Damaged},
		HotSpringGroups:  []int{1, 1, 3},
	}

	if row.shouldBePruned() {
		t.Error("Incorrect")
	}
}

func TestPrunePasses2(t *testing.T) {
	var row = HotSpringRow{
		HotSpringRecords: []HotSpring{Damaged, Unknown, Unknown, Working, Damaged, Damaged, Damaged},
		HotSpringGroups:  []int{1, 1, 3},
	}

	if row.shouldBePruned() {
		t.Error("Incorrect")
	}
}

func TestPrunePasses3(t *testing.T) {
	var row = HotSpringRow{
		HotSpringRecords: []HotSpring{Working, Unknown, Unknown, Working, Damaged, Damaged, Damaged},
		HotSpringGroups:  []int{1, 1, 3},
	}

	if !row.shouldBePruned() {
		t.Error("Incorrect")
	}
}

func TestPrunePasses4(t *testing.T) {
	var row = HotSpringRow{
		HotSpringRecords: []HotSpring{
			Unknown, Unknown, Unknown, Working, Damaged, Damaged, Damaged, Unknown,
			Unknown, Unknown, Unknown, Working, Damaged, Damaged, Damaged, Unknown,
			Unknown, Unknown, Unknown, Working, Damaged, Damaged, Damaged, Unknown,
			Unknown, Unknown, Unknown, Working, Damaged, Damaged, Damaged, Unknown,
			Unknown, Unknown, Unknown, Working, Damaged, Damaged, Damaged,
		},
		HotSpringGroups: []int{
			1, 1, 3,
			1, 1, 3,
			1, 1, 3,
			1, 1, 3,
			1, 1, 3,
		},
	}

	if row.shouldBePruned() {
		t.Error("Incorrect")
	}
}

func TestUnfold(t *testing.T) {
	var row = HotSpringRow{
		HotSpringRecords: []HotSpring{Unknown, Unknown, Unknown, Working, Damaged, Damaged, Damaged},
		HotSpringGroups:  []int{1, 1, 3},
	}

	row.unfoldRow()

	var expectedRecords = []HotSpring{
		Unknown, Unknown, Unknown, Working, Damaged, Damaged, Damaged, Unknown,
		Unknown, Unknown, Unknown, Working, Damaged, Damaged, Damaged, Unknown,
		Unknown, Unknown, Unknown, Working, Damaged, Damaged, Damaged, Unknown,
		Unknown, Unknown, Unknown, Working, Damaged, Damaged, Damaged, Unknown,
		Unknown, Unknown, Unknown, Working, Damaged, Damaged, Damaged,
	}
	var expectedGroups = []int{
		1, 1, 3,
		1, 1, 3,
		1, 1, 3,
		1, 1, 3,
		1, 1, 3,
	}

	fmt.Println(row.HotSpringRecords)

	for i, group := range row.HotSpringGroups {
		if group != expectedGroups[i] {
			fmt.Println(row)
			t.Error("Group did not match")
		}
	}
	for i, record := range row.HotSpringRecords {
		if record != expectedRecords[i] {
			t.Error("Record did not match")
		}
	}
}

func TestCountCombinations(t *testing.T) {
	var row = HotSpringRow{
		HotSpringRecords: []HotSpring{Unknown, Unknown, Unknown, Working, Damaged, Damaged, Damaged},
		HotSpringGroups:  []int{1, 1, 3},
	}

	row.unfoldRow()

	var res = countPossibleCombinations(row)

	if res != 1 {
		t.Error("Wrong answer")
	}
}
