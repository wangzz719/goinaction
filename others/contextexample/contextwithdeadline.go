package main

import (
	"context"
	"fmt"
	"time"
)

func slowOperationDeadline(ctx context.Context, duration time.Duration) {
	timeout := time.After(duration * time.Second) // 如果在 case 中使用 time.After，每次 case 都会重新计时，会导致 case 无法生效

	for {
		select {
		case <-ctx.Done():
			fmt.Println("context done: ", ctx.Err().Error())
			return
		case <-timeout:
			fmt.Println("operation done")
			return
		default:
			fmt.Println("doing...")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(10*time.Second))
	defer cancel()

	go slowOperationDeadline(ctx, 20)

	time.Sleep(10 * time.Second)
	fmt.Println("context cancel")
	time.Sleep(5 * time.Second)
}
