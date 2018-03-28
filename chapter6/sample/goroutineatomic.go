package main

import (
	"sync"
	"runtime"
	"fmt"
	"sync/atomic"
)

var (
	counteratom int64
	atomWg      sync.WaitGroup
)

func main() {
	atomWg.Add(2)
	go incrCounterAtom(1)
	go incrCounterAtom(2)

	atomWg.Wait()
	fmt.Println("Final counter: ", counteratom)
}

func incrCounterAtom(id int) {
	defer atomWg.Done()
	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counteratom, 1)
		runtime.Gosched()
	}
}
