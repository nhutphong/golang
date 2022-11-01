package main

import (
	"fmt"
	"phong/tricks"
)

func main() {
	/*
		make() use cho:
			slice = make([]int, 3, 10)
			dict = make(map[string]string)
			buffer = make(chan string, 5)
			unbuffer = make(chan string)

		new() use cho:
			cac type con lai, khi new thi variable se nhu the nay
				varName = new(int) //ptr
				// *varName = default value of tung type; "" 0 0.0 false

				// neu khong dung new() func thi:
					var varname *int
						varName = nil
						*varName // panic

		all types neu:
			varName = nil va *varName //panic
			*varName = nil // *slice, *dict, *channel = nil, thi se ko lam duoc gi ca
				vi the moi dung make()
			thi cac varable do, ko lam duoc gi ca
	*/
	fmt.Println("")
	var f *float64 // f = nil ; *i error panic
	fmt.Println("\nf = ", f)
	// fmt.Println("*f = ",*f) // panic
	// fmt.Println()
	var one = 5.0
	f = &one
	fmt.Println("f = ", f)
	fmt.Println("*f = ", *f)

	var i = new(int) // i = ptr ; *i = 0
	fmt.Println("i = ", i)
	fmt.Println("*i = ", *i) // *i = 0

	var c = new(*float64) // i = ptr ; *i = nil
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)

	tricks.Format("var info map[string]interface{}")
	var info map[string]interface{}
	// info["name"] = "thong"
	println("info = ", info)
	fmt.Println("info = ", info)

	tricks.Format("new(map)")
	var dict = new(map[string]interface{})

	println("dict = ", dict)
	fmt.Println("dict = ", dict)

	// (*dict)["name"] = "thong"
	// (*dict)["old"] = 35 =
	// println("*dict = ", *dict)
	fmt.Println("*dict = ", *dict) // *dict == nil = 0x0

	tricks.Format("make(map)")
	var data = make(map[string]interface{})

	// println("data = ", data)
	fmt.Println("data = ", data)

	data["name"] = "thong"
	data["old"] = 35
	// println("data = ", data)
	fmt.Println("data = ", data)

	tricks.Format("map empty")
	// var mapEmpty map[string]interface{} // not declared nhu nay
	var mapEmpty = map[string]interface{}{} // gán thẳng map_empty
	// println("mapEmpty = ", mapEmpty)
	fmt.Println("mapEmpty = ", mapEmpty)

	mapEmpty["name"] = "thong"
	mapEmpty["old"] = 35
	println("mapEmpty", mapEmpty)
	fmt.Println("mapEmpty", mapEmpty)

	tricks.Format("slice empty")
	var sliceNil []int // slice nil, phaỉ dùng func append(slice, 2,3,4,5,6); để thêm item vào
	println("sliceNil = ", sliceNil)
	fmt.Println("sliceNil = ", sliceNil)
	// sliceNil[0] = 5 //panic cap=0, nên không thêm được
	// sliceNil[1] = 35 //panic
	sliceNil = append(sliceNil, 3, 4, 5) //lenght=3 cap=3
	println("sliceNil = ", sliceNil)
	fmt.Println("sliceNil = ", sliceNil)

	// khi dung append() add items, neu items vao vuot qua cap-of-slice thi cap=cap*2
	tricks.Format("make([]int, 3, 5)")
	var slicemake = make([]int, 3, 5)
	println("slicemake = ", slicemake)
	fmt.Println("slicemake = ", slicemake)

	slicetwo := append(slicemake, []int{2, 3, 10}...) //new slice
	println("slicemake = ", slicemake)
	println("slicetwo = ", slicetwo)

	slicethree := append(slicetwo, []int{7, 8, 9, 10, 15}...) //new slice
	println("slicethree = ", slicethree)
	fmt.Println("slicethree = ", slicethree)

	slicefour := append(slicethree, []int{7, 8, 9, 10, 15, 44, 55, 33, 22, 99}...) //new slice
	println("slicefour = ", slicefour)
	fmt.Println("slicefour = ", slicefour)

	tricks.Format("channel")
	channew := new(chan int)
	println("channew  = ", channew)       // channew = ptr
	fmt.Println("*channew  = ", *channew) // *channew = nil

	chanmake := make(chan int)
	println("chanmake = ", chanmake)
	fmt.Println("chanmake = ", chanmake)

	tricks.Format("struct")
	var person *struct {
		name string
		old  int
	}
	println("person = ", person)     //nil
	fmt.Println("person = ", person) //nil

	human := new(*struct {
		name string
		old  int
	})
	println("human = ", human)
	fmt.Println("*human = ", *human)

	var dung struct {
		name string
		old  int
	}
	dung = struct {
		name string
		old  int
	}{}
	//println("dung = ", dung)
	fmt.Printf("dung = %#v \n", dung)

	var number int
	fmt.Println("number = ", number)

}
