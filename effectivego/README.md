# Effective Go 学习笔记

## Godoc
godoc 既是一个程序，又是一个 Web 服务器，它对 Go 的源码进行处理，并提取包中的文档内容。 出现在顶级声明之前，且与该声明之间没有空行的注释，将与该声明一起被提取出来，作为该条目的说明文档。这些注释的类型和风格决定了 godoc 生成的文档质量。

每个包都应包含一段包注释，即放置在包声明子句前的一个块注释。对于包含多个文件的包，包注释只需出现在其中的任一文件中即可。包注释应在整体上对该包进行介绍，并提供包的相关信息。它将出现在 godoc 页面中的最上面，并为紧随其后的内容建立详细的文档。

在包中，任何顶级声明前面的注释都将作为该声明的文档注释。在程序中，每个可导出（首字母大写）的名称都应该有文档注释。

## 命名
包名：应该简洁明了易于理解，包名应该以小写单词来命名，且不应该使用下划线或者驼峰法；包名应该为源码目录的 base name，如 “encoding/base64” 的包名为 “base64”

由于被导入的项总是通过它们的包名来确定，因此包名可以不用包含在变量命名中

长命名并不会使其更具可读性。一份有用的注释说明文档通常比额外的长名更有价值。

```
Getters：不要将 Get 放在 getter’s name 中
Setters：Setter’s name 可以包含 Set
```

## 接口命名：
按照约定，只包含一个方法的接口应当以该方法的名称加上 - er 后缀来命名，如 Reader、Writer、 Formatter、CloseNotifier 等。

Read、Write、 Close、Flush、 String 等都具有典型的签名和意义。为避免冲突，请不要用这些名称为你的方法命名， 除非你明确知道你的方法和它们的签名和意义相同。。反之，若你的类型实现了的方法， 与一个众所周知的类型的方法拥有相同的含义，那就使用相同的命名。请将字符串转换方法命名为 String 而非 ToString。

## 变量
在满足下列条件时，已经被声明的变量 v 可以出现在 := 声明中
- 本次声明与已声明的 v 处于同一作用域中（若 v 已在外层作用域中声明过，则此次声明会创建一个新的变量）
- 在初始化中与其类型相同的值才能赋值给 v，且在此次声明中至少有一个变量是新声明的

即便 go 中的函数形参和返回值在词法上处于大括号之外，但它们的作用域和该函数体仍然相同

```go
Switch Type
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
   fmt.Printf("unexpected type %T", t) // %T prints whatever type t has
case bool:
   fmt.Printf("boolean %t\n", t) // t has type bool
case int:
   fmt.Printf("integer %d\n", t) // t has type int
case *bool:
   fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
   fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}
```

## Functions
可命名结果形参
Go 函数的返回值或结果 “形参” 可被命名，并作为常规变量使用，就像传入的形参一样。 命名后，一旦该函数开始执行，它们就会被初始化为与其类型相应的零值； 若该函数执行了一条不带实参的 return 语句，则结果形参的当前值将被返回。
此名称不是强制性的，但它们能使代码更加简短清晰：它们就是文档。

## Defer
Go 的 defer 语句用于预设一个函数调用（即推迟执行函数）， 该函数会在执行 defer 的函数返回之前立即执行。它是特殊的能够有效处理例如：资源释放，典型的例子就是解锁互斥和关闭文件。

推迟诸如 Close 之类的函数调用有两点好处：
第一，它能确保你不会忘记关闭文件。
第二，它意味着 “关闭” 离 “打开” 很近，这总比将它放在函数结尾处要清晰明了。

被推迟函数的实参（如果该函数为方法则还包括接收者）在推迟执行时就会求值， 而不是在调用执行时才求值。这样不仅无需担心变量值在函数执行时被改变， 同时还意味着单个已推迟的调用可推迟多个函数的执行。
```go
func DeferTest() {
   fmt.Println("defer Test begin...")
   for i := 0; i < 5; i++ {
      defer fmt.Printf("%d ", i)
   }

   fmt.Println("defer Test end...")
   return
}
```

被推迟的函数按照后进先出（LIFO）的顺序执行，上述函数输出：
```
defer Test begin...
defer Test end...
4 3 2 1 0
```

