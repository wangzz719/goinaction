package main

import "fmt"

type Interface interface {
	ReceiverFunc()
}

type UserType1 struct {
	name string
}

func (ut UserType1) ReceiverFunc() {
	fmt.Println(ut.name)
}

type UserType2 struct {
	name string
}

func (ut *UserType2) ReceiverFunc() {
	fmt.Println(ut.name)
}

func main() {
	var it1 Interface = UserType1{name: "wzz"}
	it1.ReceiverFunc()

	var itp1 Interface = &UserType1{name: "wzz"}
	itp1.ReceiverFunc()

	// 下面赋值语句会报类型错误
	// cannot use UserType2 literal (type UserType2) as type Interface in assignment
	// UserType2 does not implement Interface (ReceiverFunc method has pointer receiver)
	var it2 Interface = UserType2{name: "wzz"}
	it2.ReceiverFunc()

	// 下面语句可以正确执行
	var itp2 Interface = &UserType2{name: "wzz"}
	itp2.ReceiverFunc()
}
