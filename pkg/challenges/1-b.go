package challenges

import (
	"aoc-2023/pkg/inpututils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GetAnswerOneB() {
	inputLines := inpututils.GetFileInput("1-input.txt")

	// Regex for finding single-digit numbers in a string.
	re := regexp.MustCompile("[0-9]")

	totalValue := 0
	for _, line := range strings.Fields(inputLines) {

		firstNumber := getFirstNumber(re, line)
		lastNumber := getLastNumber(re, line)

		// Get first and last number as strings, concat them and convert that into an int.
		lineValue, convErr := strconv.Atoi(firstNumber + lastNumber)
		if convErr != nil {
			fmt.Println(convErr)
		}

		output := fmt.Sprintln("line: ", line, "  Number: ", lineValue)
		fmt.Println(output)

		// Sum each line value.
		totalValue += lineValue
	}

	fmt.Println("Challenge 1(b) - total value is ", totalValue)
}

// Get the first number (digit or word).
func getFirstNumber(re *regexp.Regexp, line string) string {
	// Get the first digit in the line.
	firstDigitLocation := re.FindStringIndex(line)
	if firstDigitLocation == nil {
		fmt.Println("Line didn't contain a digit.")
	}

	// Get the subsection of the line before the first digit.
	firstDigitIndex := firstDigitLocation[0]
	subSection := line[:firstDigitIndex]

	// Search the subsection for a spelled out digit.
	numberFromSubSection, found := findNumberInString(subSection, true)

	if found {
		return fmt.Sprint(*numberFromSubSection)
	} else {
		firstDigit := line[firstDigitLocation[0]:firstDigitLocation[1]]
		return firstDigit
	}
}

// Get the last number (digit or word).
func getLastNumber(re *regexp.Regexp, line string) string {
	// Find all occurrences of digits in the line
	allDigitLocations := re.FindAllStringIndex(line, -1)

	// Get the last digit location
	lastDigitLocation := allDigitLocations[len(allDigitLocations)-1]
	lastDigitIndex := lastDigitLocation[1]

	// Get the subsection of the line after the last digit.
	subSection := line[lastDigitIndex:]

	// Search the subsection for a spelled out digit.
	numberFromSubSection, found := findNumberInString(subSection, false)

	if found {
		return fmt.Sprint(*numberFromSubSection)
	} else {
		lastDigit := line[lastDigitLocation[0]:lastDigitLocation[1]]
		return lastDigit
	}
}

func findNumberInString(s string, findFirst bool) (*int, bool) {
	spelledOutNumbers := map[string]int{
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

	bestIndex := -1
	bestNumber := 0
	found := false

	for word, number := range spelledOutNumbers {
		var index int
		if findFirst {
			index = strings.Index(s, word)
		} else {
			index = strings.LastIndex(s, word)
		}
		if index != -1 && (findFirst && (index < bestIndex || bestIndex == -1) || !findFirst && index > bestIndex) {
			bestIndex = index
			bestNumber = number
			found = true
		}
	}

	return &bestNumber, found
}
