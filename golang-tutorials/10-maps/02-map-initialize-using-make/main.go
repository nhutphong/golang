package main

import "fmt"

func main() {
	// dung make co the add del change key,value in map
	var m = make(map[string]int)

	fmt.Println(m)

	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println("m is not nil")
	}

	// make() function returns an initialized and ready to use map.
	// Since it is initialized, you can add new keys to it.
	m["one hundred"] = 100
	fmt.Println(m)

	//
	dict := map[string]string{
		"name": "thong",
		"old":  "29",
		"city": "gia lai",
	}
	value, ok := dict["name"]
	fmt.Printf("value: %s, ok: %v", value, ok)

	fmt.Println()
	v, k := dict["empty"]
	fmt.Printf("value: %s, ok: %v", v, k)
	fmt.Println()
	fmt.Println(nil)
}
