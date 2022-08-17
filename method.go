package main

import (
	"fmt"
)

// Struct type - `Point`
type Point struct {
	X, Y float64
}

/*
 Method with receiver=self `Point` chi access, get value, NOT UPDATE

*/
func (p Point) IsAbove(y float64) bool {
	return p.Y > y
}

func IsAboveFunc(p Point, y float64) bool {
	return p.Y > y
}

// method co pointer
// in-place, update=add del change chinh no
// cac thao tac set
func (p *Point) Method(dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

// NO pointer chi get value tu field=attr ko the add del change
func (p Point) NoPointer(dx, dy float64) (X, Y float64) {
	X = p.X + dx
	Y = p.Y + dy
	return
	// return X, Y
}

//func(p *Point) pass *Pointer nen update X, Y
func TranslateFunc(p *Point, dx, dy float64) {
	p.X = p.X + dx
	p.Y = p.Y + dy
}

func format() {
	fmt.Println()
	fmt.Println("--------------------------------------------------------------------------------")
	fmt.Println()
}

type MyString string

func (myStr MyString) reverse() string {
	s := string(myStr)
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}


func main() {
	p := Point{2.0, 4.0}

	fmt.Println("Point : ", p)

	fmt.Println("Is Point p located above the line y = 1.0 ? : ", p.IsAbove(2))

	format()

	p1 := Point{2.5, -3.0}

	fmt.Println("Point : ", p1)

	fmt.Println("Is Point p1 located above the line y = 1.0 ? : ", IsAboveFunc(p1, 1))

	format()

	p2 := Point{3, 4}
	fmt.Println("Point p = ", p2)

	p2.Method(7, 9) // pointer *receiver
	fmt.Println("After Translate, p2 = ", p2)

	one, two := p.NoPointer(10, 20) // receiver not update X va Y
	fmt.Println("after StaticMethod, p2 = ", p2)
	fmt.Printf("one = %v, two = %v", one, two)

	format()

	p3 := Point{5, 8}
	fmt.Println("Point p = ", p3)

	TranslateFunc(&p3, 7, 10)
	fmt.Println("TranslateFunc(&p3, 7, 10) = ", p3)

	format()
	myStr := MyString("OLLEH")
	fmt.Println(myStr.reverse())
}
