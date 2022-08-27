package main

import (
	"fmt"
	"phong/tricks"
)


func main() {
	tricks.Format()

	fmt.Println([]byte("thong"))
	fmt.Println([]rune("thong"))

	tricks.Format()

	name:="chi thong"
	for index, char := range name {
		fmt.Printf("index %d char %c\n", index, char)
	}
	
}