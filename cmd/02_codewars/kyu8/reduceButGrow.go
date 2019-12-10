package kyu8

func Grow(arr []int) int {
	result := 1
	for _, e := range arr {
		result *= e
	}
	return result
}
