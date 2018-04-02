package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("context done: ", ctx.Err().Error())
				return
			default:
				fmt.Println("doing...")
				time.Sleep(2 *time.Second)
			}
		}
	}(ctx)
	time.Sleep(10 * time.Second)
	fmt.Println("context cancel")
	cancel()
	time.Sleep(5 * time.Second)
}