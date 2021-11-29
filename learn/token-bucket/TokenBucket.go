package slot

import (
	"fmt"
	"os"
	"reflect"
)

// 将位置再度赋值给一个对象，然后观察他的实际对应
// *int 指的是 指向一个int变量的指针
// todo 同一个目录下面不能有个多 package main
// 它直接指向了这个值的内存地址，并且只占用了4个或者8个字节。

// &a
// *addr，可以作为一个函数，也可以作为一种定义
func Demo1() {
	a := 3
	// var intP *int = &a
	intP := &a
	fmt.Fprintln(os.Stderr, "实际的数据为：", a)
	fmt.Fprintln(os.Stderr, "在内存中的地址为：", intP)
	fmt.Println(reflect.TypeOf(intP))
	fmt.Fprintln(os.Stderr, "实际的数据为：", *intP)
}

func Demo2() {
	a := 3
	fmt.Fprintln(os.Stderr, "实际的数据为：", a)
	fmt.Fprintln(os.Stderr, "在内存中的地址为：", &a)
}
