package main

 // 结构



//有些语言支持多重继承，
//但是如果
//多个父类存在相同的属性或方法（单继承就不会有）
//就会发生冲突，有些语言为了防止这种情况而只支持单继承，很明显就没有了复用多个父类的属性和方法的优势。
// todo 如果实现的两个抽象类 ，有相同的方法，会发生什么




// 如何给Mother 和 Father 统一一个类呢？
type Father struct {
	MingZi string
}

func (this *Father) Say() string {
	return "大家好，我叫 " + this.MingZi
}

type Mother struct {
	Name string
}

func (this *Mother) Say() string {
	return "Hello, my name is " + this.Name
}


// OO中class的"子继承父"是：子 is_a 父; GO中type的"内嵌式继承"是：子 has_a 父; 父不过是子的一种属性而已。
type Child struct {
	Father
	Mother
}

func main() {
	c := new(Child)
	c.MingZi = "张小明"
	c.Name = "Tom Zhang"
	c.Mother.Say()
}