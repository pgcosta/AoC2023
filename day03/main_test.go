package main

import (
	"regexp"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestLoadGrid(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"......755.",
		"...$.*....",
		".664.598..",
	}
	grid := loadGrid(lines)

	assert.Equal(t, len(grid.schematic), 10)
	assert.Equal(t, len(grid.schematic[9]), 10)
	assert.Equal(t, returnPartNumbers(lines), []string{"467", "35", "633", "617", "592", "755", "664", "598"})
}

func TestAnyStringSymbol(t *testing.T) {
	assert.Equal(t, anyStringSymbol("asdfd..324"), false)
	assert.Equal(t, anyStringSymbol("1"), false)
	assert.Equal(t, anyStringSymbol("."), false)
	assert.Equal(t, anyStringSymbol("saf%gasd"), true)
	assert.Equal(t, anyStringSymbol("as.df&sdf"), true)
	assert.Equal(t, anyStringSymbol("..*."), true)
}

func TestCustomRegex(t *testing.T) {
	input := "...*.*.\n.**...."
	matches := regexp.MustCompile(`\*`).FindAllStringSubmatchIndex(input, -1)

	assert.Equal(t, matches, [][]int{{3, 4}, {5, 6}, {9, 10}, {10, 11}})
}

func TestRangeIntersects(t *testing.T) {
	range1 := []int{6, 9} //633
	range2 := []int{5, 6} //*

	assert.Equal(t, rangeIntersects(range1, range2), true)
}

func TestReturnGearRatios(t *testing.T) {
	lines := []string{
		"467..114..",
		"...*......",
		"..35..633.",
		"......#...",
		"617*......",
		".....+.58.",
		"..592.....",
		"45*3......",
		"......755.",
		"...$.*....",
		".664.598..",
	}

	result := returnGearRatios(lines)
	expected := [][]string{{"467", "35"}, {"45", "3"}, {"755", "598"}}

	assert.Equal(t, result, expected)
}

func TestCalculateGearRatios(t *testing.T) {
	lines := []string{
		"............",
		"2*2......2*2",
		"..$.....*...",
		"..78...2.4..",
		"..*....60...",
		"78.........9",
		".5.....23..$",
		"8...90*12...",
		"............",
		"2.2......12.",
		".*.........*",
		"1.1..503+.56",
		"......120...",
		"........*...",
		".........410",
	}

	result := calculateGearRatios(returnGearRatios(lines))
	assert.Equal(t, 78*78+12*56+120*410+4+4, result)
}
