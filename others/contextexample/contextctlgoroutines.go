package main

import (
	"context"
	"fmt"
	"time"
)

func Run(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, " context done: ", ctx.Err().Error())
			return
		default:
			fmt.Println(name, " doing...")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go Run(ctx, "goroutine1")
	go Run(ctx, "goroutine2")
	go Run(ctx, "goroutine3")

	time.Sleep(10 * time.Second)

	fmt.Println("context cancel")
	cancel()

	time.Sleep(5 * time.Second)
}
