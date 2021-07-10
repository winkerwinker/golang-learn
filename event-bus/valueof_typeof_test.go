package event_bus_test

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"testing"
)

// valueof 和 typeof 的测试

// golang's reflect
// reflect is a kind of application ,they can describe themselves.that means the kind of
// application can use a mechanism to implement self-representation and examination.and
// adjusting  or updating the application's state and semantic by the state and result
// plus： golang'grpc implement by reflect

// 设计原则
// 变量包括 type/value  ->  nil!=nil
// todo 静态类型和具体类型
// type包括 static type(when you code) 和 concrete type(seen by runtime system)

// interface 类型和静态类型
// 在Golang的实现中，每个interface变量都有一个对应pair，pair中记录了实际变量的值和类型
// value是实际变量值，type是实际变量的类型。一个interface{}类型的变量包含了2个指针，一个指针指向值的类型【对应concrete type】，另外一个指针指向实际的值【对应value】。

func Test1(t *testing.T) {
	test1()
}

func test1() {

	tty, _ := os.OpenFile("/dev/tty", os.O_RDWR, 0)

	// 接口变量r的pair中将记录如下信息：(tty, *os.File)，这个pair在接口变量的连续赋值过程中是不变的，将接口变量r赋给另一个接口变量w:

	var r io.Reader
	r = tty

	//*os.File
	//&{0xc0001102a0}
	//*os.File
	//&{0xc0001102a0}
	fmt.Println(reflect.TypeOf(r))
	fmt.Println(reflect.ValueOf(r))
	var w io.Writer
	w = r.(io.Writer)
	fmt.Println(reflect.TypeOf(w))
	fmt.Println(reflect.ValueOf(w))
}

//reflect.TypeOf： 直接给到了我们想要的type类型，如float64、int、各种pointer、struct 等等真实的类型
//reflect.ValueOf：直接给到了我们想要的具体的值，如1.2345这个具体数值，或者类似&{1 "Allen.Wu" 25} 这样的结构体struct的值
//也就是说明反射可以将“接口类型变量”转换为“反射类型对象”，反射类型指的是reflect.Type和reflect.Value这两种
// todo 可以“接口类型变量”转换为“反射类型对象”

//作者：吴德宝AllenWu
//链接：https://juejin.cn/post/6844903559335526407
//来源：掘金
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
