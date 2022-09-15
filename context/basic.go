package main

import (
	"fmt"
	"context"
)

const NOTE = ` 
context=ctx=map=dict

type Context interface {
    // Channel listen to cancellation
    Done() <-chan struct{}

    // Return error if context is cancelled
    Err() error
    
    // Return the deadline set to the context
    Deadline() (deadline time.Time, ok bool)
  
    // Return the value for a given key
    Value(key interface{}) interface{}
}

// Return empty root context
func Background() Context
func TODO() Context

// Return cancellable context
func WithCancel(parent Context) (ctx Context, cancel CancelFunc)
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc)
func WithDeadline(parent Context, d time.Time) (Context, CancelFunc)

// Return context with key-value pairs
func WithValue(parent Context, key, val interface{}) Context

`

func main() {
	ctx := context.Background()
	// ctx2 := context.TODO()

	// ctx["myKey"] = "myValue" ; return context
	ctx = context.WithValue(ctx, "myKey", "myValue")
	ctx = context.WithValue(ctx, "name", "chi thong")
	ctx2 := context.WithValue(ctx, "name", "thanh dung")

	keys := []string{"myKey", "name"}
	for _, key := range keys {
		fmt.Printf("key:value %s\n", ctx.Value(key))
		fmt.Printf("key:value %s\n", ctx2.Value(key))
	}
}

func doSomething(ctx context.Context) {
	// ctx.Value("myKey") == ctx["myKey"]
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
}