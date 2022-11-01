package main

import (
	"fmt"
)

func main() {
	var b = Book{pages: 123}
	var p = &b // pointer
	var f1 = b.Pages
	var f2 = p.Pages
	b.pages = 789
	var g1 = p.Pages2 // co pointer receiver
	var g2 = b.Pages2 // co pointer receiver
	fmt.Println(f1()) // 123
	fmt.Println(f2()) // 123
	fmt.Println(g1()) // 789 ; update
	fmt.Println(g2()) // 789 ; update

}

// struct-non giong nhu alias cua types ma ta muon gan thui
type Age int

func (age Age) LargerThan(a Age) bool {
	return age > a
}
func (age *Age) Increase() {
	*age++ // phai dung operator *, de update, giong nhu types: string slice, ...,
}
func (age *Age) Reset() {
	*age = 5
	// *age=Age(5) // OK
}
func (age *Age) Set(a Age) {
	*age = a
}

// func (age *Age) Set(a int) {
// 	*age = Age(a)
// }

// Receiver of custom defined function type.
type FilterFunc func(in int) bool

func (f FilterFunc) Filte(in int) bool {
	return f(in)
}

// Receiver of custom defined map type.
type StringSet map[string]struct{}

func (ss StringSet) Has(key string) bool {
	_, present := ss[key]
	return present
}
func (ss StringSet) Add(key string) {
	ss[key] = struct{}{} // struct empty
}
func (ss StringSet) Remove(key string) {
	delete(ss, key)
}

// Receiver of custom defined struct type.
type Book struct {
	pages int // field=attribute=state
}

func (b Book) Pages() int {
	return b.pages
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}

func (b *Book) Pages2() int {
	return b.Pages()
}
