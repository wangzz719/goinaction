# Chapter1

## goroutine
goroutine 是可以与其他 goroutine 并行执行的函数，同时也会与主程序(程序的入口)并行执行。
goroutine 使用的内存比线程更少，Go 语言运行时会自动在配置的一组逻辑处理器上调度执行 goroutine。
每个逻辑处理器绑定到一个操作系统线程上

```go
// 启动 goroutine
go func()
```

## Channel
通道是一种数据结构，可以让 goroutine 之间进行安全的数据通信。

goroutine 安全通信理念：
`
通过通信来共享变量，而不是通过共享变量来通信
`

这种在 goroutine 之间安全传输数据的方法不需要任何锁或者同步机制。

**需要强调的是，channel 并不提供跨 goroutine 的数据访问保护机制**

## Type System
Go 开发者使用组合(composition)设计模式，只需简单地将一个类型嵌入到另一个类型，就能复用所有的功能。
在 Go 语言中，一个类型由其他更微小的类型组合而成，避免了传统的基于继承的模型。

### Interface
在 Go 语言中，如果一个类型实现了一个接口的所有方法，那么这个类型的实例就可以存储在这个接口类型的实例中，不需要额外声明。
```
鸭子类型——如果它叫起来像鸭子，那它就可能是只鸭子
```