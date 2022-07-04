package day01

import (
	_ "embed"
	"strconv"
	"strings"
)

//go:embed data.txt
var data string

//go:embed testdata.txt
var testData string

var Data, TestData []int

func init() {
	TestData = parseData(testData)
	Data = parseData(data)
}

func parseData(data string) []int {
	parsedText := strings.Split(data, "\n")

	numbers := make([]int, len(parsedText))
	for i, s := range parsedText {
		n, _ := strconv.Atoi(s)

		numbers[i] = n
	}

	return numbers
}
