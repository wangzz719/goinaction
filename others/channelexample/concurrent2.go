package main

import (
	"fmt"
	"time"
)

func run2(taskId, sleepTime int, ch chan string) {
	time.Sleep(time.Duration(sleepTime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", taskId, sleepTime)
	return
}
func main() {
	input := []int{3, 2, 1}
	chs := make([]chan string, len(input))
	startTime := time.Now()
	fmt.Println("Multirun start")
	for i, sleepTime := range input {
		chs[i] = make(chan string)
		go run2(i, sleepTime, chs[i])
	}
	for _, ch := range chs {
		fmt.Println(<-ch)
	}
	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of tasks is %d\n", endTime.Sub(startTime), len(input))
}
