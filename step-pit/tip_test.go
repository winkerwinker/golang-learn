package tips

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	sList := []string{"1", "2", "3"}

	for i := 0; i < len(sList); i++ {

		fmt.Printf("%p \n", &sList[i])

	}
	/*	0xc000100300
		0xc000100310
		0xc000100320
		0xc00011e340
		1
		0xc00011e340
		2
		0xc00011e340
		3*/
	for _, v := range sList {
		// %p是 指针
		// v的地址
		fmt.Printf("%p \n", &v)
		fmt.Printf("%s \n", v)
	}
}

/*
=== RUN   Test1
0xc0000a04f0
1
0xc0000a04f0
2
0xc0000a04f0
3 */

type student struct {
	Name string
	Age  int
}

func Test1_1(t *testing.T) {
	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		// todo 在range里面修改变量都是无法修改成功的
		m[stu.Name] = &stu
	}
	fmt.Println(m)
}

/*

=== RUN   Test1_2
0xc00013e100 0xc00013e100
0xc000120028 0xc000120030
123 456
--- PASS: Test1_2 (0.00s)
PASS

todo 地址的地址块为什么反而不同了
*/
func Test1_2(t *testing.T) {

	ch1 := make(chan int, 3)
	// tood 赋值给引用
	// https://www.huaweicloud.com/articles/a4da83fd0722655d896e04160aa0f96c.html
	ch2 := ch1
	ch1 <- 123
	ch2 <- 456
	// 相当于给一个chan 先进去 123 再进去 456
	// 所以先出来 123 ，再出来 456 ？？/
	fmt.Printf("%p %p \n ", ch1, ch2)

	//
	//824635023616 824635023616
	// todo 是不同的地址块
	//0xc000120028 0xc000120030
	fmt.Printf("%p %p \n ", &ch1, &ch2)
	fmt.Printf("%d %d \n ", <-ch2, <-ch2)

}

// 测试 slice
func Test1_3(t *testing.T) {

	s := make([]int, 3, 100)
	s1 := s
	s2 := s
	s[0] = 123
	s1[1] = 456
	s2[2] = 789
	fmt.Println("原始切片：", s)
	fmt.Println("赋值切片：", s1)
	fmt.Println("赋值切片：", s2)
	fmt.Println("切片的长度、容量为：", len(s), cap(s))
	fmt.Printf("追加之前的地址：s:%p ;s1:%p; s2:%p\n", s, s1, s2)
	fmt.Println("====执行append操作====")
	// 是当容量不够所以拷贝到了新地址，
	s2 = append(s2, 123)
	// todo 重点
	fmt.Printf("追加之后的地址：s:%p ;s1:%p; s2:%p\n", s, s1, s2)

	//追加之前的地址：s:0xc0000ec000 ;s1:0xc0000ec000; s2:0xc0000ec000
	//====执行append操作====
	//追加之后的地址：s:0xc0000ec000 ;s1:0xc0000ec000; s2:0xc0000a8120

	fmt.Printf("原始切片s：%p\n", s)
	fmt.Printf("赋值切片s1：%p\n", s1)
	fmt.Printf("赋值切片s2：%p\n", s2)

	// !!!todo有问题待解决
	//https://www.huaweicloud.com/articles/a4da83fd0722655d896e04160aa0f96c.html
	// todo 同一个内存，但是围栏到的最后的index不同
	//原始切片s： 3
	//赋值切片s1： 3
	//赋值切片s2： 4
	fmt.Println("原始切片s：", len(s))
	fmt.Println("赋值切片s1：", len(s1))
	fmt.Println("赋值切片s2：", len(s2))
}

// todo map 中如果没有对应的元素，那么会自动返回类型对应的默认值的
func Test1_5(t *testing.T) {
	//Go 则会返回元素对应数据类型的零值，比如 nil、'' 、false 和 0，取值操作总有值返回，故不能通过取出来的值来判断 key 是不是在 map 中。
	//检查 key 是否存在可以用 map 直接访问，检查返回的第二个参数即可：

}

// 更新字符串
func Test1_6(t *testing.T) {

	/*	x := "text"
		x[0] = "T"		// error: cannot assign to x[0]
		fmt.Println(x)

		x := "text"
		xBytes := []byte(x)
		xBytes[0] = 'T'	// 注意此时的 T 是 rune 类型
		x = string(xBytes)
		fmt.Println(x)	// Text*/

	/*
		因为一个 UTF8 编码的字符可能会占多个字节，比如汉字就需要 3~4 个字节来存储，此时更新其中的一个字节是错误的。
		更新字串的正确姿势：将 string 转为 rune slice（此时 1 个 rune 可能占多个 byte），直接更新 rune 中的字符
	*/

	x := "text"
	xRunes := []rune(x)
	xRunes[0] = '我'
	x = string(xRunes)
	fmt.Println(x) // 我ext
}

/*
interface 的坑
https://codery.cn/post/golang/interface/
这里我去看了一下interface的底层代码，interface 并不是一个指针类型，它是一个结构体。众所周知，interface 和 空interface 是两种不同的数据结构。分别是：

type eface struct {
    _type *_type
    data  unsafe.Pointer
}

type iface struct {
    tab  *itab
    data unsafe.Pointer
}
todo eface 就是空 interface，不包含任何方法的空接口，也称为 empty interface。
iface 就是正常的 interface，包含方法的接口。

*/
