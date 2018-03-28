package main

import (
	"time"
	"fmt"
)

func Run1(taskId, sleepTime, timeout int, ch chan string) {
	chRun := make(chan string)
	go run4(taskId, sleepTime, chRun)
	select {
	case re := <-chRun:
		ch <- re
	case <-time.After(time.Duration(timeout) * time.Second):
		re := fmt.Sprintf("task id %d , timeout", taskId)
		ch <- re
	}
}

func run4(taskId, sleepTime int, ch chan string) {
	time.Sleep(time.Duration(sleepTime) * time.Second)
	ch <- fmt.Sprintf("task id %d , sleep %d second", taskId, sleepTime)
	return
}

func main() {
	input := []int{3, 2, 1}
	timeout := 2
	chLimit := make(chan bool, 2)
	chs := make([]chan string, len(input))

	limitFunc := func(chLimit chan bool, ch chan string, taskId, sleeptime, timeout int) {
		Run1(taskId, sleeptime, timeout, ch)
		<-chLimit
	}

	startTime := time.Now()
	fmt.Println("Multirun start")

	for i, sleepTime := range input {
		chs[i] = make(chan string, 1) // an buffered channel need, otherwise will lead to deadlock
		chLimit <- true
		go limitFunc(chLimit, chs[i], i, sleepTime, timeout)
	}

	for _, ch := range chs {
		fmt.Println(<-ch)
	}

	endTime := time.Now()
	fmt.Printf("Multissh finished. Process time %s. Number of task is %d", endTime.Sub(startTime), len(input))
}

// chs[i] = make(chan string, 1) an buffered channel need, otherwise will lead to deadlock
// if we let chs[i] = make(chan string)
// if chLimit <- true is blocked, then chs[i] consume function cannot be reached, this will lead to deadlock
// chs[i] consume function is :
// for _, ch := range chs {
//     fmt.Println(<-ch)
// }
