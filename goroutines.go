package main

import (
	"fmt"
	"time"
)

// goroutines giong voi async def myfunc python
//

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go numbers() // run 1 goroutines
	go alphabets()

	// run toi lun neu ko co time.Sleep()
	fmt.Println("end goroutines!!!")
	fmt.Println("end goroutines!!!")
	fmt.Println("end goroutines!!!")

	// bac buoc phai co time.Sleep() de 2 goroutines numbers va alphabets co
	// time scheduler va run
	// goroutine main() phai co time.Sleep() > time.Sleep() cua goroutine no
	// call la go numbers() va go alphabets()

	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
