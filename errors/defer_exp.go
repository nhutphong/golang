package main

import (
	"fmt"
	"errors"
)

func main() {
	// not()

	solution()
	solution2()
	solution3()

	fmt.Println("notRelease()", notRelease())
	fmt.Println("Release()", release())
	fmt.Println("main() end")
}

func not() {
	for i := 0; i < 3; i++ {
		defer func() {
			fmt.Println(i)
		}() // cac def func(), chi nhin thay i cuoi cung
	}
	fmt.Println("not() end")
}

func solution() {
	for i := 0; i < 3; i++ {
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("solution() end")
}


func solution2() {
	for i := 0; i < 3; i++ {
		i := i
		defer func(i int) {
			fmt.Println(i)
		}(i)
	}
	fmt.Println("solution2() end")
}

func solution3() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("solution3() end")
}


func notRelease() error {
	defer func() error {
		return errors.New("error") // not return 
	}()

	return nil
}

// use named return 
func release() (err error) {
	defer func() {
		err = errors.New("xin chao err") //
	}() // return err

	return nil
}
