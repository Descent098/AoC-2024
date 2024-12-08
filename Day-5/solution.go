package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
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
	var (
		section1 string
		section2 string
	)

	temp := strings.Split(input, "\r\n")

	var section1Pattern = regexp.MustCompile(`(?m)(\d{2}\|\d{2})`)

	for index, line := range temp {
		if line == "" {
			section1 = strings.Join(temp[0:index-1], "\n")
			section2 = strings.Join(temp[index+1:], "\n")
		}

	}

	fmt.Printf("Section 1: %s\nSection2: %s\n", section1, section2)

	for i, match := range section1Pattern.FindAllString(section1, -1) {
		fmt.Println(match, "found at index", i)
	}
	return result
}

func part2(input string) int {
	var result int

	return result
}
