package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	// context.WithCancel return a new context and a function to cancel context
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}

	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel = context.WithDeadline(context.Background(), d)
	defer cancel()

	select {
	case <- time.After(1 * time.Second):
		fmt.Println("overSlept")
	case <- ctx.Done():
		fmt.Println(ctx.Err())
		fmt.Println(ctx.Deadline())
	}

	type vlcontext string
	f := func(ctx context.Context, k vlcontext) {
		if a := ctx.Value(k); a != nil {
			fmt.Println("found key value: ", a)
		} else {
			fmt.Println("not found key value", k)
		}
	}

	a := vlcontext("language")
	ctx = context.WithValue(context.Background(), a , "golang")
	f(ctx, a)
	f(ctx, vlcontext("color"))
}
