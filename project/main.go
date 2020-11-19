package main

import (
	"fmt"
	"rootname/page"
	pro "rootname/product"
	// giong as pro python
)

/*
	go run main.go se run init()
	or duoc import duoc cung run init()
*/
func init() {
	fmt.Println("tao la init main.go")
}

func main() {
	fmt.Println("hello main() thuoc main.go")
	page.Sport()
	pro.Milk()
}
