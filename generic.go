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


// thiết lập Constraint = hạn chế
type Number interface {
 	int | int8 | int16 | int32 | int64 | float64
}
// nếu func() dùng generic thì các types của arguments sẽ thuộc 1 trong các types chúng ta định nghĩa 
// trong interface


/*
	[T Number] = thiết lặp type dùng chung cho arguments trong func
	T in (int | int8 | int16 | int32 | int64 | float64)

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


// generic cho struct
type Vector[T any] []T

type LinkedList[T any] struct {
	next *LinkedList[T]
	val  T
}

type Pair[T1, T2 any] struct {
	v1 T1
	v2 T2
}

type Tuple[T1, T2, T3 any] struct {
	v1 T1
	v2 T2
	v3 T3
}


func Map[K, V any](slice []K, transform func(K) V) []V {
	new_slice := make([]V, len(slice))
	for index, item := range slice {
		new_slice[index] = transform(item)
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
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}


// SumNumbers sums the values of map m. Its supports both integers
// and floats as map values.
func SumNumbers[K comparable, V Number](dict map[K]V) V {
	var total V
	for _, value := range dict {
		total += value
	}
	return total
}

type NumberOne interface {
    constraints.Integer | constraints.Float
}

//
func test[T NumberOne](a,b T) float64 {
	return math.Pow(float64(a), float64(b))
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

}