package main

import (
	"fmt"
)

func getMostRepeatedString(inputList []string) string {
	counts := make(map[string]int)
	maxCount := 0
	mostRepeatedString := ""

	for _, string := range inputList {
		if count, exists := counts[string]; exists {
			counts[string] = count + 1
		} else {
			counts[string] = 1
		}

		if counts[string] > maxCount {
			maxCount = counts[string]
			mostRepeatedString = string
		}
	}

	return mostRepeatedString
}

func main() {
	inputList := []string{"apple", "pie", "apple", "red", "red", "red"}
	fmt.Println(getMostRepeatedString(inputList))
}