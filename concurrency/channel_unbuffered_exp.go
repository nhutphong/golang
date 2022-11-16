package main

import (
	"fmt"
	"time"
)

const LOOP = 4

func main() {
	ch := make(chan int)
	defer close(ch)
	name := "one"
	name2 := "two"
	name3 := "three"

	go write(LOOP, ch, name)
	go write(LOOP, ch, name2)
	go write(LOOP, ch, name3)

	go read(LOOP, ch, name)
	go read(LOOP, ch, name2)
	go read(LOOP, ch, name3)

	fmt.Println("start main()")
	time.Sleep(time.Second)
	fmt.Println("end main()")
}

func read(loop int, ch chan int, name string) {
	fmt.Printf("\t%v read start\n", name)
	for i := 0; i < loop; i++ {
		fmt.Printf("\t\t%v read for start%v\n", name, i)
		fmt.Printf("\t\t%v read ch %v\n", name, <-ch)
		fmt.Printf("\t\t%v read for end%v\n", name, i)
	}
	fmt.Printf("\t%v read end\n", name)
}

func write(loop int, ch chan int, name string) {
	fmt.Printf("\t%v write start\n", name)

	for i := 0; i < loop; i++ {
		fmt.Printf("\t\t%v write for start%v\n", name, i)
		ch <- i
		fmt.Printf("\t\t%v write for end%v\n", name, i)
	}

	fmt.Printf("\t%v write end\n", name)
}
