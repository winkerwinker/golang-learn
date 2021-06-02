package main

import (
	"fmt"
	"strings"
)

// 匿名、 组合


/*
继承
一个结构体嵌到另一个结构体，称作组合
匿名和组合的区别

如果一个struct嵌套了另一个匿名结构体，那么这个结构可以直接访问匿名结构体的方法，从而实现继承

如果一个struct嵌套了多个匿名结构体，那么这个结构可以直接访问多个匿名结构体的方法，从而实现多重继承
todo 以下是匿名结构体

如果一个struct嵌套了另一个【有名】的结构体，那么这个模式叫做组合

*/


type Car struct {
	weight int
	name   string
}

func (p *Car) Run() {
	fmt.Println("running")
}

type Bike struct {
	// 可以直接使用内嵌的功能 ？？ ！！！ todo
	Car
	lunzi int
}
type Train struct {
	Car
}

func (p *Train) String() string {
	str := fmt.Sprintf("name=[%s] weight=[%d]", p.name, p.weight)
	return str
}

// todo init 的时候跑出panic ，外面如何处理

func main() {

	var ok bool
	fmt.Println(ok)

	fmt.Println(ChooseStrategy("model"))
	fmt.Println(ChooseStrategy("model--"))
	fmt.Println(nil)

	//var (
	//	strategy Car
	//	ok       bool
	//)

	strategy, ok := StrategyMap["model--"]

		fmt.Println(strategy)
		fmt.Println(ok)

	//{0 }
	//false
	//{88 hello}
	//true
	// todo  不会为nil 会初始化一个默认的值得，所以只能用第二个值来判断




	strategy, ok = StrategyMap["model"]

		fmt.Println(strategy)
		fmt.Println(ok)







	file :="CategoryPage_Global_v3.model"
	// split := strings.Split(file, "_")
	index := strings.LastIndex(file, "_")
	modelCode  := file[:index]
	modelVersion  := strings.Split(file[index+1:], ".")[0]
	fmt.Println(modelCode)
	fmt.Println(modelVersion)



	var a Bike
	a.weight = 100
	a.name = "bike"
	a.lunzi = 2
	fmt.Println(a)
	a.Run()

	var b Train
	b.weight = 100
	b.name = "train"
	b.String()
	b.Run()
	fmt.Printf("%s", &b)
}



func ChooseStrategy(bizCode string) Car {
	// 为什么会给默认值


	return StrategyMap[bizCode]
}

var StrategyMap map[string]Car = map[string]Car{
	"model":   Car{weight: 88,name: "hello"},
	"feature": Car{weight: 88,name: "hello111"},
}


