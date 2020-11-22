package main

import (
	"fmt"
)

type Point struct {
	X, Y float64
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

func main() {
	p := Point{3, 4}
	fmt.Println("Point p = ", p)

	p.Method(7, 9)
	fmt.Println("After Translate, p = ", p)

	one, two := p.NoPointer(10, 20)
	fmt.Println("after StaticMethod, p = ", p)
	fmt.Printf("one = %v, two = %v", one, two)
}
