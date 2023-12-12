package main

import "testing"

//func Test_On_Input(t *testing.T) {
//	content, err := os.ReadFile("input.txt")
//	if err != nil {
//		t.Error("An error occurred while reading the input. Exiting...")
//		return
//	}
//
//	games := strings.Split(string(content), "\n")
//
//	if part1(games) != 6852 {
//		t.Error("Incorrect!")
//	}
//}

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
