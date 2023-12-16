package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindCalibration(t *testing.T) {
	result1 := findCalibration("1abc2")
	expected1 := 12
	result2 := findCalibration("a1b2c3d4e5f")
	expected2 := 15
	result3 := findCalibration("treb7uchet")
	expected3 := 77

	assert.Equal(t, result1, expected1)
	assert.Equal(t, result2, expected2)
	assert.Equal(t, result3, expected3)
}

func TestFindCalibrationRegex(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"onetwothree", 13},
		{"one two three four five", 15},
		{"zeroone88888twozeronine", 19},
		{"abcone2threexyz", 13},
		{"two1nine", 29},
		{"eightwothree", 83},
		{"abcone2threexyz", 13},
		{"xtwone3four", 24},
		{"4nineeightseven2", 42},
		{"zoneight234", 14},
		{"7pqrstsixteen", 76},
	}
	for _, testCase := range testCases {
		assert.Equal(t, testCase.expected, findCalibrationRegex(testCase.input))
	}
}

func TestSumAllCalibrations(t *testing.T) {
	result := sumAllCalibrations("two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen")
	expected := 281

	assert.Equal(t, expected, result)
}

func TestReverseString(t *testing.T) {
	assert.Equal(t, "654321", reverseString("123456"))
	assert.Equal(t, "a", reverseString("a"))
	assert.Equal(t, "ba", reverseString("ab"))
}
