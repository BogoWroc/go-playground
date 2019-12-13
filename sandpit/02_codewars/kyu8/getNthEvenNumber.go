package kyu8

func NthEven(n int) int {
	var count, even, current int

	for {
		if current%2 == 0 {
			even = current
			count++
		}

		current++

		if count == n {
			break
		}
	}
	return even
}
