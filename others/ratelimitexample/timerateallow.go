package main

import (
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	r := rate.Every(100 * time.Millisecond) // rate.Every 指每隔多长时间向桶中新增一个 token
	limit := rate.NewLimiter(r, 10)         // 创建一个令牌桶，最大容量是 10 个 token，每隔 r 时间新增一个令牌

	fmt.Println("maximum tokens: ", limit.Burst(), " limit: ", limit.Limit())

	c := make(chan int, 100)
	go func() {
		start := time.Now()
		i := 0

		for i < 100 {
			if limit.Allow() {
				c <- i
				i++
			}
		}

		close(c) // if not close, deadlock will occur

		fmt.Println(time.Now().Sub(start))
	}()

	for i := range c {
		fmt.Println(i)
	}

	time.Sleep(5 * time.Second)
}
