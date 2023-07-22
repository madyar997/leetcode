package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(cellsInRange("A1:F2"))
}

func cellsInRange(s string) []string {
	input := strings.Split(s, ":")

	var ans []string

	for i := input[0][0]; i <= input[1][0]; i++ {
		for j := input[0][1]; j <= input[1][1]; j++ {
			ans = append(ans, fmt.Sprintf("%s%s", string(i), string(j)))
		}
	}

	return ans
}
