package main

import (
	"fmt"
	"math"
	"golang.org/x/exp/constraints"

	"phong/tricks"
)
// chưa nên xài
// dùng generic code sẽ ngắn ngọn dễ đọc hơn khi declare type cho arguments trong func
// các types dùng generic, để get value là chính, chưa được hỗ trợ mạnh các operators: + - * /, < ==. ...
// phải convert về type của argument mà mình dùng đến 

/*
	syntax new: ~
	// tilde forms
	~int
	~[]byte
	~map[int]string
	~chan struct{}
	~struct{x int}

	// unions of terms
	uint8 | uint16 | uint32 | uint64
	~[]byte | ~string
	map[int]int | []int | [16]int | any
	chan struct{} | ~struct{x int}


	comparable: only embed in interface{}, NOT use declared var and argument func

	type _ interface {
		int | ~int // error
	}

	type _ interface {
		interface{int} | interface{~int} // okay
	}

	type _ interface {
		int | interface{~int} // okay
	}

*/


// Integer is made up of all the int types
type Integer interface {
        ~int | ~int8 | ~int16 | ~int32 | ~int64
}

// Float is made up of all the float type
type Float interface {
        ~float32 | ~float64
}

type Number interface {
	Integer | Float
}

type Ordered interface {
	~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 |
	~int32 | ~uint32 | ~int64 | ~uint64 | ~uintptr |
	~float32 | ~float64 | ~string
}


// use package: constraints
func min[T constraints.Ordered](x, y T) T {
    if x > y {
        return x
    } else {
        return y
    }
}



/*

	vi gán type ís any or comparable , không thể thực hiện các opertors như:  + - * / , == <=. && || !=, ...

*/
func Smallest[T Number](slice []T) T {
	firstItem := slice[0]
	for _, v := range slice[1:] {
		if firstItem > v {
			firstItem = v
		}
	}
	return firstItem
}


func Map[key, value any](slice []key , callback func(item key) value) []value {
	new_slice := make([]value, len(slice))
	for index, item := range slice {
		new_slice[index] = callback(item)
	}
	return new_slice
}


// generic func
func genericFunc[T any](slice ...T) {
	for _, item := range slice {
		fmt.Println(item)
	}
}


// SumIntsOrFloats sums the values of map m. It supports both floats and integers
// as map values.
func SumIntsOrFloats[key comparable, value int64 | float64](dict map[key]value) value {
	var total value
	for _, val := range dict {
		total += val
	}
	return total
}


// use generic voi map [key comparable], NOT [key any]
func SumNumbers[key comparable, value Number](dict map[key]value) value {
	var total value
	for _, value := range dict {
		total += value
	}
	return total
}

//
func test[T Number](a,b T) float64 {
	return math.Pow(float64(a), float64(b))
}

func luk[val any](obj struct{name val}) {}
func kit[val any](slice []val) {} 
func wet[val any](callback func() val) {}


func sumNumbers[num Number](slice []num) num {
    var total num
    for _, num := range slice {
        total += num
    }
    return total
}


func main() {
	tricks.Format("generic func")
	genericFunc("Hello", "generic func")

	tricks.Format("Map")
	slice := []int{1, 2, 3}
	resultSlice := Map(slice, func(v int) int { return v * 2 })
	fmt.Println(resultSlice)

	tricks.Format("*")
	slice2 := []string{"a", "b", "c"}
	resultSlice2 := Map(slice2, func(v string) string { return v + v })
	fmt.Println(resultSlice2)

	tricks.Format("func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {}")

	// Initialize a map for the integer values
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}

	fmt.Printf("Generic Sums, type parameters inferred: %v and %v\n",
		SumIntsOrFloats(ints),
		SumIntsOrFloats(floats))

	tricks.Format("func SumNumbers[K comparable, V Number](m map[K]V) V {}")

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
		SumNumbers(ints),
		SumNumbers(floats))

	fmt.Println("test: ", test(2,10))

	luk(struct{name string}{name:"chi thong"})        // okay
	kit([]string{"go", "c"})       // okay
	wet(func() bool {return true}) // okay


	intxxx := []int64{32, 64, 96, 128}    
    floatxxx := []float64{32.0, 64.0, 96.1, 128.2}
    bytes := []int8{8, 16, 24, 32}  

    fmt.Println(sumNumbers(intxxx))
    fmt.Println(sumNumbers(floatxxx))    
    fmt.Println(sumNumbers(bytes))
}