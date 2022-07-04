package day01

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed testdata.txt
var testData string

func TestData() []int {
	parsedText := strings.Split(testData, "\n")

	numbers := make([]int, len(parsedText))
	for i := 0; i < len(parsedText); i++ {
		n, _ := strconv.Atoi(parsedText[i])

		numbers[i] = n
	}

	return numbers
}