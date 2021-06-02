package alltype

/*
申请内存
*/

type T struct{ a, b int }

func applyMemory() {
	// todo new 返回的是指针
	var t *T = new(T)
	t = new(T)

	t = &T{a: 1, b: 1}

	t1 := T{a: 1, b: 1}
	_ = t1
	_ = t

}

/*
maps, slices 和 channels是引用语意（reference semantics）

todo 引用语意
对于这三类引用类型的变量，需要用另一个内建的make()分配并初始化空间：


new(T) 返回的类型是 *T , 而 make(T) 返回的是引用语意的 T
如果你(错误的)使用 new()` 分配了一个引用对象，你将会得到一个指向 nil 引用的指针
这个相当于声明了一个未初始化引用变量并取得 它的地址。
*/
func applyMemory1() {
	m := make(map[string]int)

	var m1 map[string]int

	// todo 如果不用make 直接进行分配 会不会有什么区别
	// 1. make 会分配内存
	// 2. 
	m1 = map[string]int{
		"1": 1,
		"2": 1,
		"3": 1,
	}
	_ = m1
	_ = m

}

func referenceSemantics() {
	//m := make(map[string]int)

}
