package slot

import (
	"fmt"
	"reflect"
)

/*
空指针学习
*/

func DemoNil1() {
	var ptr *int
	// todo Println 没有办法使用format的方式

	fmt.Printf("ptr 的值为 : %x\n", ptr)

	fmt.Println(" -------   ")

	fmt.Println(reflect.TypeOf(ptr))

	if ptr != nil {
		fmt.Printf("ptr 不是空指针 : %x\n", ptr)
	} else {
		fmt.Printf("ptr 是空指针 : %x\n", ptr)
	}

}

// todo 一个包下不能有重复的方法
//func Demo1() {
//}
