package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("An error occurred while reading the input. Exiting...")
		return
	}

	lines := strings.Split(string(content), "\n")
	count := 0

	isNumeric := func(r rune) bool {
		return unicode.IsDigit(r)
	}

	for _, line := range lines {
		digit1Index := strings.IndexFunc(line, isNumeric)
		if digit1Index != -1 {
			count += int(line[digit1Index]-'0') * 10
		}

		digit2Index := strings.LastIndexFunc(line, isNumeric)
		if digit2Index != -1 {
			count += int(line[digit2Index] - '0')
		}
	}

	fmt.Println("Total:", count)
}
