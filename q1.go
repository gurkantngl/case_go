package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input_strings := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklklkl", "l"}

	sort.Slice(input_strings, func(i, j int) bool {
		countI := strings.Count(input_strings[i], "a")
		countJ := strings.Count(input_strings[j], "a")
		if countI != countJ {
			return countI > countJ
		}
		return len(input_strings[i]) > len(input_strings[j])
	})

	fmt.Println(input_strings)
}