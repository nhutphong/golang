package main

import (
	"fmt"
	str "strings" // alias
	"math"
	"errors"
)

/*
	// variadic arguments
	func hello(a int, args ...int) {  
		// syntax ok
	}

	func hello(args ...int, b int) {  
		// syntax error
	}

*/

func format() {
    fmt.Println()
    fmt.Println("--------------------------------------------------------------------------------")
    fmt.Println()
}

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

	#golang
	func aciton(args ...string) args ...string {
		return args
	}

	#python
	def action(*args):
		return args

*/

func avg(x float64, y float64) float64 {
	return (x + y) / 2
}

func getStockPriceChange(prevPrice, currentPrice float64) (float64, float64) {
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange
}

func getStockPriceChangeWithError(prevPrice, currentPrice float64) (float64, float64, error) {
	if prevPrice == 0 {
		err := errors.New("Previous price cannot be zero")
		return 0, 0, err
	}
	change := currentPrice - prevPrice
	percentChange := (change / prevPrice) * 100
	return change, percentChange, nil
}

func getNamedStockPriceChange(prevPrice, currentPrice float64) (change, percentChange float64) {
	change = currentPrice - prevPrice
	percentChange = (change / prevPrice) * 100
	return change, percentChange
}


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


	x := 5.75
	y := 6.25

	result := avg(x, y)

	fmt.Printf("Average of %.2f and %.2f = %.2f\n", x, y, result)


	prevStockPrice := 0.0
	currentStockPrice := 100000.0

	change, percentChange := getStockPriceChange(prevStockPrice, currentStockPrice)

	if change < 0 {
		fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
	} else {
		fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
	}


	prevStockPriceOne := 0.0
	currentStockPriceOne := 100000.0

	change, percentChange, err := getStockPriceChangeWithError(prevStockPriceOne, currentStockPriceOne)

	if err != nil {
		fmt.Println("Sorry! There was an error: ", err)
	} else {
		if change < 0 {
			fmt.Printf("The Stock Price decreased by $%.2f which is %.2f%% of the prev price\n", math.Abs(change), math.Abs(percentChange))
		} else {
			fmt.Printf("The Stock Price increased by $%.2f which is %.2f%% of the prev price\n", change, percentChange)
		}
	}
}
