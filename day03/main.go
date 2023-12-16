package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Grid struct {
	schematic [][]string
}

func loadInput(filePath string) []string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("File reading error", err)
		return []string{}
	}

	return strings.Split(string(data), "\n")
}

func loadGrid(lines []string) Grid {
	var grid [][]string
	var row []string
	for _, line := range lines {
		for _, char := range line {
			row = append(row, string(char))
		}
		grid = append(grid, row)
		row = []string{}
	}
	return Grid{schematic: grid}
}

func returnPartNumbers(lines []string) []string {
	validParts := []string{}

	regex := `\d+`
	expression := regexp.MustCompile(regex)

	for rowIndex, line := range lines {
		matches := expression.FindAllStringSubmatchIndex(line, -1)
		for _, match := range matches {
			// [column row]
			matchColumn := match[0]
			matchRow := match[1]
			checkLeftSymbolIndex := matchColumn - 1
			checkRightSymbolIndex := matchRow + 1

			if checkLeftSymbolIndex < 0 {
				checkLeftSymbolIndex = 0
			}
			if checkRightSymbolIndex > len(line)-1 {
				checkRightSymbolIndex = len(line) - 1
			}
			// check if there is any symbol on the left or right of the number
			if anyStringSymbol(string(lines[rowIndex][checkLeftSymbolIndex])) ||
				anyStringSymbol(string(lines[rowIndex][checkRightSymbolIndex-1])) {
				validParts = append(validParts, lines[rowIndex][matchColumn:matchRow])
				continue
			}
			// check if there is any symbol on the row above the number
			if rowIndex > 0 {
				if anyStringSymbol(lines[rowIndex-1][checkLeftSymbolIndex:checkRightSymbolIndex]) {
					validParts = append(validParts, lines[rowIndex][matchColumn:matchRow])
					continue
				}
			}
			// check if there is any symbol on the row bellow the number
			if rowIndex < len(lines)-1 {
				if anyStringSymbol(lines[rowIndex+1][checkLeftSymbolIndex:checkRightSymbolIndex]) {
					validParts = append(validParts, lines[rowIndex][matchColumn:matchRow])
				}
			}
		}
	}
	return validParts
}

// [0,1] [1,2]
func rangeIntersects(range1 []int, range2 []int) bool {
	if range1[0] <= range2[0] && range1[1] >= range2[0] ||
		range1[0] >= range2[0] && range1[0] <= range2[0]+1 {
		return true
	}
	return false
}

func returnGearRatios(lines []string) [][]string {
	gearRatios := [][]string{}
	var currentGearRatio []string
	var moreThanTwoNeighbours bool

	for rowIndex, line := range lines {
		gearMatches := regexp.MustCompile(`\*`).FindAllStringSubmatchIndex(line, -1)
		for _, gearMatch := range gearMatches {
			//digits := regexp.MustCompile(`(\d+)\*(\d+)`).FindAllStringSubmatch(line, -1)

			//if len(digits) > 0 {
			//currentGearRatio = []string{digit[0][1], digit[0][2]}
			//}

			digitsOnCurrentRow := regexp.MustCompile(`\d+`).FindAllStringSubmatchIndex(lines[rowIndex], -1)
			for _, digitMatchSameLine := range digitsOnCurrentRow {
				if rangeIntersects(digitMatchSameLine, gearMatch) {
					if len(currentGearRatio) >= 2 {
						moreThanTwoNeighbours = true
						continue
					}
					currentDigit := lines[rowIndex][digitMatchSameLine[0]:digitMatchSameLine[1]]
					currentGearRatio = append(currentGearRatio, currentDigit)
				}
			}

			if rowIndex > 0 {
				digitsAbove := regexp.MustCompile(`\d+`).FindAllStringSubmatchIndex(lines[rowIndex-1], -1)
				for _, digitMatchAbove := range digitsAbove {
					if rangeIntersects(digitMatchAbove, gearMatch) {
						if len(currentGearRatio) >= 2 {
							moreThanTwoNeighbours = true
							continue
						}
						currentDigit := lines[rowIndex-1][digitMatchAbove[0]:digitMatchAbove[1]]
						currentGearRatio = append(currentGearRatio, currentDigit)
					}
				}
			}

			if rowIndex < len(lines)-1 {
				digitsBellow := regexp.MustCompile(`\d+`).FindAllStringSubmatchIndex(lines[rowIndex+1], -1)
				for _, digitMatchBellow := range digitsBellow {
					if rangeIntersects(digitMatchBellow, gearMatch) {
						if len(currentGearRatio) >= 2 {
							moreThanTwoNeighbours = true
							continue
						}
						currentDigit := lines[rowIndex+1][digitMatchBellow[0]:digitMatchBellow[1]]
						currentGearRatio = append(currentGearRatio, currentDigit)
					}
				}
			}

			if len(currentGearRatio) == 2 && !moreThanTwoNeighbours {
				gearRatios = append(gearRatios, currentGearRatio)
			}
			currentGearRatio = []string{}
			moreThanTwoNeighbours = false
		}
	}
	return gearRatios
}

func calculateGearRatios(gearRatios [][]string) int {
	var sum int
	for _, gearRatio := range gearRatios {
		ratio1, err := strconv.Atoi(gearRatio[0])
		if err != nil {
			panic("kaboom")
		}
		ratio2, err := strconv.Atoi(gearRatio[1])
		if err != nil {
			panic("kaboom")
		}
		sum += ratio1 * ratio2
	}
	return sum
}

func anyStringSymbol(word string) bool {
	return strings.ContainsAny(word, "!$#%$&/()=?'+*~^\\|{}[]<>/-,;_:@€€£‰¶÷[]≠±≤≥∞≈©®™")
}

func sumAllPartsNumbers(parts []string) int {
	sum := 0
	for _, part := range parts {
		number, err := strconv.Atoi(part)
		if err != nil {
			panic("kaboom")
		}
		sum += number
	}
	return sum
}

func main() {
	lines := loadInput("input.txt")
	part1 := sumAllPartsNumbers(returnPartNumbers(lines))
	part2 := calculateGearRatios(returnGearRatios(lines))

	fmt.Println("Part 1: ", part1)
	fmt.Println("Part 2: ", part2)
}