## New 和 Make
- 内置函数 new 分配空间。传递给new 函数的是一个类型，不是一个值。返回值是 指向这个新分配的零值的指针。
- 内建函数 make 分配并且初始化 一个 slice, 或者 map 或者 chan 对象。 并且只能是这三种对象。 和 new 一样，第一个参数是 类型，不是一个值。 但是make 的返回值就是这个类型（即使一个引用类型），而不是指针。 具体的返回值，依赖具体传入的类型。

```go
var p *[]int = new([]int) // allocates slice structure; *p == nil; rarely useful
var v []int = make([]int, 100) // the slice v now refers to a new array of 100 ints
// Unnecessarily complex:
var p *[]int = new([]int)
*p = make([]int, 100, 100)
// Idiomatic:
v := make([]int, 100)
```

## 数组
在 Go 中：
- 数组是值。将一个数组赋予另一个数组会复制其所有元素。
- 若将某个数组传入某个函数，它将接收到该数组的一份副本而非指针。
- 数组的大小是其类型的一部分。类型 [10]int 和 [20]int 是不同的。

## Slice
Slice 是引用类型，存在内存的重新分配

## Map
映射是方便而强大的内建数据结构，它可以关联不同类型的值。其键可以是任何相等性操作符支持的类型， 如整数、浮点数、复数、字符串、指针、接口（只要其动态类型支持相等性判断）、结构以及数组。 切片不能用作映射键，因为它们的相等性还未定义。与切片一样，映射也是引用类型。 若将映射传入函数中，并更改了该映射的内容，则此修改对调用者同样可见。

## 方法
以指针或值为接收者的区别在于：值方法可通过指针和值调用， 而指针方法只能通过指针来调用。之所以会有这条规则是因为指针方法可以修改接收者；通过值调用它们会导致方法接收到该值的副本，因此任何修改都将被丢弃，语言不允许这种错误。（通过变量并不一定总能取到相应的指针，而通过指针总能取到相应的值）

## 接口
Go 中的接口为指定对象的行为提供了一种方法：如果某样东西可以完成这个， 那么它就可以用在这里。
如果一个动物走路像鸭子，叫得像鸭子，name 它就是鸭子。

## 类型断言
```
switch - case : switch v := value.(type)
v, ok := value.(type)
```

## 空白标识符
检查某个类型是否满足某个接口：
```
var _ Interface = (*Type)(nil)，如：
var _ json.Marshaler = (*RawMessage)(nil)
```

## 内嵌
### 接口内嵌
```go
type Reader interface {
   Read(p []byte) (n int, err error)
}
type Writer interface {
   Write(p []byte) (n int, err error)
}
// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
   Reader
   Writer
}
```

### struct 内嵌
```go
type A struct {

}

type B struct {

}

type AB struct {
   A // 匿名成员；a A 命名成员
   B // 匿名成员; b B 命名成员
}
```

## 并发
通过通信共享内存。go channel 是并发安全的。

```go
func Serve(queue chan *Request) {
   for req := range queue {
      req := req // 为该 Go 程创建 req 的新实例。
      sem <- 1
      go func() {
         process(req)
         <-sem
      }()
   }
}
```
req := req 在 Go 中这样做是合法且惯用的。你用相同的名字获得了该变量的一个新的版本， 以此来局部地刻意屏蔽循环变量，使它对每个协程保持唯一。（闭包也需要这么做）

## panic & recover
实际的库函数应避免 panic。若问题可以被屏蔽或解决， 最好就是让程序继续运行而不是终止整个程序。一个可能的反例就是初始化：若某个库真的不能让自己工 作，且有足够理由产生 panic，那就由它去吧

当 panic 被调用后（包括不明确的运行时错误，例如切片检索越界或类型断言失败）， 程序将立刻终止当前函数的执行，并开始回溯 Go 协程的栈，运行任何被推迟的函数。若回溯到达 Go 协程栈的顶端，程序就会终止。不过我们可以用内建的 recover 函数来重新或来取回 Go 协程的控制权限并使其恢复正常执行。

调用 recover 将停止回溯过程，并返回传入 panic 的实参。 由于在回溯时只有被推迟函数中的代码在运行，因此 recover 只能在被推迟的函数中才有效。