package kyu8

func ReversedString(word string) string {
	var result string
	for _, v := range word {
		result = string(v) + result
	}
	return result
}
