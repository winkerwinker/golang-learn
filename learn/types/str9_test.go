package cache_sdk

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_create9(t *testing.T) {
	//randsArr("float64", 7)
	strings := make([]string, 0, 10)
	string := make([]string, 10)
	fmt.Println(strings)
	fmt.Println(string)
	//indexs := []string{"ok", "hello", "bye"}
	//
	//for index, str := range indexs {
	//	str := str
	//	index := index
	//	go func() {
	//		fmt.Println(strconv.Itoa(index) + str)
	//	}()
	//}
	//time.Sleep(10 * time.Second)
}

func randsArr(theType string, length int) interface{} {
	var t reflect.Type
	switch theType {
	case "int":
		t = reflect.TypeOf(int(2))
	case "int64":
		t = reflect.TypeOf(int64(3))
	case "float32":
		t = reflect.TypeOf(float32(3))
	case "float64":
		t = reflect.TypeOf(float64(3))
	default:
		return nil
	}
	t = reflect.ArrayOf(length, t)

	of := reflect.ArrayOf(1, t)
	result := reflect.Zero(of)
	fmt.Println(result)

	// interface {} is [1][7]float64, not [][]float64 [recovered]

	i := result.Interface().([1][7]float64)
	fmt.Println(reflect.TypeOf(result.Interface()))
	fmt.Println(reflect.TypeOf(i))

	return result
}
