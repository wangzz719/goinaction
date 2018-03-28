package main

import (
	"log"
	"os"
	"io/ioutil"
	"io"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	file, err := os.OpenFile("errors.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}

	// Discard是一个io.Writer，所有的Write调用都不会有动作，但是会成功返回
	// var Discard io.Writer = devNull(0)
	// 基于 devNull 类型实现的 Write 方法，会忽略所有写入这一变量的数据。
	// 当某个等级的日志不重要时，使用 Discard 变量可以禁用这个等级的日志。
	Trace = log.New(ioutil.Discard, "TRACE: ", log.Ldate|log.Lmicroseconds|log.Llongfile)

	Info = log.New(os.Stdout, "INFO: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
	Warning = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Lmicroseconds|log.Llongfile)

	//MultiWriter 函数是一个变参函数，可以接受任意个实现了 io.Writer 接口的值。
	// 这个函数会返回一个 io.Writer 值，这个值会把所有传入的 io.Writer 的值绑在一起。
	// 当对这个返回值进行写入时，会向所有绑在一起的 io.Writer 值做写入。
	// 这让类似 log.New 这样的函数可以同时向多个 Writer 做输出。
	// 现在，当我们使用 Error 记录器记录日志时，输出会同时写到文件和 stderr。
	Error = log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", log.Ldate|log.Lmicroseconds|log.Llongfile)
}

func main() {
	Trace.Println("I have something standard to say")
	Info.Println("Special Information")
	Warning.Println("There is something you need to know about")
	Error.Println("Something has failed")
}
