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
	for i := 0; i < len(measurements) - 3; i++ {
		sum := measurements[i] + measurements[i+1] + measurements[i+2]
		if previousSum < sum {
			increases++
		}
		previousSum = sum
	}

	return increases
}
