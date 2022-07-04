package main

import (
	"fmt"

	"github.com/tiesmaster/advent-of-code-2021-go/day01"
)

func main() {
	fmt.Println("Day 01: count measurements increases of test data: ", day01.CountMeasurementIncreases(day01.TestData()))
	fmt.Println("Day 01: count measurements increases of data: ", day01.CountMeasurementIncreases(day01.Data()))
}
