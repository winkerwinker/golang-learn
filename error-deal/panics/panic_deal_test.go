package panics

import (
	"fmt"
	"testing"
)

// panic
// recover  捕获异常，让程序继续正常运行。类似 java 等语言中的 try ... catch 。
// 正常逻辑下的函数调用栈，是逐个回溯的，而异常捕获可以理解为：
//程序调用栈的长距离跳转。这点在 C 语言里，是通过 setjump 和 longjump 这两个函数来实现的。

// 异常处理直接将调用栈回溯到函数到起点

/*
try catch 、 recover 、setjump 等机制会将程序当前状态
（主要是 cpu 的栈指针寄存器 sp 和程序计数器 pc ，
Go 的 recover 是依赖

// todo  回到调用到入口，然后继续执行下一句。
// 会不会有panic 循环， 因为错过执行某一句，而又导致下一句的panic
// 所以在panic 要做扫尾工作
// 1. 释放资源、
// 2. 删除默认值


todo defer 来维护 sp 和 pc ）
保存到一个与 throw、panic、longjump共享的内存里。当有异常的时候，
从该内存中提取之前保存的 sp 和 pc 寄存器值，直接将函数栈调回到 sp 指向的位置，
并执行 ip 寄存器指向的下一条指令，将程序从异常状态中恢复到正常状态
*/
// recover 是相对一个协程而言的
func Test_1(t *testing.T) {

	foo1()
	fmt.Println("----main---")
}

func foo1() {

	foo()
	fmt.Println("----foo1---")
}

func foo() {
	// todo  recover 完毕会跳到他的上一个调用层级
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("----recover----")
			fmt.Println(err)
		}
	}()
	var bar = []int{1}
	fmt.Println(bar[1])

	fmt.Println("-------")
	fmt.Println("-------")
	fmt.Println("-------")
	fmt.Println("-------")
}

//var zid = new(string)
////*zid = GenZid()
//*zid = traceID
//
//st := time.Now()
//
//response = &recommend.Response{
//GoodsIDs: []int64{},
//AlgoInfo: zid,
//}
