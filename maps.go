package main

import (
	"fmt"
)

func main() {
	// var demoMap map[string]int //cach1
	dict := make(map[string]int) //cach2

	dict["old"] = 27
	value, ok := dict["old"]
	fmt.Println("value: ", value, "ok: ", ok)

	delete(dict, "old")
	fmt.Println(dict["old"])

	// Tiếp theo ví dụ trên
	v, k := dict["old"]
	fmt.Println("Giá trị của phần tử là: ", v)
	fmt.Println("Kiểm tra phần tử có tồn tại hay không: ", k)
}
