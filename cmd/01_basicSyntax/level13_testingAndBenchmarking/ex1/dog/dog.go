// Package dog allows us to more fully understand dogs.
package dog

// HumanYearsInDogYears converts human years to dog years.
func HumanYearsInDogYears(n int) int {
	return n * 7
}

// HumanYearsInDogYearsTwo converts human years to dog years.
func HumanYearsInDogYearsTwo(n int) int {
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}
