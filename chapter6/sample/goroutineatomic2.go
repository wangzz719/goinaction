package main

import (
	"sync"
	"fmt"
	"sync/atomic"
	"time"
)

var (
	shutdown int64
	atomWg2  sync.WaitGroup
)

func main() {
	atomWg2.Add(2)
	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)
	fmt.Println("Shutdown now")
	atomic.StoreInt64(&shutdown, 1)
	atomWg2.Wait()
}

func doWork(name string) {
	defer atomWg2.Done()
	for {
		fmt.Printf("Doing %s work\n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1 {
			fmt.Printf("Shutting %s down\n", name)
			break
		}
	}
}
