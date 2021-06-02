package alltype

import "fmt"

/*
常量
*/

func ConstDemo() {

	const hardEight = (1 << 100) >> 97

	fmt.Println("------------")

	// 在中间程度、这个是合法的
	fmt.Println(hardEight)
}
