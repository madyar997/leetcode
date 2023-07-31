package main

import (
	"strings"
)

func main() {
	toGoatLatin("I speak Goat Latin")
}

func toGoatLatin(sentence string) string {
	words := strings.Split(sentence, " ")
	result := make([]string, 0, len(words))

	for _, word := range words {
		if strings.ContainsAny(strings.ToLower(string(word[0])), "aeiou") {
			word = word + "ma"
			result = append(result, word)
		} else {
			firstLetter := word[0]
			word = word[1:] + string(firstLetter) + "ma"
			result = append(result, word)
		}
	}

	for i := 1; i <= len(result); i++ {
		for j := 0; j < i; j++ {
			result[i-1] += "a"
		}

	}

	return strings.Join(result, " ")
}
