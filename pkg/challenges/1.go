package challenges

import (
	"aoc-2023/pkg/inpututils"
	"fmt"
	"strings"
)

func GetAnswerOne() {
	inputLines := inpututils.GetFileInput("1-input.txt")

	for _, line := range strings.Fields(inputLines) {
		fmt.Println(line)
	}
}
