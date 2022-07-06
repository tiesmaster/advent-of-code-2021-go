package day02

import "strconv"

func CalculateLocation(instructions []string) int {
	x, z := 0, 0
	for _, line := range instructions {
		switch line[:1] {
		case "f":
			x += parse(line[8:])
		case "u":
			z -= parse(line[3:])
		case "d":
			z += parse(line[5:])
		}
	}

	return x * z
}

func parse(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}