package main

import (
	"fmt"
	"time"
)

// 多重继承测试
func main() {
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}

	for _, file := range balance {
		//  todo 如何单独recover 一个 一个风险
		//  Possible resource leak, 'defer' is called in a 'for' loop
		fmt.Println(file)
		fmt.Println("----")
		// 依赖于外部变量
		//如果循环体中会启动协程（并且协程会使用循环变量），就需要格外注意了，因为很可能循环结束后协程才开始执行，
		//此时，所有协程使用的循环变量有可能已被改写。（是否会改写取决于引用循环变量的方式）


		go func(t float32) {
			fmt.Println(t)
		}(file)
	}
	time.Sleep((30) * time.Second)

}
