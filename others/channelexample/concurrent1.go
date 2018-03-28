package main

import (
	"fmt"
	"time"
)

func run(taskId, sleepTime int, ch chan string) {
	time.Sleep(time.Duration(sleepTime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", taskId, sleepTime)
	return
}
func main() {
	input := []int{3, 2, 1}
	ch := make(chan string)
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleepTime := range input {
		go run(i, sleepTime, ch)
	}
	for range input {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of tasks is %d\n", endTime.Sub(startTime), len(input))
}
