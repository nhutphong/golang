package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

//import "folder/to/folder1"
// folder1.class1, folder1.func1, ..., truy cap duoc tat ca cac class, func cua tat ca cac files trong folder1
// ko can phai thong qua file_name, chi can folder_name.

func add(a int, b int) int {
	return a + b
}

func swap(a, b string) (string, string) {
	return b, a
}

func split(sum int) (y, x int) {
	x = sum * 4 / 9
	y = sum - x
	// return y, x
	return //nen ghi ro rang
}

func forloop1(stop int) int {
	sum := 0
	// for i,j := 0,1; i < 10 && j < 10_; i,j = i+1, j+1  {
	// }
	for i := 0; i < stop; i++ {
		sum += i
	}
	return sum
}

func forloop2(total int) int {
	sum := 1
	// for condition == while
	for sum < total {
		sum += sum
	}
	return sum
	// for {
	// vo tan ko co dk dung
	// }
}

func forever() {
	for {

	}
	// for {
	// vo tan ko co dk dung
	// }
}

func pownhut(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
		// can't use v here, though
	}
	return lim
}

func osmain() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func weekday() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func morning() {
	t := time.Now()
	fmt.Println(t)
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func deferNhut(stop int) {
	fmt.Println("start")
	for i := 0; i < stop; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("end")
}

func point() {
	i, j := 42, 2701

	p := &i                             // point to i
	fmt.Println("*p =", *p, "=> value") // read i through the pointer
	fmt.Println("p =", p, "=> id=memory")
	*p = 21        // set i through the pointer
	fmt.Println(i) // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

func main() {
	// var c, python, java bool // bool:=false
	// var i int // i:=0
	fmt.Println("Hello, how are you?")
	fmt.Println("Bay gio la: ", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	fmt.Println("Tao la math.Pi :", math.Pi)
	fmt.Println("Tao la func add(a,b) ", add(3, 8))

	a, b := swap("nhut", "tan")
	fmt.Println("Tao la swap(a,b) = ", a, b)

	fmt.Println(split(17))
	fmt.Println("for_loop1(stop) =", forloop1(10))
	fmt.Println("for_loop2 =", forloop2(1000))
	defer osmain()
	//defer <expr> => lam expr chay cuoi chung trinh
	weekday()
	morning()
	deferNhut(11)
	point()
}
