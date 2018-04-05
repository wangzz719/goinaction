package main

import (
	"github.com/juju/ratelimit"
	"time"
	"fmt"
)

func main()  {
	limit := ratelimit.NewBucket(100*time.Millisecond, 10)
	fmt.Println("maximum tokens: ", limit.Capacity(), " rate: ", limit.Rate(), " avaliable tokens:", limit.Available())
	c := make(chan int, 100)
	go func() {
		start := time.Now()
		for i := 0; i < 100; i++ {
			limit.Wait(1)
			c <- i
		}
		close(c) // if not close, deadlock will occur
		fmt.Println(time.Now().Sub(start))
	}()
	for i := range c {
		fmt.Println(i)
	}
	time.Sleep(5 * time.Second)
}
