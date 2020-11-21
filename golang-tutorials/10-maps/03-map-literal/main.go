package main

import "fmt"

func main() {
	// co keys va values
	var m = map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5, // Comma is necessary
	}

	fmt.Println(m)
}
