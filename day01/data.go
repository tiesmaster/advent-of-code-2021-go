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

//go:embed data.txt
var data string

func Data() []int {
	return parseData(data)
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