package main

import (
	"context"
	"fmt"
	"time"
)

func parentSomeThing(ctx context.Context) {
	var index int
	for {
		index++
		fmt.Println("index", index)
		
		select {
		case <-ctx.Done():
			fmt.Println("run du 2s thoat thoi")
			err := ctx.Err()
			fmt.Println(err)
			fmt.Println("\tparentSomeThing end")
			return // end func
		default:
			fmt.Println("defalut select")
		}

		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	fmt.Println("Go Context Tutorial")
	//ctx=event run 2s la thoat
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go parentSomeThing(ctx)

	select {
	case <-ctx.Done():
		fmt.Println("\nmain() ctx.Done()()\n")
	}

	time.Sleep(time.Second) // jumpto <-ctx.Done() in parentSomeThing
	fmt.Println("\t\tmain end")
}
