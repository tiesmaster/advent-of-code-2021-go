package day01

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed testdata.txt
var testData string

func TestData() []int {
	return parseData(testData)
}

func parseData(data string) []int {
	parsedText := strings.Split(data, "\n")

	numbers := make([]int, len(parsedText))
	for i := 0; i < len(parsedText); i++ {
		n, _ := strconv.Atoi(parsedText[i])

		numbers[i] = n
	}

	return numbers
}