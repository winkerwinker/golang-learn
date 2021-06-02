package slot

import "fmt"

/*
入参是指针,参数值可以在函数内部被修改
*/


func SlotFuncCompare() {
	a := 1
	b := 2

	SwapSlot1(&a, &b)
	fmt.Printf("after : %d  ->  %d \n ", a, b)
	// re init

	a = 1
	b = 2
	Swap1(a, b)
	fmt.Printf("after : %d  ->  %d \n ", a, b)

}

// 交换函数这样写更加简洁，也是 Golang 的特性
func SwapSlot1(x *int, y *int) {

	// 交换的是指向
	*x, *y = *y, *x
}

func Swap1(x int, y int) {

	x, y = y, x
}
