package main

import "log"

func init() {
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

func main() {
	log.Println("message") 			// Println写到标准日志记录器
	log.Fatalln("fatal message")		// Fatalln 在调用 Println()之后会接着调用 os.Exit(1)
	log.Panic("panic message")		// Panicln 在调用 Println()之后会接着调用 panic()
}

// Fatal 系列函数用来写日志消息，然后使用 os.Exit(1)终止程序。
// Panic 系列函数用来写日志消息，然后触发一个 panic。 除非程序执行 recover 函数，否则会导致程序打印调用栈后终止。
// Print 系列函数是写日志消息的标准方法。
// 日志记录器是多 goroutine 安全的。这意味着在多个 goroutine 可以同时调用来自同一个日志记录器的这些函数，
// 而不会有彼此间的写冲突