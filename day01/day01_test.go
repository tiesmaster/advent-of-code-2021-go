package day01

import "testing"

func TestCountMeasurementIncreases(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{TestData, 7},
		{Data, 1553},
	}
	for _, c := range cases {
		got := CountMeasurementIncreases(c.in)
		if got != c.want {
			t.Errorf("CountMeasurementIncreases(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestCountMeasurementIncreasesSlidingWindow(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{TestData, 5},
		{Data, 1597},
	}
	for _, c := range cases {
		got := CountMeasurementIncreasesSlidingWindow(c.in)
		if got != c.want {
			t.Errorf("CountMeasurementIncreasesSlidingWindow(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}
