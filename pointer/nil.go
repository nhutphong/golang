package main

import (
	"fmt"
	"unsafe"
	"phong/tricks"
)
func main() {
	tricks.Format("nil thuoc types: interface{} pointer chan slice map func()")
	// compiler to deduce the type of a nil value.
	_ = (*struct{})(nil)  // pointer
	_ = []int(nil)
	_ = map[int]bool(nil)
	_ = chan string(nil)
	_ = (func())(nil)
	_ = interface{}(nil)

	// These lines are equivalent to the above lines.
	var _ *struct{} = nil
	var _ []int = nil
	var _ map[int]bool = nil
	var _ chan string = nil
	var _ func() = nil
	var _ interface{} = nil

	// // This following line doesn't compile.
	// var _ = nil

	tricks.Format("use func *new()")
	fmt.Println(*new(*int) == nil)         // true
	fmt.Println(*new(chan string) == nil)  // true
	fmt.Println(*new(interface{}) == nil)  // true

	fmt.Println(*new([]int) == nil)        // true
	fmt.Println(*new(map[int]bool) == nil) // true
	fmt.Println(*new(func()) == nil)       // true


	tricks.Format("package: unsafe")
	var p *struct{} = nil
	fmt.Println( unsafe.Sizeof( p ) ) // 8

	var s []int = nil
	fmt.Println( unsafe.Sizeof( s ) ) // 24

	var m map[int]bool = nil
	fmt.Println( unsafe.Sizeof( m ) ) // 8

	var c chan string = nil
	fmt.Println( unsafe.Sizeof( c ) ) // 8

	var f func() = nil
	fmt.Println( unsafe.Sizeof( f ) ) // 8

	var i interface{} = nil
	fmt.Println( unsafe.Sizeof( i ) ) // 16


	tricks.Format("not comparable")
	println(NOTE2)


	tricks.Format("pointer camparable with pointer, phải cùng type")
	type IntPtr *int
	// The underlying of type IntPtr is *int.
	var _ = IntPtr(nil) == (*int)(nil) //true

	tricks.Format("chan comparable with chan, phải cung type")
	var _ = (chan int)(nil) == (chan<- int)(nil) //true
	var _ = (chan int)(nil) == (<-chan int)(nil) //true


	tricks.Format("interface{} comparable with pointer, chan")
	// Every type in Go implements interface{} type.
	var _ = (interface{})(nil) == (*string)(nil) //false
	var _ = (interface{})(nil) == (chan string)(nil) //false
	var _ = (interface{})(nil) == (chan int)(nil) //false


	tricks.Format("slice map func() chi camparable with == nil")
	// The following lines compile okay.
	var _ = ([]int)(nil) == nil
	var _ = (map[string]int)(nil) == nil
	var _ = (func())(nil) == nil



}
const NOTE2 string = ` 
Compilation failure reason: mismatched types.
var _ = (*int)(nil) == (*bool)(nil)         // error ; defference pointer type
var _ = (chan int)(nil) == (chan bool)(nil) // error ; defference chan type
var _ = ([]int)(nil) == ([]int)(nil)		// error
var _ = (map[string]int)(nil) == (map[string]int)(nil) //error
var _ = (func())(nil) == (func())(nil) //error
`

const NOTE string = `
nil thuoc types: pointer slice map func() interface{} chan
`