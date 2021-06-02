package receiver

import "fmt"

/*
value method 可以被 pointer和value 对象调用，而
pointer method 只能被 pointer 对象调用
*/
type MyInterface interface {
	foo()
}

type MyStruct struct {
	ii int64
}

// 因为foo被声明成了pointer receiver，而实际需要的是value receiver。
// func (m *MyStruct) foo() {
// todo 对于go没有基本类型，都是引用的概念，全部都要走引用。所以如果在一个value对象上调用pointer method，编译器会对原来的值做一份拷贝(参考函数传参规范)
// https://www.jianshu.com/p/d1a9bbd0ae36
func (m MyStruct) foo() {
	fmt.Println(m.ii)
	m.ii++
}

func Hello(p MyInterface) {
	p.foo()
}

func main() {
	var a int64 = 10
	// m := MyStruct { 10 }
	m := MyStruct{a}
	Hello(m)

	fmt.Println(m.ii)
}
