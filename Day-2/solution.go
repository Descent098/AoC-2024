package main

import (
	"fmt"
	"os"
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
	// Safe if
	// The levels(columns in the rows) are either all increasing or all decreasing.
	// Any two adjacent levels(columns in the row) differ by at least one and at most three.
	var safeResults int

	for _, report := range strings.Split(input, "\r\n") {
		temp := strings.Split(report, " ")

		report := make([]int, 0)

		for _, level := range temp {
			currentLevel, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}
			report = append(report, currentLevel)
		}

		if checkReportPart1(report) {
			safeResults++
		}

	}

	return safeResults
}

func part2(input string) int {
	// Safe if
	// The levels(columns in the rows) are either all increasing or all decreasing.
	// Any two adjacent levels(columns in the row) differ by at least one and at most three.
	// The Problem Dampener is a reactor-mounted module that lets the reactor safety systems tolerate a single bad level in what would otherwise be a safe report. It's like the bad level never happened!
	var safeResults int

	for _, report := range strings.Split(input, "\r\n") {
		temp := strings.Split(report, " ")

		report := make([]int, 0)

		for _, level := range temp {
			currentLevel, err := strconv.Atoi(level)
			if err != nil {
				panic(err)
			}
			report = append(report, currentLevel)
		}

		if checkReportPart2(report) {
			safeResults++
		}

	}

	return safeResults
}

func checkReportPart1(report []int) bool {
	var increasing bool

	if report[0] < report[1] {
		increasing = true
	}

	if increasing {
		for i := 0; i < len(report); i++ {
			// Check if at end of list
			if i == len(report)-1 {
				return true
			}
			difference := report[i+1] - report[i]

			if difference <= 0 || difference > 3 {
				return false
			}
		}
	} else {
		for i := 0; i < len(report); i++ {
			// Check if at end of list
			if i == len(report)-1 {
				return true
			}
			difference := report[i] - report[i+1]

			if difference <= 0 || difference > 3 {
				return false
			}
		}
	}

	return true
}

func checkReportPart2(report []int) bool {
	var increasing bool
	var firstError bool

	if report[0] < report[1] {
		increasing = true
	}

	if increasing {
		for i := 0; i < len(report); i++ {
			// Check if at end of list
			if i == len(report)-1 {
				return true
			}
			difference := report[i+1] - report[i]

			if difference <= 0 || difference > 3 {
				if firstError {
					return false
				} else {
					firstError = true
				}
			}
		}
	} else {
		for i := 0; i < len(report); i++ {
			// Check if at end of list
			if i == len(report)-1 {
				return true
			}
			difference := report[i] - report[i+1]

			if difference <= 0 || difference > 3 {
				if firstError {
					return false
				} else {
					firstError = true
				}
			}
		}
	}

	return true
}
