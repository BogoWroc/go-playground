package kyu8

import "math"

func SquareSum(numbers []int) (sum int) {
	for _, e := range numbers {
		sum += int(math.Pow(float64(e), 2))
	}
	return
}
