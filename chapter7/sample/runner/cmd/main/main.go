package main

import (
	"time"
	"log"
	"github.com/wangzz719/goinaction/chapter7/sample/runner/pkg/runner"
	"os"
)

const timeout = 3 * time.Second

func main() {
	log.Println("Starting work.")
	r := runner.New(timeout)
	r.Add(createTask(), createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			log.Println("Termination due to timeout.")
			os.Exit(1)
		case runner.ErrInterrupt:
			log.Println("Termination due to interrrupt")
			os.Exit(2)
		}
	}
	log.Println("Process ended.")
}

func createTask() func(int) {
	return func(id int) {
		log.Printf("Processor - Task #%d.\n", id)
		time.Sleep(time.Duration(id) * time.Second)
	}
}
