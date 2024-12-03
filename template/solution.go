package main

import (
	"fmt"
	"os"
)

func main() {
	input := parseInput()
	fmt.Printf("Part 1: %d\n", part1(input))
	fmt.Printf("Part 2: %d\n", part2(input))

}

func parseInput() string {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	return string(input)
}

func part1(input string) int {
	var result int

	return result
}

func part2(input string) int {
	var result int

	return result
}
