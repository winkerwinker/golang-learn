package main

import "math"
import "fmt"


/*
一个对象可以实现多个接口
一个接口可以被多个对象实现
和java中的概念一样
 */
type geometry interface {
	// 实现了接口的方法就是实现了接口
	area() float64
	perim() float64
}

// 定义长方形
type rect struct {
	width, height float64
}

// 定义圆
type circle struct {
	radius float64
}

// 实现了接口中声明的所有方法即实现了该接口，这里 rects 实现了 geometry 接口
func (r rect) area() float64 {
	return r.width * r.height
}


// 如果没有完全实现所有的接口那么返回
// Cannot use 'r' (multipleInherit rect) as multipleInherit geometry Type does not implement 'geometry' as some methods are missing: perim() float64


func (r rect) perim() float64 {
	return (r.width + r.height) * 2
}

// todo 为什么有些要注明返回
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 1, height: 4}
	c := circle{radius: 1}
	measure(r)
	measure(c)

}
