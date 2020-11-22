package main

import (
	"fmt"
	"rootpath/model"
)

func main() {
	c := model.Customer{
		Id:   1,
		Name: "Rajeev Singh",
	}

	c.married = true // Error: can not refer to unexported field or method

	//a := model.address{} // Error: can not refer to unexported name

	fmt.Println("Programmer = ", c)
}
