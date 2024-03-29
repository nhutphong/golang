package main

import (
	"fmt"
	// "math"
	"reflect"
	"phong/tricks"
)


func main() {
	/*

		Arithmetic Operators: + - * / % ++ --
		a++ a--
		Assignment Operators: = += -= *= /= %= 
		Comparison Operators: == != < <= > >=
		Logical Operators:	: && || ! 				
		vd: !(x == y && x > z) có thể hiểu: (x != y || x < z)

		Bitwise Operators	: & | ^ << >> 			AND OR XOR


		CHECK TYPE:
		var name is type =  any | interface{}, moi dung duoc
			value, ok := ageInt.(int)
			value, ok := nameString.(string)
			value, ok := salaryFloat.(float64)


		bool

		string

		int int8 int16 int32 int64
		uint uint8 uint16 uint32 uint64 uintptr

		byte // alias for uint8

		rune // alias for int32
		     // represents a Unicode code point

		float32 float64

		complex64 complex128

		error

		any interface{}
		comparable: chưa nên dùng

		var myNil error = nil // nil thuoc error

		type varName any | interface{} // chỉ đẻ get value, không dùng được operations: + - *, ... != ,...


		type any = interface{} // alias
		type error interface {
			Error() string
		}
		declared any thi variable co the nhan bat ky kieu du lieu nao

		// func make() new()
		For slice, map and channel: use make()
		For arrays, struct and all value types: use new()

		byte rune: number or 1 char (chứa in quote ''): 'a', 'b', ; NOT double "a" "b" "A"

	*/


	tricks.Format("check type: nameVar.(type) ")
	var haha interface{} = 2.3
	switch v := haha.(type) {
	case int:
	    fmt.Println("int:", v)
	case float64:
	    fmt.Println("float64:", v)
	default:
	    fmt.Println("unknown")
	    fmt.Println()
	}

	tricks.Format("iota = increase, only use for const varName = iota")
	const (
		// iota = 0
		Failed = iota - 1 // == -1
		Unknown           // == 0
		Succeeded         // == 1
	)

	const (
		Readable = 1 << iota // == 1 * 2^0 = 1
		Writable             // == 1 * 2^1 = 2
		Executable           // == 1 * 2^2 = 4
	)
	fmt.Println(Readable, Writable, Executable)

	// chỉ declared, NOT assign value, thì default value = nil 
	tricks.Format("nil is default values of: <slice map channel func interface{} pointer>")
	var dict map[bool]bool //nil
    fmt.Print("Default Zero Value of a map: ")
    fmt.Println(dict)
    fmt.Println("dict == nil", dict == nil)

    var channel chan int //nil
    fmt.Print("Default Zero Value of a channel: ")
    fmt.Println(channel)

    var face interface{} //nil
    fmt.Print("Default Zero Value of a interface: ")
    fmt.Println(face)

    var slice []int //nil
    fmt.Println("slice == nil", slice == nil)
    fmt.Print("Default Zero Value of a slice: ")
    fmt.Println(slice)

    var f func() //nil
    fmt.Print("Default Zero Value of a func: ")
    fmt.Println(f)

    var pointer *int //nil
    fmt.Print("Default Zero Value of a pointer: ")
    fmt.Println(pointer)


	tricks.Format("check type: use packages reflect")
	var xyz interface{} = []int{1, 2, 3}
	xType := reflect.TypeOf(xyz)
	xValue := reflect.ValueOf(xyz)
	fmt.Println("reflect.TypeOf(xyz) is",xType) // []int
	fmt.Println("reflect.ValueOf(xyz) is",xValue) // [1 2 3]


	tricks.Format("type any")
	var my_number any = 100
	var my_any any = "thong dung"
	my_any = [...]int{1,2,3,4,5} // vi type la any nen co the assign type khac
	fmt.Println("my_number =",my_number, "\tmy_any =", my_any)


	tricks.Format("int")
	// nil == None python
	// khi declared name_var dung = thi pha co: var name string = 'thong'
	// name := 'thong' // := thi khong can keyword var
	var myInt8 int8 = 97

	/*
	  When you don't declare any type explicitly, the type inferred is `int`
	  (The default type for integers)
	*/
	var myInt = 1200

	var myUint uint = 500

	var myHexNumber = 0xFF  // Use prefix '0x' or '0X' for declaring hexadecimal numbers
	var myOctalNumber = 034 // Use prefix '0' for declaring octal numbers

	var myFloat32 float32 = 4.5
	var myFloat = 9.12 // // Type inferred as `float64` (the default type for floating-point numbers)
	fmt.Printf("%d, %d, %d, %#x, %#o %f %f\n", myInt8, myInt, myUint, myHexNumber, myOctalNumber, myFloat32, myFloat)


	// type rune byte chỉ là 1 char, sẽ được convert to hexa
	tricks.Format(`byte rune: dùng format: %c=string %U=unicode`)
	var myByte byte = 'a'
	var myRune rune = '♥'
	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)


	tricks.Format("so sanh")
	var myBoolean bool = true
	var anotherBoolean = false // Inferred type

	var truth = 3 <= 5
	var falsehood = 10 != 10

	// Short Circuiting
	var res4 = 10 > 20 && 5 == 5     // Second operand is not evaluated since first evaluates to false
	var res5 = 2*2 == 4 || 10%3 == 0 // Second operand is not evaluated since first evaluates to true
	fmt.Println(myBoolean, anotherBoolean, truth, falsehood, res4, res5)


	tricks.Format("complex64")
	// === Creating complex numbers ====
	/*
		complex64: both real and imaginary parts are of float32 type.
		complex128: both real and imaginary parts are of float64 type.
	*/
	var x complex64 = 3.4 + 2.9i
	var y = 5 + 7i // Type inferred as `complex128` (default type for complex numbers)
	fmt.Println(x, y)

	// Creating complex no from variables
	var a1 = 4.5
	var a2 = 7.1

	var c = complex(a1, a2) // a1 + a2*i won't work
	fmt.Println(c)

	// ===== Complex No Operations =====
	var one = 3 + 5i
	var two = 2 + 4i

	var res9 = one + two
	var res6 = one - two
	var res7 = one * two
	var res8 = one / two
	fmt.Println(res9, res6, res7, res8)


	tricks.Format("back-tics `raw string not render format`")
	// Normal String (Can not contain newlines, and can have escape characters like `\n`, `\t` etc)
	var website = "\thttps://www.callicoder.com\t\n"

	// Raw String (Can span multiple lines. Escape characters are not interpreted)
	var siteDescription = `\t\tCalliCoder is a programming blog where you can find
                           practical guides and tutorials on programming languages, 
                           web development, and desktop app development.\t\n`
	fmt.Println(website, siteDescription)


	tricks.Format("type Conversion")
	// Type Conversion
	var a3 int64 = 4
	var b1 int = int(a3) // Explicit Type Conversion

	var c1 float64 = 6.5

	// Explicit Type Conversion
	var result = float64( b1) +  c1 // Works
	fmt.Println(result)

	// ==================

	// The general syntax for converting a value v to a type T is T(v)
	var my_Int int = 65
	var my_Uint uint = uint(my_Int)
	var my_Float float64 = float64(my_Int)
	fmt.Println(my_Int, my_Uint, my_Float)


	tricks.Format("keywork const, multi declared variables")
	// Untyped Constant
	const myFavLanguage = "Python"
	const sunRisesInTheEast = true

	// Multiple declaration
	const country, code = "USA", 91

	const (
		employeeId string  = "E101"
		salary     float64 = 50000.0
	)

	// Typed Constant
	const typedInt int = 1234
	const typedStr string = "Hi"
	fmt.Println(myFavLanguage, sunRisesInTheEast, country, code, employeeId, salary, typedInt, typedStr)


}

const NOTE string = `
	type (
		MyInt int
		Age   int
		Text  string
	)	

	type IntPtr *int
	type Book struct{author, title string; pages int}
	type Convert func(in0 int, in1 bool)(out0 int, out1 string)
	type StringArray [5]string
	type StringSlice []string

	// The underlying types of the following ones are both int.
	type (
		MyInt int
		Age   MyInt
	)

	// The following new types have different underlying types.
	type (
		IntSlice   []int   // underlying type is []int
		MyIntSlice []MyInt // underlying type is []MyInt
		AgeSlice   []Age   // underlying type is []Age
	)

	// The underlying types of []Age, Ages, and AgeSlice
	// are all the unnamed type []Age.
	type Ages AgeSlice

	
`
