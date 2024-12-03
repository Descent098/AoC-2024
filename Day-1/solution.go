package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	dailyInput, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	p1 := part1(string(dailyInput))
	fmt.Println(p1)

	p2 := part2(string(dailyInput))
	fmt.Println(p2)
}

func part1(input string) int {

	leftList, rightList := parseInputText(input)

	difference := 0
	for index := range leftList {
		difference += int(math.Floor(math.Abs(float64((leftList[index] - rightList[index])))))
	}

	return difference
}

func part2(input string) int {
	leftList, rightList := parseInputText(input)

	similarity := 0
	for _, lValue := range leftList {
		appearances := 0
		for _, rValue := range rightList {
			if lValue == rValue {
				appearances += 1
			} else if lValue < rValue {
				break
			}
		}
		similarity += lValue * appearances
	}

	return similarity
}

func parseInputText(input string) ([]int, []int) {
	leftList := make([]int, 0)
	rightList := make([]int, 0)

	for _, line := range strings.Split(input, "\n") {
		temp := strings.Split(line, "   ")
		left, _ := strconv.Atoi(temp[0])

		right, _ := strconv.Atoi(strings.TrimSpace(temp[1]))

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)
	return leftList, rightList
}
