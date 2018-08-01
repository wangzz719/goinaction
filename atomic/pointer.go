/*

使用 atomic.StorePointer 和 atomic.LoadPointer 实现对共享变量的读写操作

 */

package main

import (
"unsafe"
"time"
"fmt"
"sync/atomic"
"os"
)

type Config struct {
	Name  string
	Value int64
}

var DefaultConfig unsafe.Pointer

func ChangeConfig() {
	ticker := time.NewTicker(time.Second * 2)
	for range ticker.C {
		config := &Config{
			Name:  fmt.Sprintf("name: %d", time.Now().Unix()),
			Value: time.Now().Unix(),
		}
		atomic.StorePointer(&DefaultConfig, unsafe.Pointer(config))
	}
}

func SetConfig() {
	config := &Config{Name: "name: 123", Value: 123}
	atomic.StorePointer(&DefaultConfig, unsafe.Pointer(config))
}

func GetConfig() {
	ticker := time.NewTicker(time.Second * 1)
	for range ticker.C {
		config := (*Config)(atomic.LoadPointer(&DefaultConfig))
		fmt.Println(config.Name, config.Value)
	}
}

func main() {
	SetConfig()
	go ChangeConfig()
	go GetConfig()
	go GetConfig()

	interrupt := make(chan os.Signal, 1)
	<-interrupt

	fmt.Println("shut down")
}

