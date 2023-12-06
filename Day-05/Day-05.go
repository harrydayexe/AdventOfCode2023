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

	games := strings.Split(string(content), "\n")

	//println("Part 1:", part1(games))
	println("Part 2:", part2(games))
}
