package main

import (
	"fmt"
	"phong/tricks"
)


func main() {
	fmt.Println("")
	var f *float64
	fmt.Println("\nf",f)
	// fmt.Println(*f)


	var i = new(int)
	fmt.Println(i)
	fmt.Println(*i == 0)

	var c = new(*float64)
	fmt.Println("\nc", c == nil)
	fmt.Println("\n*c",*c)

	tricks.Format("map")
	var info map[string]interface{}
	println("info", info)
	fmt.Println("info", info)


	tricks.Format("new(map)")
	var dict = new(map[string]interface{})

	println("dict", dict)
	fmt.Println("dict", dict)

	// (*dict)["name"] = "thong"
	// (*dict)["old"] = 35
	println("*dict", *dict)
	fmt.Println("*dict", *dict) // *dict == nil = 0x0


	tricks.Format("make(map)")
	var data = make(map[string]interface{})

	println("data", data)
	fmt.Println("data", data)

	data["name"] = "thong"
	data["old"] = 35
	println("data", data)
	fmt.Println("data", data)


	tricks.Format("map empty")
	// var mapEmpty map[string]interface{} // not declared nhu nay
	var mapEmpty = map[string]interface{}{} // gán thẳng map_empty
	println("mapEmpty", mapEmpty)
	fmt.Println("mapEmpty", mapEmpty)

	mapEmpty["name"] = "thong"
	mapEmpty["old"] = 35
	println("mapEmpty", mapEmpty)
	fmt.Println("mapEmpty", mapEmpty)


	tricks.Format("slice empty")
	var sliceNil []int // slice nil, phaỉ dùng func append(slice, 2,3,4,5,6); để thêm item vào 
	println("sliceNil", sliceNil)
	fmt.Println("sliceNil", sliceNil)
	// sliceNil[0] = 5 //panic cap=0, nên không thêm được 
	// sliceNil[1] = 35 //panic 
	sliceNil = append(sliceNil,3,4,5)
	println("sliceNil", sliceNil)
	fmt.Println("sliceNil", sliceNil)

}