package main

import (
	"sync"
	"runtime"
	"fmt"
)

var (
	counterMutex int
	mutexWg      sync.WaitGroup
	mutex        sync.Mutex
)

func main() {
	mutexWg.Add(2)
	go incrCounterMutex(1)
	go incrCounterMutex(2)

	mutexWg.Wait()
	fmt.Println("Final counter: ", counterMutex)
}

func incrCounterMutex(id int) {
	defer mutexWg.Done()
	for count := 0; count < 2; count++ {
		mutex.Lock() // 加锁
		{
			value := counterMutex
			runtime.Gosched()
			value ++
			counterMutex = value
		}
		mutex.Unlock() // 释放锁
	}
}
