package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("An error occurred while reading the input. Exiting...")
		return
	}

	lines := strings.Split(string(content), "\n")

	part1Ans := part1(lines)
	part2Ans := part2(lines)

	fmt.Println("Part 1 Answer:", part1Ans)
	fmt.Println("Part 2 Answer:", part2Ans)
}
