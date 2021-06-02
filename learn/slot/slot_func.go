package slot

import "fmt"

/*
参数传递指针函数
todo 注意辨析
*/

func SlotFunc() {
	a := 1
	b := 2

	//ptra := &a
	//ptrb := &b
	fmt.Printf("before: %d  ->  %d \n ",  a, b)
	// todo 如果穿过去不是饮用，那么后期发生的事情就与这里无关
	// SwapSlot(ptra, ptrb)

	SwapSlot(&a, &b)

	fmt.Printf("after : %d  ->  %d \n ", a, b)

}

// 交换指针
// 因为传递过来的是引用，而不是具体的类型
func SwapSlot(a *int, b *int) {
	sentinel := a
	a = b
	b = sentinel
}

func SlotFunc1() {
	var a int = 100
	var b int = 200


	fmt.Printf("交换前 a 的值 : %d\n", a)
	fmt.Printf("交换前 b 的值 : %d\n", b)

	swap1(&a, &b)

	fmt.Println("-------")

	fmt.Printf("交换后 a 的值 : %d\n", a)
	fmt.Printf("交换后 b 的值 : %d\n", b)
}

/* 交换函数这样写更加简洁，也是 Golang 的特性 */
func swap1(x *int, y *int) {

	// 交换的是指向
	*x, *y = *y, *x
}
