package alltype

// 有两种构建数组的方式
// [...]T 这种初始化方式也只是 Go 语言为我们提供的一种语法糖
//

func createArray() {
	arr1 := [3]int{1, 2, 3}
	// 编译期推导size
	arr2 := [...]int{1, 2, 3}
	_ = arr1
	_ = arr2
}

// 创建切片的三种方式
//
// 通过下标的方式获得数组或者切片的一部分；
// 使用字面量初始化新的切片；
// 使用关键字 make 创建切片：
func createSlice() {
	arr := [3]int{1, 2, 3}
	slice := []int{1, 2, 3}

	slice1 := arr[0:3]
	slice1 = slice[0:3]
	slice1 = make([]int, 10)
	_ = slice1
	_ = slice

}
