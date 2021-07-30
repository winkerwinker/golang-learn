package map_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

/*

Kind为ptr指针种类，Type为float64的指针类型注意

*/

//reflect.TypeOf(variable).Kind()
func TestDo2(t *testing.T) {
	var maps map[string]string
	ofMap := CreateCompositeObjectsOfMap(reflect.TypeOf(maps))
	inter := ofMap.Interface()
	m := inter.(map[string]string)
	fmt.Println(m)
	// map[] 但是也是nil， 过程中reflect value 、 interface 都不是nil ，但是map是nil
	// todo  判断是否是nil 的方式，为什么有了type还是nil 呢？，和空map有什么区别呢
	fmt.Println(inter == nil)
	fmt.Println(m == nil)
	m["ok"] = "ok"

}

// 反射创建map
// todo 会创建盛一个nil
func CreateCompositeObjectsOfMap(t reflect.Type) reflect.Value {
	return reflect.Zero(t)
}

func CreateMap(key, elem reflect.Type) reflect.Value {
	var mapType reflect.Type
	mapType = reflect.MapOf(key, elem)
	return reflect.MakeMap(mapType)
	// 使用interface进行转化
}

func extractMap(v reflect.Value) (interface{}, error) {
	if v.Kind() != reflect.Map {
		return nil, errors.New("invalid input")
	}
	var mapVal interface{}
	mapVal = v.Interface()
	return mapVal, nil
}

// 创建指针对象
//注意到上面的功能同样可以使用 reflect.New(t) 来实现。
func CreatePtr(t reflect.Type) reflect.Value {
	var ptrType reflect.Type
	ptrType = reflect.PtrTo(t)
	return reflect.Zero(ptrType)
}

func extractElement(v reflect.Value) (interface{}, error) {
	if v.Kind() != reflect.Ptr {
		return nil, errors.New("invalid input")
	}
	var elem reflect.Value
	// Elem returns the value that the interface v contains
	// or that the pointer v points to.
	elem = v.Elem()
	//var elem interface{}
	//elem = v.Interface()
	// 返回的是具体的类型吗？

	return elem, nil
}
