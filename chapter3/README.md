# Chapter3

## package
所有 Go 语言的程序都会组织成若干组文件，每组文件被称为一个包。这样每个包的代码都可以作为很小的复用单元，被其他项目引用。

所有的.go 文件，除了空行和注释，都应该在第一行声明自己所属的包。每个包都在一个单独的目录里。
不能把多个包放到同一个目录中，也不能把同一个包的文件分拆到多个不同目录中。
这意味着，同一个目录下的所有 .go 文件必须声明同一个包名。

**给包命名的惯例是使用包所在目录的名字。**

一般情况下，包被导入后会使用你的包名作为默认的名字，不过这个导入后的名字可以修改。
这个特性在需要导入不同目录的同名包时很有用。

## package main
所有用 Go 语言编译的可执行程序都必须有一个名叫 main 的包。
当编译器发现某个包的名字为 main 时，它一定也会发现名为 main()的函数，否则不会创建
可执行文件。 main()函数是程序的入口，所以，如果没有这个函数，程序就没有办法开始执行。
程序编译时，会使用声明 main 包的代码所在的目录的目录名作为二进制可执行文件的文件名。

## import package
import 语句告诉编译器到磁盘的哪里去找想要导入的包。
导入包需要使用关键字 import，它会告诉编译器你想引用该位置的包内的代码。
如果需要导入多个包，习惯上是将 import 语句包装在一个导入块中

```go
package test
import (
    "fmt"
    "strings"
)
```

### rename imported package
重名的包可以通过命名导入来导入。命名导入是指，在 import 语句给出的包路径的左侧定义一个名字，将导入的包命名为新名字。

```go
package test
import (
    "fmt"
    myfmt "mylib/fmt"
)

```

编译器会使用 Go 环境变量设置的路径，通过引入的相对路径来查找磁盘上的包。
标准库中 的包会在安装 Go 的位置找到。Go 开发者创建的包会在 GOPATH 环境变量指定的目录里查找。
GOPATH 指定的这些目录就是开发者的个人工作空间。

### Blank Identifier 「_」 （空白标识符）
用户可能需要导入一个包，但是不需要引用这个包的标识符。在这种情况，可以使用空白标识符_来重命名这个导入。

**空白标识符**下划线字符(_)在 Go 语言里称为空白标识符，有很多用法。这个标识符用来抛弃不 想继续使用的值，如给导入的包赋予一个空名字，或者忽略函数返回的你不感兴趣的值。

```go
package test
import (
    "fmt"
    _ "mylib/fmt"   // 不需要使用包里的标识符
)

func TestFunc() {
    _ = OtherFunc() // 忽略 OtherFunc() 的返回

    for _, n := range Array {   // 忽略数组下标
        // do something
    }
}

```

## GOPATH
设置 GOPATH
```bash
export $GOPATH=yourdirectory
```

## go tools

### go build
```bash
go build -o target source.go  # 将 source.go 生成可执行程序 target，source.go 的包名必须为 main 且含有 main() 函数
```

### go run
```bash
go run source.go  # 运行 source.go，source.go 的包名必须为 main 且含有 main() 函数

```

## dependency manage
### dep
**安装 dep**
```bash
go get -u github.com/golang/dep/cmd/dep
```

**dep init**
```bash
dep init
```
在项目目录下执行上述命令。

_目录必须位于 `$GPATH/src` 下_

执行成功后会在项目目录下生成 `Gopkg.toml` 和 `Gopkg.lock`

**dep ensure**
```bash
dep ensure
```
根据 `Gopkg.toml` 和 `Gopkg.lock` 中的数据构建vendor目录和同步里面的包。

## Project directory structure

Project/

|--- pkg/   程序源文件目录

|---|--- package1/

|---|--- package2/

|--- cmd/   程序入口文件目录

|---|--- cmd1/

|---|--- cmd2/