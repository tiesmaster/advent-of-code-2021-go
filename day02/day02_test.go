package day02

import "testing"

func TestCalculateLocation(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{TestData, 150},
		{Data, 1561344},
	}
	for _, c := range cases {
		got := CalculateLocation(c.in)
		if got != c.want {
			t.Errorf("CalculateLocation(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

func TestCalculateLocationWithAim(t *testing.T) {
	cases := []struct {
		in   []string
		want int
	}{
		{TestData, 900},
		{Data, 1848454425},
	}
	for _, c := range cases {
		got := CalculateLocationWithAim(c.in)
		if got != c.want {
			t.Errorf("CalculateLocationWithAim(%v) == %v, want %v", c.in, got, c.want)
		}
	}
}

