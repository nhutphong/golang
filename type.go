package main

import (
	"fmt"
	"math"
	"phong/tricks"
)

func format() {
    fmt.Println()
    fmt.Println("--------------------------------------------------------------------------------")
    fmt.Println()
}

func main() {
	/*

		type any = interface{}
		type error interface {
			Error() string
		}
		declared any thi variable co the nhan bat ky kieu du lieu nao

		For slices, maps and channels: use make()
		For arrays, structs and all value types: use new()

	*/
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


	tricks.Format("byte rune")
	var myByte byte = 'a'
	var myRune rune = '♥'

	fmt.Printf("%c = %d and %c = %U\n", myByte, myByte, myRune, myRune)


	var a, b = 4, 5
	var res1 = (a + b) * (a + b) / 2 // Arithmetic operations

	a++ // Increment a by 1

	b += 10 // Increment b by 10

	var res2 = a ^ b // Bitwise XOR

	var r = 3.5
	var res3 = math.Pi * r * r // Operations on floating-point type

	fmt.Printf("res1 : %v, res2 : %v, res3 : %v\n", res1, res2, res3)

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


	tricks.Format("keywork const")
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
