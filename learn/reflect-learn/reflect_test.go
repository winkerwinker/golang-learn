package map_test

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

/*
type VersionType string
var Version VersionType

Version 的 type 是 VersionType，
Version 的 kind 是 string。

类型是程序员定义的关于数据和函数的员数据
种类是编译器和运行时定义的关于数据和函数的员数据，
（运行时 runtime,最后是根据kind 来给变量和函数分配内存和栈空间）

*/

//1、 根据versionType来创建Version ，因为version有原始的kind ，所以可以直接使用零值来创建

/*
 初始化的逻辑
1.给定Type
2. Zero获得value
3. value与interface转换

*/
type VersionType string

func TestDo(t *testing.T) {
	var Version VersionType

	objects := CreatePrimitiveObjects(reflect.TypeOf(Version))
	i, err := extractVersionTypeI(objects)
	// 其实是获取了value的类型
	if err == nil {
		fmt.Print(fmt.Sprintf("初始化：%s\n", i))
		return
	}
	fmt.Print(fmt.Sprintf("错误为：%v\n", err))
}

func valueToInterface(value reflect.Value) interface{} {
	//  todo 这里面可能抛出的问题
	return value.Interface()
}

func CreatePrimitiveObjects(t reflect.Type) reflect.Value {
	return reflect.Zero(t)
}

func extractVersionTypeI(v reflect.Value) (string, error) {
	if v.Type().String() != "map_test.VersionType" {
		return "", errors.New("invalid input s")

	}

	// 返回具体的类型
	return v.Type().String(), nil
}

/*
所有的kind
Bool
Int
Int8
Int16
Int32
Int64
Uint
Uint8
Uint16
Uint32
Uint64
Uintptr
Float32
Float64
Complex64
Complex128
String
*/
