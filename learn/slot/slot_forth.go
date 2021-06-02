package slot

import "fmt"

/*
四重指针
*/

func DemoForthSlot() {
	var a int = 10
	ptr := &a
	pptr := &ptr
	ppptr := &pptr
	pppptr := &ppptr

	fmt.Printf("变量 a = %d\n", a)
	fmt.Printf("指针变量 *ptr = %d\n", *ptr)
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
	fmt.Printf("指向指针的指针的指针变量 ***ppptr = %d\n", ***ppptr)
	fmt.Printf("指向指针的指针的指针变量的指针 ****pppptr = %d\n", ****pppptr)
	fmt.Println("------")

}
