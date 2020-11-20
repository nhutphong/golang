package main

import "fmt"

func main() {
	/*
		The length of the slice is the number of elements in the slice.
		slice=cap la do dai tu start_index den end cua array no tham tri=value
	*/
	a := []int{10, 20, 30, 40, 50, 60, 80, 100, 120, 130, 150}
	s := a[1:6]
	last := s[2:]

	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
	fmt.Printf("s = %v, len = %d, cap = %d\n", last, len(last), cap(last))
}
