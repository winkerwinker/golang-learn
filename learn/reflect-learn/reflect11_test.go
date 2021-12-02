package map_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

/*
reflect
*/

func TestDo11(t *testing.T) {

	m := map[string]int{"11": 1, "22": 2}
	fmt.Println(m)
	fmt.Printf("%v", m)
	fmt.Printf("================%v", m)
	p := &Person{Name: "nihao", Age: 11}

	for i := 0; i < 10000; i++ {
		i := i
		go func() {
			toStruct := SetValueToStruct(p, fmt.Sprintf("zhang%d", i), 18)
			fmt.Println(toStruct)
			toStruct1 := SetValueToStruct1(p, fmt.Sprintf("zhang%d", i), i)
			fmt.Println(toStruct1)
		}()
	}

	time.Sleep(1 * time.Minute)
	//fmt.Println(*p)
	//fmt.Println(p.Name)
	//fmt.Println(p.Age)
}

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func SetValueToStruct(p *Person, name string, age int) *Person {
	// 必须是地址结构体

	v := reflect.ValueOf(p).Elem()

	// value := reflect.Zero(v.FieldByName("Name").Type())

	v.FieldByName("Name").Set(reflect.ValueOf(name))

	// reflect: call of reflect.Value.Set on zero Value [recovered]
	//v.FieldByName("Name").
	// json.Marshal(p)
	//v.FieldByName("Name").Set(reflect.ValueOf(name))
	//v.FieldByName("Age").Set(reflect.ValueOf(age))
	return p
}

func SetValueToStruct1(p *Person, name string, age int) *Person {
	// 必须是地址结构体

	v := reflect.ValueOf(p).Elem()

	// value := reflect.Zero(v.FieldByName("Name").Type())

	v.FieldByName("Age").Set(reflect.ValueOf(age))

	// reflect: call of reflect.Value.Set on zero Value [recovered]
	//v.FieldByName("Name").
	// json.Marshal(p)
	//v.FieldByName("Name").Set(reflect.ValueOf(name))
	//v.FieldByName("Age").Set(reflect.ValueOf(age))
	return p
}

//
//1.结构体首字母必须大写，否则会出现：panic: reflect: reflect.Value.Set using value obtained using unexported field
//在Golang中首字母的大小写代表着访问权限，首字母小写则包外无法访问

//2.   反射需要使用地址 否则会出现：panic: reflect: reflect.Value.Set using unaddressable value
//如下代码就会出现panic
//p := Person{}
//v := reflect.ValueOf(p)
//v.FieldByName("Name").Set(reflect.ValueOf(name))
//v.FieldByName("Age").Set(reflect.ValueOf(age))
