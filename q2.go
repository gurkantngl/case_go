package main

import (
	"fmt"
)

func recursive(n int) {
	if n == 1 {
		return
	}
	recursive(n / 2)
	fmt.Println(n)
}

func main() {
	recursive(9)
}