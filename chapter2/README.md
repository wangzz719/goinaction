# Chapter2

## Sample

### Entrance
go 语言的程序入口是 main package 里的 `main()` 函数，如
```go
package main

func main(){ // entrance
}
```
chapter/sample/main.go 是示例程序入口

### package
Go 语言的每个代码文件都属于一个包，一个包定义一组编译过的代码，包的名字类似命名空间，可以用来间接访问包内声明的标识符。
包的声明在 go 文件的第一行， package 关键字开头，随后跟着包的名字。按照惯例，包和文件夹同名。
在 Go 语言里，标识符要么从包里公开，要么不从包里公开。公开标识符以大写字母开头，不公开标识符以小写字母开头。

### import
关键字 import 就是导入一段代码，让用户可以访问其中的标识符，如类型、函数、常量和接口。
所有处于同一个文件夹里的代码文件，必须使用同一个包名。

### init
所有代码文件里的 init 函数都会在 main 函数执行前调用。

### defer
延迟执行函数。关键字 defer 会安排随后的函数调用在函数返回时才执行。

### 变量零值
在 Go 语言中，所有变量都被初始化为其零值。
对于数值类型，零值是 0;对于字符串类型，零值是空字符串;对于布尔类型，零值是 false;
对于指针，零值是 nil。对于引用类型来说，所引用的底层数据结构会被初始化为对应的零值。
但是被声明为其零值的引用类型的变量，会返回 nil 作为其值。

### make vs new
new 和 make 都可以用来分配空间，初始化类型，但是它们确有不同。
#### new
new(T) 返回的是 T 的指针
new(T) 为一个 T 类型新值分配空间并将此空间初始化为 T 的零值，返回的是新值的地址，
也就是 T 类型的指针 *T，该指针指向 T 的新分配的零值。

#### make
make 只能用于 slice，map，channel 三种类型，make(T, args) 返回的是初始化之后的 T 类型的值，
这个新值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。

#### 示例
```go
package main

import (
	"fmt"
)

func main() {
	newList := new([]int)
	fmt.Printf("new []int : %#v\n", newList)
	var varList []int
	fmt.Printf("var []int : %#v\n",varList)
	makeList := make([]int, 2)
	fmt.Printf("make []int : %#v\n",makeList)

	// output:
	//  new []int : &[]int(nil)
	//  var []int : []int(nil)
	//  make []int : []int{0, 0}
}
```

### sync.WaitGroup
sync.WaitGroup 用于等待一组 goroutine 完成。WaitGroup内部实现了一个计数器，
用来记录未完成的操作个数，它提供了三个方法：
- Add()用来添加计数
- Done()用来在操作结束时调用，使计数减 1。
- Wait()用来等待所有的操作结束，即计数变为0，该函数会在计数不为0时等待，在计数为0时立即返回。

### 方法接收者
- 指针接收者
- 值接收者

因为大部分方法在被调用后都需要维护接收者的值的状态，所以，一个最佳实践是，将方法的接收者声明为指针。

**使用指针作为接收者声明的方法，只能在接口类型的值是一个指针的时 候被调用。使用值作为接收者声明的方法，在接口类型的值为值或者指针时，都可以被调用。**
```go
package main

import "fmt"

type Interface1 interface {
	ValueReceiver()
}

type Implementation1 struct {
}

func (impl Implementation1) ValueReceiver() {
	fmt.Println("value receiver")
}

type Interface2 interface {
	PointReceiver()
}
type Implementation2 struct {
}

func (impl *Implementation2) PointReceiver() {
	fmt.Println("point receiver")
}

func main() {
	var concreteImpl1 Implementation1
	var impl11 Interface1 = concreteImpl1
	impl11.ValueReceiver()

	var impl12 Interface1 = &concreteImpl1
	impl12.ValueReceiver()

	var concreteImpl2 Implementation2

	// cannot do like below
	//var impl21 Interface2 = concreteImpl2

	// can do like below
	var impl21 Interface2 = &concreteImpl2
	impl21.PointReceiver()
}
```

