package main

import (
	"fmt"
	"strings"
)

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
func main() {
	fmt.Println(strings.Map(func(r rune) rune {
		return r + 1
	}, "HAL-9000"))

	f := squares
	fmt.Println(f()())
}
