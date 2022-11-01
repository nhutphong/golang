package main

import (
	"fmt"
	// "time"
)

func main() {
	/*
		khi so sanh(==) cac structs neu cac fields co value=nhau thi se la true, else false
	*/
	h := Human{name: "thong", old: 38}
	humancopy := h.NewHuman()
	humanpointer := h.Pointer()

	fmt.Println("h = ", h)
	fmt.Println("humanpointer = ", humanpointer)
	fmt.Println("humancopy = ", humancopy)
	println()

	humanpointer.name = "nguyen thong dung"
	humancopy.name = "copy-update"

	fmt.Println("h = ", h)
	fmt.Println("humanpointer = ", humanpointer)
	fmt.Println("humancopy = ", humancopy)
	println()

	fmt.Println("h == humancopy is", h == humancopy)         //false
	fmt.Println("&h == &humancopy is ", &h == &humancopy)    //false
	fmt.Println("&h == humanpointer is", &h == humanpointer) //true
	fmt.Println("h == *humanpointer is", h == *humanpointer) //true
	println()

	println("&humancopy = ", &humancopy)
	println("&h = ", &h)
	println("humanpointer = ", humanpointer)
}

type Human struct {
	name string
	old  int
}

func (h Human) NewHuman() Human {
	return h //new Human == copy
}
func (h *Human) Pointer() *Human {
	return h //pointer
}
