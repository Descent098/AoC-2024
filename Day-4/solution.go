package main

import (
	"fmt"
	"os"
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

	lines := strings.Split(input, "\r\n")

	rows := len(lines)
	var count int
	for rowIndex, line := range lines {
		for columnIndex, character := range line {
			columns := len(line)

			if string(character) == "X" {
				if columns > columnIndex+3 {
					// Horizontal right way (i.e X M A S)
					l := strings.Split(line, "")
					if strings.Join(l[columnIndex:columnIndex+4], "") == "XMAS" {
						fmt.Printf("\nFOUND XMAS 1")
						count++
					}
				}

				if rows > rowIndex+3 {
					// Vertical right way i.e. X
					//						   M
					//						   A
					//						   S
					temp := fmt.Sprintf("%s%s%s%s", string(lines[rowIndex][columnIndex]), string(lines[rowIndex+1][columnIndex]), string(lines[rowIndex+2][columnIndex]), string(lines[rowIndex+3][columnIndex]))
					if temp == "XMAS" {
						fmt.Printf("\nFOUND XMAS 2")
						count++
					}
				}

				// Diagonal (Top left start) i.e. X
				//					    		   M
				//								    A
				//							         S
				if (columns > (columnIndex + 3)) && (rows > (rowIndex + 3)) {
					temp := fmt.Sprintf("%s%s%s%s", string(lines[rowIndex][columnIndex]), string(lines[rowIndex+1][columnIndex+1]), string(lines[rowIndex+2][columnIndex+2]), string(lines[rowIndex+3][columnIndex+3]))
					if temp == "XMAS" {
						fmt.Printf("\nFOUND XMAS 3")
						count++
					}
				}

				// Diagonal (Top Right start) i.e. X
				//								 M
				//								A
				//							   S
				if (columns < (columnIndex - 3)) && (rows < (rowIndex - 3)) {
					temp := fmt.Sprintf("%s%s%s%s", string(lines[rowIndex][columnIndex]), string(lines[rowIndex-1][columnIndex-1]), string(lines[rowIndex-2][columnIndex-2]), string(lines[rowIndex-3][columnIndex-3]))
					if temp == "XMAS" {
						fmt.Printf("\nFOUND XMAS 4")
						count++
					}
				}

			} else if string(character) == "S" {
				if columns > columnIndex+3 {
					// Horizontal backwards (i.e. S A M X)
					l := strings.Split(line, "")
					if strings.Join(l[columnIndex:columnIndex+4], "") == "SAMX" {
						fmt.Printf("\nFOUND SAMX 5")
						count++
					}
				}

				if rows > rowIndex+3 {
					// Vertical Backwards way i.e. S
					//						   	   A
					//						   	   M
					//						   	   X
					temp := fmt.Sprintf("%s%s%s%s", string(lines[rowIndex][columnIndex]), string(lines[rowIndex+1][columnIndex]), string(lines[rowIndex+2][columnIndex]), string(lines[rowIndex+3][columnIndex]))
					if temp == "SAMX" {
						fmt.Printf("\nFOUND SAMX 6")
						count++
					}
				}

				// Diagonal (Top left start) i.e. S
				//					    		   A
				//								    M
				//							         X
				if (columns > (columnIndex + 3)) && (rows > (rowIndex + 3)) {
					temp := fmt.Sprintf("%s%s%s%s", string(lines[rowIndex][columnIndex]), string(lines[rowIndex+1][columnIndex+1]), string(lines[rowIndex+2][columnIndex+2]), string(lines[rowIndex+3][columnIndex+3]))
					if temp == "SAMX" {
						fmt.Printf("\nFOUND SAMX 7")
						count++
					}
				}

				// Diagonal (Top Right start) i.e. S
				//								 A
				//								M
				//							   X
				if (columns < (columnIndex - 3)) && (rows < (rowIndex - 3)) {
					temp := fmt.Sprintf("%s%s%s%s", string(lines[rowIndex][columnIndex]), string(lines[rowIndex-1][columnIndex-1]), string(lines[rowIndex-2][columnIndex-2]), string(lines[rowIndex-3][columnIndex-3]))
					if temp == "SAMX" {
						fmt.Printf("\nFOUND SAMX 8")
						count++
					}
				}

			}
		}
	}

	// Search Diagonal

	return count
}

func part2(input string) int {
	var result int

	return result
}
