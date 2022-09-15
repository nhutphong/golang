package main

import (
	"fmt"
)

var list [2]int
var slice = []int{}
var num *int
var person struct{name string; old int}

func main() {
	fmt.Println()
	person  = struct{name string; old int}{name: "chithong", old:50}
	fmt.Printf("list %#v %v",person, slice ==nil)
}