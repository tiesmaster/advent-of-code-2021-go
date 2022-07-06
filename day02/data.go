package day02

import (
	_ "embed"
	"strings"
)

var (
	Data     = parseData(data)
	TestData = parseData(testData)
)

//go:embed data.txt
var data string

var testData = `forward 5
down 5
forward 8
up 3
down 8
forward 2`

func parseData(data string) []string {
	return strings.Split(data, "\n")
}
