package challenges

import (
	"aoc-2023/pkg/inpututils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func GetAnswerOne() {
	inputLines := inpututils.GetFileInput("1-input.txt")

	// Regex for finding single-digit numbers in a string.
	re := regexp.MustCompile("[0-9]")

	totalValue := 0
	for _, line := range strings.Fields(inputLines) {

		// Extract all numbers from line into an array of strings.
		numbers := re.FindAllString(line, -1)

		// Get first and last number as strings, concat them and convert that into an int.
		lineValue, convErr := strconv.Atoi(numbers[0] + numbers[len(numbers)-1])
		if convErr != nil {
			fmt.Println(convErr)
		}

		// Sum each line value.
		totalValue += lineValue
	}

	fmt.Println("Challenge 1 - total value is ", totalValue)
}
