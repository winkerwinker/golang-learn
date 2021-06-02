package alltype

import (
	"fmt"
	"unsafe"
)

/*

1 ->  范围：
int8: 2 ^ (8 - 1) ~ 2 ^ (8 - 1) - 1
uint8: 0~ 2 ^ 8

2  ->
int 64位操作系统默认为int64,32位操作系统为int32,但是类型检查时时int
	int8 (byte 1字节)
	int16 (short 2字节)
	int32 (int 4字节)
	int64 (long 8字节)

3  ->
无符号的表示范围更大，因为表示数值得位多1
uint uint8 uint16 uint32 uint64 代表无符号int 无符号只能表示正数

4  -> 自动推导
a := 2
这里a是int, 在64位操作系统中就是int64

*/

func SizeOfInt() {

	var n1 int32 = 666
	var n2 int64 = 666
	// 默认就会是64
	n3 := 666
	fmt.Printf("n1 multipleInherit is %T, 占用字节数:%d \n", n1, unsafe.Sizeof(n1))
	fmt.Printf("n2 multipleInherit is %T, 占用字节数:%d \n", n2, unsafe.Sizeof(n2))
	fmt.Printf("n3 multipleInherit is %T, 占用字节数:%d \n", n2, unsafe.Sizeof(n3))
}

/*
1 ->  浮点数都是有符号的
float32 (4字节 float)
float64 (8字节 double)

2 ->
浮点数=符号位+指数位+尾数位，尾数位可能丢失，造成精度损失
默认声明为float64

3 ->
4.223E2 = 4.223* 10的2次方
*/
func SizeOfFLoat() {

}

/*
字符型

golang没有专门的char类型，一般用单个byte保存单个字母字符
1byte=8bit 一般说是8位

UTF-8编码中：1个英文字符或英文标点占1byte
1个中文汉字或一个中文标点占3byte


# command-line-arguments
.\main.go:12:16: constant 29399 overflows byte
字符'狗'unicode码值是29399,byte的范围-128~127，所以overflows byte

golang中的字符本质是一个整数，可以进行运算,如'A'+23
*/
func SizeOfChar() {
	// var c3 byte = '狗'
	// fmt.Println(c3)

	var c4  int = '狗'
	//  也可使用自动推导 c5 := '狗'
	// Println(cc) 输出的是码值 29399 想输出字符需要

	fmt.Printf("%c",c4)
	fmt.Printf("%c,%d byte,multipleInherit is %T",c4, unsafe.Sizeof(c4), c4)
}


func SizeOfStr() {
	var Message = `I can not love \n golang\t any more!`
	// 反引号中的字符串是原始字符串，就像python中的 r" something "
	var name string ="Jack"
	msg :="hello"+"world"
	msg+="!!!!"
	info := "asjdkljflskdf"+"adsfasdf"+
		"xxxx"

	fmt.Printf(Message)
	fmt.Printf(name)
	fmt.Printf(info)
}





/*
golang中的数据类型转换必须显示转换，没有自动转换
*/


// fmt.Sprintf("%参数", 表达式) // 字符串格式化
// %c  char
// %t  bool
// %d  十进制整数 int、double
// %f  浮点数 float