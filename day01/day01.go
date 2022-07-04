package day01

func CountMeasurementIncreases(measurements []int) int {

	increases := 0
	for i := 1; i < len(measurements); i++ {
		if measurements[i - 1] < measurements[i] {
			increases++
		}
	}

	return increases
}