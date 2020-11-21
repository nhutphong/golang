package main

import "fmt"

func main() {
	var personAge = map[string]int{
		"Rajeev": 25,
		"James":  32,
		"Sarah":  29,
	}

	for key, value := range personAge {
		fmt.Println(key, value)
	}

}
