package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input := readFile()
	fmt.Println(sumAllCalibrations(input))
}

func sumAllCalibrations(text string) int {
	var calibrationValue int
	lines := splitLinesIntoSlice(text)
	for word := range lines {
		calibrationValue += findCalibrationRegex(lines[word])
	}
	return calibrationValue
}

func readFile() string {
	file, err := os.Open("input.txt")
	if err != nil {
		panic("File input.txt must be present!")
	}
	defer file.Close()

	buf := make([]byte, 1000*100)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if n == 0 {
			break
		}
	}

	return string(buf[:])
}

func splitLinesIntoSlice(buf string) []string {
	return strings.Split(buf, "\n")
}

func findCalibration(word string) int {
	var err error
	var firstIntDetected, lastIntDetected bool
	var firstInt, lastInt, aInt int
	for i := 0; i < len(word); i++ {
		aInt, err = strconv.Atoi(string(word[i]))
		if err != nil {
			continue
		}
		if !firstIntDetected {
			firstInt = aInt
			firstIntDetected = true
		} else {
			lastInt = aInt
			lastIntDetected = true
		}
	}

	if !lastIntDetected {
		lastInt = firstInt
	}

	return firstInt*10 + lastInt
}

func findCalibrationRegex(word string) int {
	dict := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}
	regex_exp := "one|two|three|four|five|six|seven|eight|nine"
	regex := regexp.MustCompile(`\d|` + regex_exp)
	firstDigit := regex.FindString(word)

	regex = regexp.MustCompile(`\d|` + reverseString(regex_exp))
	lastDigit := regex.FindString(reverseString(word))

	return dict[firstDigit]*10 + dict[reverseString(lastDigit)]
}

func reverseString(s string) string {
	if len(s) < 2 {
		return s
	}
	return reverseString(s[1:]) + string(s[0])
}
