package page

import (
	"fmt"
	pro "my_project/product"
)

func init() {
	fmt.Println("tao la init sport.go")
}

func Sport() {
	fmt.Println("tao func Sport()")
	pro.Tea("page")
}
