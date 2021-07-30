package map_test

import (
	"errors"
	"reflect"
	"testing"
)

/*
创建复杂类型的对象
可以像原始对象一样使用零值来创建复合对象。但是，仅使用一个空值而不做其他额外的操作的话，它们并不能被初始化。下面一节将详细讨论如何初始化复合种类。

*/

func TestDo1(t *testing.T) {

}

func CreateCompositeObjects(t reflect.Type) reflect.Value {
	return reflect.Zero(t)
}

func CreateArray(t reflect.Type, length int) reflect.Value {
	var arrayType reflect.Type
	// 使用reflect.ArrayOf创建默认数组
	// 其中类型决定于传进来的参数 t ，数组长度决定于length
	arrayType = reflect.ArrayOf(length, t)
	return reflect.Zero(arrayType)

	// 注意 Slice() 方法也可以用来提取数组的值，但是该方法需要在计算 reflect.Value 之前先将数组转换为一个可寻址的数组。
	// 提取的时候从一个interfcae中提取

}

func extractArray(v reflect.Value) (interface{}, error) {

	// todo 可以转换为slice v.Slice()
	if v.Kind() != reflect.Array {
		return nil, errors.New("invalid input ")

	}

	array := v.Interface()
	return array, nil
}

/*
复合类型

Array
Chan
Func
Interface
Map
Ptr
Slice
Struct
*/

//type signature 来创建复合信道对象
func CreateCompositeChan(t reflect.Type) reflect.Value {
	return reflect.Zero(t)
}

//ChanOf 是用来创建信道的
//type signature，MakeChan(Type, int) 方法可以用来给信道分配内存。
//信道中元素的类型取决于传递进来的参数 t。信道的容量取决于参数 buffer。
//要想使用 reflect 包将值提取进一个信道中，最好的办法是将信道作为接口处理。信道的方向通过传入 ChanOf 的第一个参数来控制，可能的取值有：
func CreateChan(t reflect.Type, buffer int) reflect.Value {
	chanType := reflect.ChanOf(reflect.BothDir, t)
	return reflect.MakeChan(chanType, buffer)
	//使用interfcae 转换为具体的类型
}

// todo 创建函数对象
// 根据f func(args []reflect.Value) (results []reflect.Value), //具体的函数
// 创建 func
func CreateFunc(
	fType reflect.Type, // 函数类型
	f func(args []reflect.Value) (results []reflect.Value), //具体的函数
) (reflect.Value, error) {

	if fType.Kind() != reflect.Func {
		return reflect.Value{}, errors.New("invalid input")
	}
	var ins, outs *[]reflect.Type
	ins = new([]reflect.Type)
	outs = new([]reflect.Type)
	// 返回一个func类型的入参数量
	for i := 0; i < fType.NumIn(); i++ {
		*ins = append(*ins, fType.In(i))
	}
	for i := 0; i < fType.NumOut(); i++ {
		*outs = append(*outs, fType.Out(i))
	}
	var variadic bool
	variadic = fType.IsVariadic()
	//
	return AllocateStackFrame(*ins, *outs, variadic, f), nil

}

func AllocateStackFrame(
	ins []reflect.Type,
	outs []reflect.Type,
	variadic bool,
	f func(args []reflect.Value) (results []reflect.Value),
) reflect.Value {
	var funcType reflect.Type
	// 有了in 、out 还有一个 f，然后产生结果？？？
	funcType = reflect.FuncOf(ins, outs, variadic)
	// 这个f 应该是怎么传过去的
	return reflect.MakeFunc(funcType, f)

	// outs := v.Call(ins)
}

// todo 求场景
