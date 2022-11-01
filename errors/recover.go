package main

import (
	"fmt"
)

func main() {
	// defer func() {
	//     fmt.Println(recover()) // 3
	// }()

	// anonymous := func() string {
	//     return "anonymous"
	// }
	// fmt.Println(anonymous())

	// panic(3) // will replace panic 2
	// panic(2) // will replace panic 1
	// panic(1) // will replace panic 0
	// panic(0)

	sum(5, 0)                 //letgo to 1
	fmt.Println("main() end") //8
}

const (
	NOTE = `use keyword defer, recover() chi nằm trong func or anonymous func, chỉ bắt được func,
        NOT goroutine  `
)

func catch() {
	fmt.Println("catch() start") //5
	if err := recover(); err != nil {
		fmt.Println("err = ", err) //6
		// err = "runtime error: integer divide by zero"
	}
	fmt.Println("catch() end") //7
}

func sum(a int, b int) {
	fmt.Println("sum() start") //1
	defer catch()              //syntax: defer nameFunc() lun o tren cac func cua the xay ra panic, de bat no

	fmt.Printf("%d + %d = %d\n", a, b, a+b) //2
	divide(a, b)                            // se xay ra error=panic in func; //3
	fmt.Println("sum() end")
}

func divide(a int, b int) {
	fmt.Println("divide() start")           //4
	fmt.Printf("%d / %d = %d\n", a, b, a/b) // a / 0 ; error = panic ; jumpto func catch()
	fmt.Println("divide() end")
}
