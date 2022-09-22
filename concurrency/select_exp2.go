package main

import (
	"fmt"
	"time"
)

func process(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func loop() float64 {
	ketqua := 100.0
	for i := 0; i < 5; i++ {
		if i==4 {
			return ketqua
			// return //error, phai co named return, moi return kieu nay duoc
		} else {
			fmt.Println(i)
		}
	}

	// end func return phai co, bat buoc
	return 500.0
}


func main() {
	fmt.Println(loop())
	fmt.Println("main end")

}
