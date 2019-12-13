package kyu8

func monkeyCount(n int) []int {
	var monkeys []int
	for i := 1; i <= n; i++ {
		monkeys = append(monkeys, i)
	}
	return monkeys
}
