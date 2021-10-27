package main

import "fmt"

// go build -gcflags="-m -l" main.go

//go build -gcflags="-m" main.go
//
//go:noinline
func NewCat(name string, age string) *Cat {
	c := new(Cat) // c will excape to heap，./main.go:15:10: new(Cat) escapes to heap
	c.Name = name
	c.Age = age
	return c
}

type Cat struct {
	Name string
	Age  string
}

func main() {
	//NewCat("Tom", "5")
	x := test()
	//这是因为fmt.Println(a ...interface{}),fmt接受的参数是interface{}，这是类型不确定的。 编译期间不能确定参数的具体的类型，逃逸就会产生。
	fmt.Println(*x)
	// todo 编译期间不能确定参数的具体的类型？？ 每一个interface 都会有逃逸产生
	// todo 每一个fmt.println 都会进行逃逸吗
}

// 动态类型逃逸
func test() *int {
	s := 3 //moved to heap: s，进行了内存逃逸
	return &s
}
