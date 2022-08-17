package main

import (
	"fmt"
	"time"
)

/*
	goroutines giong voi async def myfunc python
	dung keywork go func_name()
*/

func numbers() {
	for i := 1; i <= 10; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}

func alphabets() {
	for i := 'a'; i <= 'g'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}

func main() {
	go numbers() // run lun
	go alphabets() run lun

	// run toi lun neu ko co time.Sleep()
	fmt.Println("end goroutines!!!")
	fmt.Println("end goroutines!!!")
	fmt.Println("end goroutines!!!")

/* 
	bac buoc phai co time.Sleep() de 2 goroutines numbers va alphabets co
	time scheduler va run
	goroutine main() phai co time.Sleep() > time.Sleep() của các goroutines trong nó như: 
	go numbers() va go alphabets()
	or call func() nào có thời gian chạy lâu hơn các goroutines phía trên nếu không sẽ run hết func main()
	các phần còn lại trong goroutines sẽ đc bỏ qua
*/ 
	
	//  
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("main terminated")
}
