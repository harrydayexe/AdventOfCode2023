package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("An error occurred while reading the input. Exiting...")
		return
	}

	n, _ := strconv.Atoi(os.Args[2])

	lines := strings.Split(string(content), "\n")

	println("Part 1:", part1(lines, n))
	println("Part 2:", part2(lines))
}
