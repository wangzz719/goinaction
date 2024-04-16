package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	rcWg    sync.WaitGroup
)

func main() {
	rcWg.Add(2)

	go incrCounter(1)
	go incrCounter(2)

	rcWg.Wait()
	fmt.Println("Final counter: ", counter)
}

func incrCounter(id int) {
	defer rcWg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		runtime.Gosched()

		value++
		counter = value
	}
}
