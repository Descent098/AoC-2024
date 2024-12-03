package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
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

	var re = regexp.MustCompile(`(?m)mul\((\d{1,3})\,(\d{1,3})\)`)

	for _, match := range re.FindAllString(input, -1) {
		temp := strings.Split(match[4:len(match)-1], ",")
		leftOperator, err := strconv.Atoi(temp[0])
		if err != nil {
			panic(err)
		}
		rightOperator, err := strconv.Atoi(temp[1])
		if err != nil {
			panic(err)
		}
		result = result + (leftOperator * rightOperator)
	}

	return result
}

func part2(input string) int {
	var result int
	do := true

	var re = regexp.MustCompile(`(?m)mul\((\d{1,3})\,(\d{1,3})\)|do\(\)|don\'t\(\)`)

	for _, match := range re.FindAllString(input, -1) {
		if match == "do()" {
			do = true
		} else if match == "don't()" {
			do = false
		}

		if do && strings.HasPrefix(match, "mul(") {
			temp := strings.Split(match[4:len(match)-1], ",")
			leftOperator, err := strconv.Atoi(temp[0])
			if err != nil {
				panic(err)
			}
			rightOperator, err := strconv.Atoi(temp[1])
			if err != nil {
				panic(err)
			}
			result = result + (leftOperator * rightOperator)
		}

	}

	return result
}
