package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(sortSentence("is2 sentence4 This1 a3"))
}

func sortSentence(s string) string {
	words := strings.Split(s, " ")

	var (
		ind int
		err error
	)

	wordArr := make([]string, len(words), len(words))

	for _, word := range words {
		ind, err = strconv.Atoi(string(word[len(word)-1]))
		if err != nil {
			fmt.Println(fmt.Errorf("cannot convert to int %s", err))
		}

		wordArr[ind-1] = word[:(len(word) - 1)]
	}

	return strings.Join(wordArr, " ")
}
