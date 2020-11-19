package main

import (
    "fmt"
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


func rectPropsB(length, width float64)(area, perimeter float64) {
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


func main() {
    // _ underscore se duoc ignoze di
    area, _ := rectProps(10.8, 5.6)
    fmt.Printf("Area %f ", area)
}
