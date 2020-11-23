package main

import (
	"fmt"
	str "strings" // alias
)

//func functionname(parametername type) returntype {
//function body
//}

func calculateBillA(price int, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func calculateBillB(price, no int) int {
	var totalPrice = price * no
	return totalPrice
}

func rectPropsB(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
	// return area, perimeter
}

// return 2 value
func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}

// Passing anonymous function
// as an argument
func GFG(lambda func(args ...string) string, args ...string) {
	fmt.Println(lambda(args...))
}

/*
	anonymous function
	lambda = func(args ...interface{}) returntype {
				statement........
			}
*/

func main() {
	// _ underscore se duoc ignoze di
	area, _ := rectProps(10.8, 5.6)
	fmt.Printf("Area %f ", area)
	fmt.Println()

	GFG(func(args ...string) string {
		fmt.Println("Tao la lambda")
		fmt.Println("len(args) = ", len(args))
		return str.Join(args, " = ")
	}, "thanh dung", "chi thong", "tan heo", "trau cho")
}
