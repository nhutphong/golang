package main

import (
	"context"
	"fmt"
	"time"
)

const LOOP int = 5

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	parentSomeThing(ctx)

	time.Sleep(time.Second)
	fmt.Println("\t\t\tmain() end")
}


func parentSomeThing(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	// defer cancel()
	
	printCh := make(chan int)
	go childSomeThing(ctx, printCh)

	for num := 0; num < LOOP; num++ {
		printCh <- num+1
	}

	cancel()
	time.Sleep(500 * time.Millisecond)

	fmt.Printf("\t\tparentSomeThing: end\n")
}

func childSomeThing(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("childSomeThing err: %s\n", err)
			}
			fmt.Printf("\tchildSomeThing: end\n")
			return
		case num := <-printCh:
			fmt.Printf("childSomeThing: %d\n", num)
		}
	}
}