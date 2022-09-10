package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

const NOTE string = ` 
	Race condition: hiện tượng nhiều tiến trình cùng truy cập và muốn thay đổi giá trị của biến, 
	nhưng không theo quy tắc nào, khiến kết quả không như mong muốn.
	Golang xử lý race condition bằng việc dùng atomic hoặc mutex hoặc channel.

	atomic: chỉ hỗ trợ số nguyên int 
`

var x int64 = 0

func addOne(wg *sync.WaitGroup) {
        // Xài hàm của atomic để tăng giá trị.
	//x = x + 1
	atomic.AddInt64(&x, 1)
	wg.Done()
}
func main() {
	var w sync.WaitGroup

	for i := 0; i < 1000; i++ {
		w.Add(1)
		go addOne(&w)
	}

	w.Wait()
	fmt.Println("Giá trị của x là: ", x)
}