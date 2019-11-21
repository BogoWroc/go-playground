package capstone

import "strings"

type Sentence string

func (sentence Sentence) CountWords() map[string]int {
	countedWords := make(map[string]int, 100)
	fields := strings.Fields(string(sentence))

	for _, value := range fields {
		countedWords[strings.ToLower(strings.Trim(value, ".,;"))] += 1
	}

	return countedWords
}
