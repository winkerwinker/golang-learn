package word

/*
主要作用都是返回指针
*/
func makeDemo() {

	//slice := make([]int, 0, 100)
	//hash := make(map[int]bool, 10)
	//ch := make(chan int, 5)

}

// 两种方式是等价的
func newDemo() {

	a := new(int)

	var v int
	b := &v

	_ = a
	_ = b
}
