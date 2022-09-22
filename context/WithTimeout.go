package main

import (
	"context"
	"fmt"
	"time"
)

func doSomethingCool(ctx context.Context) {
	var index int
	for {
		index++
		fmt.Println("index", index)
		
		select {
		case <-ctx.Done():
			fmt.Println("run du 2s thoat thoi")
			err := ctx.Err()
			fmt.Println(err)
			return
		default:
			fmt.Println("defalut select")
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Go Context Tutorial")
	//ctx=event run 2s la thoat
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	go doSomethingCool(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("\nctx.Done() main()\n")
	}

	fmt.Println("main end")
	time.Sleep(3 * time.Second)

	fmt.Println("main end2")
	time.Sleep(1 * time.Second)
	fmt.Println("main end3")
}
