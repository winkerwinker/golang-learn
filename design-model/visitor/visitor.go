package visitor

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type visitor func(shape Shape)

type Shape interface {
	accept(visitor)
}

type Circle struct {
	Radius int
}

//对圆要做一些操作，
// 我这种类型的东西 接受 你来对我做一些 不同类型的操作
func (c Circle) accept(Visitor visitor) {
	//panic("implement me")
	Visitor(c)
}

type Rectangle struct {
	Width, Heigh int
}

func (r Rectangle) accept(v visitor) {
	v(r)
}


func JsonVisitor(shape Shape) {
	bytes, err := json.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}
func XmlVisitor(shape Shape) {
	bytes, err := xml.Marshal(shape)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))
}


// 将操作和对象完全解耦
func main() {
	c := Circle{10}
	r := Rectangle{100, 200}
	shapes := []Shape{c, r}
	for _, s := range shapes {
		s.accept(JsonVisitor)
		s.accept(XmlVisitor)
	}
}