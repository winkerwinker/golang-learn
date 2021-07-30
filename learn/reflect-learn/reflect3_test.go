package map_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"testing"
)

/*

复合结构体对象

*/

type Rectangle struct {
	Width, Heigh int
	str          string
}

func TestDo3(t *testing.T) {
	r2 := Rectangle{Width: 1, Heigh: 20, str: "test"}
	marshal, _ := json.Marshal(r2)
	fmt.Println(string(marshal))
	fmt.Println("==========")
	rectangle := Rectangle{}
	of := reflect.TypeOf(rectangle)
	value := CreateCompositeStructObjects(of)
	r := value.Interface().(Rectangle)
	fmt.Println(value.Interface())
	fmt.Println(value.Interface() == nil)
	fmt.Println("==========")
	fmt.Println(r)
	fmt.Println(&r == nil)
	json.Unmarshal([]byte("{\"Width\":1,\"Heigh\":20}"), &r)
	fmt.Println("==========")
	fmt.Println(r)
}

func CreateCompositeStructObjects(t reflect.Type) reflect.Value {
	return reflect.Zero(t)
}

func CreateStruct(fields []reflect.StructField) reflect.Value {
	var structType reflect.Type
	structType = reflect.StructOf(fields)
	return reflect.Zero(structType)
}

func extractStruct(v reflect.Value) (interface{}, error) {
	if v.Kind() != reflect.Struct {
		return nil, errors.New("invalid input")
	}
	var st interface{}
	st = v.Interface()
	return st, nil
}
