package main

import "fmt"

func main() {

	// 1 parent array chua 2 child array chua 3 item=int
	// child array co the co item it hon len=qui dinh, nhung ko dc > hon
	a := [2][3]int{
		{1, 3, 5},
		{7, 9, 20}, // This comma is necessary
	}
	fmt.Println(a)

	b := [3][4]float64{
		{1, 3},
		{4.5, -3, 7.4, 2},
		{6, 2, 11},
	}
	fmt.Println(b)
}
