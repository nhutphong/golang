package main

import "fmt"

func main() {
	// khong the add, del, change key value in map
	// syntax thuong dung la argument cho func
	var m map[string]int
	// key=string
	// value=int
	fmt.Println(m)
	if m == nil {
		fmt.Println("m is nil")
	}

	// The following statement will result in a runtime error
	m["key"] = 100
}
