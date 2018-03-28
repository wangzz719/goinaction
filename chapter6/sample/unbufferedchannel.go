package main

import (
	"sync"
	"math/rand"
	"time"
	"fmt"
)

var channelWg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	court := make(chan int)

	channelWg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)
	court <- 1
	channelWg.Wait()
}

func player(name string, court chan int) {
	defer channelWg.Done()
	for {
		ball, ok := <-court
		if ! ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(court)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball ++
		court <- ball
	}
}
