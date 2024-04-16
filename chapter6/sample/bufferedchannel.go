package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const numberGoroutines = 4
const taskLoad = 10

var bufChannelWg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	tasks := make(chan string, taskLoad)

	bufChannelWg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	fmt.Println("Close Tasks")
	close(tasks)

	bufChannelWg.Wait()
}

func worker(tasks chan string, worker int) {
	defer bufChannelWg.Done()

	for {
		task, ok := <-tasks
		if !ok {
			fmt.Printf("Worker: %d : Shutting Down\n", worker)

			return
		}

		fmt.Printf("Worker: %d : Started %s\n", worker, task)

		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		fmt.Printf("Worker: %d : Completed %s\n", worker, task)
	}
}
