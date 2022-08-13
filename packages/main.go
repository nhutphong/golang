package main

import (
	"fmt"
	"my_project/page"
	pro "my_project/product"
	// giong as pro python
)

/*
	1 file chi thuoc ve 1 package
	file nao thuoc package main	// thi file do run: go run file.go
	
	go run main.go se run init()
	or duoc import duoc cung run init()

	package page
	all variables, struct=class func=method deu thuoc page.<cac vars, func, method_name>
	khong can phai nho ten filename.go, chi dung ten packages_name=folder chua modules.go
*/
func init() {
	fmt.Println("tao la init main.go")
}

func main() {
	fmt.Println("hello main() thuoc main.go")
	page.Sport()
	pro.Milk()
}
