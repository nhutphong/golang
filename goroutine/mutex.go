package main

import (
	"fmt"
	"sync"
)

const NOTE string  = ` 
 	sync.WaitGroup: đảm bảo các goroutines sẽ hoàn thành trước các SuperGoroutine() or sau main()
 	sync.Mutex{}: xử lý race condition và shared data giữa các goroutines
 `


var x int64 = 0
// Khai báo mutex
var mutex = &sync.Mutex{}

func write(wg *sync.WaitGroup) {
	// defer wg.Done()

	// Lock lại
	mutex.Lock()
	x = x + 1
	// code trong đây sẽ được từng goroutine xử lý, không tranh giành 
	// Unlock
	mutex.Unlock()

	wg.Done() // lun nằm ở cuối goroutine 
}

func main() {
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(2)
		go write(&w)
		go write(&w)
	}

	w.Wait() // nằm sau các đoạn code: go childGoroutines()
	fmt.Println("Giá trị của x là: ", x)
}

/*
	// The capacity must be one.
	mutex := make(chan struct{}, 1)

	counter := 0
	increase := func() {
		mutex <- struct{}{} // lock
		counter++
		<-mutex // unlock
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done; <-done
	fmt.Println(counter) // 2000
*/