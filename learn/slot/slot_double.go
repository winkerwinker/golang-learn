package slot

import "fmt"

/*
双重指针  -> 指向指针的指针
*/

func DemoDoubleSlot() {
	var a int = 10
	ptr := &a
	pptr := &ptr

	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

	fmt.Println("------")

	fmt.Printf("指针变量 *ptr = %d\n", ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", pptr)
}
