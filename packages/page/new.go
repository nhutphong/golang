package page

import (
	"fmt"
)

/*
	
	go run main.go se run init()
	or duoc import duoc cung run init()

	package page
	all variables, struct=class func=method deu thuoc page.<cac vars, func, method_name>
	khong can phai nho ten filename.go, chi dung ten packages_name=folder chua modules.go
*/
func init() {
	fmt.Println("tao la init new.go")
}

func New() {
	fmt.Println("tao func New()")
}

func page() {
	fmt.Println("Tao la main() new.go")
}
