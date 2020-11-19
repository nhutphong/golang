package page

import (
	"fmt"
)

/*
	go run main.go se run init()
	or duoc import duoc cung run init()
*/
func init() {
	fmt.Println("tao la init new.go")
}

func New() {
	fmt.Println("tao func New()")
}

func main() {
	fmt.Println("Tao la main() new.go")
}
