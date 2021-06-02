package slot

import (
	"fmt"
	"reflect"
)

/*
结构体穿参
*/

// 接口 的实现 过程中 不是 类和接口 产生关系，
// 方法将他们关联
type Student struct {
	Name string
	Age  int
}

/**
在实际开发中,一般传递的结构体对象都是指针传递 ，如果我把struct直接传递，那么其实并不是java中的传递方式
*/

func addAgePointer(s *Student) {
	fmt.Println(reflect.TypeOf(s))
	// todo 指针可以直接调用方法吗
	s.Age += 2
}

func addAge(s Student) {
	s.Age += 1
}

func SlotStruct() {
	student := Student{Name: "acxie", Age: 11}

	addAge(student)

	fmt.Println(student.Age)
	addAgePointer(&student)
	fmt.Println(student.Age)
}
