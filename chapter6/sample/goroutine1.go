package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Start goroutine")
	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char ++ {
				fmt.Printf("%c\n", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char ++ {
				fmt.Printf("%c\n", char)
			}
		}
	}()

	fmt.Println("Waiting to finish")

	wg.Wait()
	fmt.Println("\n Terminating program")
}
