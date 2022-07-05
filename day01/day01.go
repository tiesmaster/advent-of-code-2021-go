package day01

func CountMeasurementIncreases(measurements []int) int {
	increases := 0
	for i := 1; i < len(measurements); i++ {
		if measurements[i-1] < measurements[i] {
			increases++
		}
	}

	return increases
}

func CountMeasurementIncreasesSlidingWindow(measurements []int) int {
	increases := 0
	previousSum := 0
	for i := 0; i < len(measurements)-3; i++ {
		sum := sum(measurements[i : i+3])
		if previousSum < sum {
			increases++
		}
		previousSum = sum
	}

	return increases
}

func sum(numbers []int) int {
	sum := 0
	for _, x := range numbers {
		sum += x
	}
	return sum
}
