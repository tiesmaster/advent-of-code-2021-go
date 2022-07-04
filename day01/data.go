package day01

import (
	_ "embed"
	"strconv"
	"strings"
)

var (
	Data = parseData(data)
	TestData = parseData(testData)
)

//go:embed data.txt
var data string

var testData = `199
200
208
210
200
207
240
269
260
263`

func parseData(data string) []int {
	parsedText := strings.Split(data, "\n")

	numbers := make([]int, len(parsedText))
	for i, s := range parsedText {
		n, _ := strconv.Atoi(s)

		numbers[i] = n
	}

	return numbers
}
