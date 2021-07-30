package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func test1() {
	//x := 0
	//y := 0
	// 可以
	//z := x +
	//y
	// 不可以
	//z := x
	//+ y
}

// go fmt 会将代码格式重写
// go import

/*
命令行参数
来自外部的输入
slice 的切片包含第一个元素，但是不包含最后一个元素，如果元素缺失直接不填写
*/
func args() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

/*
// 开头是给程序员看的注释，编译器会忽略
声明一个包的时候，使用注释对其概括
对于main包，注释是一个或多个完整的橘子，用来对这个程序进行完整的概括

var 变量会在声明的时候就进行初始化，如果变量没有明确的初始化，那么他将隐形的初始化为这个类型的空值
eg。 数字是0
字符串是""
？ 对于一个指针个一个类型呢？
*/

func slotTest() {
	var re rect
	fmt.Println(re)
	r := &re
	// 出来是这个 &{0 0}
	fmt.Println(r)
	if &re == nil {
		fmt.Println("是nil")
	} else {
		fmt.Println("不是nil")
	}

	var rect = new(rect)
	rect.width = 100
	rect.height = 200
	//0xc00010a020
	fmt.Println(&rect)
}

// 定义长方形
type rect struct {
	width, height float64
}

func main() {

}

// for 是唯一的循环语句
// initialization、condition、post
// 复制while循环只需要只使用condition
func loopTest() {
	var s, sep string
	for _, arg := range os.Args[1:] {
		// 每一次range 产生一堆值，索引和这个索引处的元素的值，一般配合空标识符_
		s += sep + arg
		sep = " "
	}
	// 代价比较大，一般替代join函数使用
	_ = strings.Join(os.Args[1:], "")
	fmt.Println(s)
}

/*
1. 找到重复行
2. printf的使用
%d 十进制
$x %o %b 16、8、2进制
%s 浮点数
%t 布尔
%S 字符串
%q 带引号的字符串或者字符
%v 内置格式的任何值
%% 百分号本身
转义符号被成为verb
\t 制表符 \n 换行符
*/
func findRepeatLine() {

}

func getUrl() {

	for _, url := range os.Args[1:] {

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintln(os.Stderr, "fetch:%v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		// 避免资源泄漏
		resp.Body.Close()

		if err != nil {
			fmt.Fprintln(os.Stderr, "fetch:reading%s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("%s", b)
	}
}

//并发获取多个url的内容
func fetchUrls() {

	start := time.Now()
	ch := make(chan string, len(os.Args[1:]))

	for _, url := range os.Args[1:] {
		go fetchUrl(url, ch) // 启动一个
	}

	// 只需要次数
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs alapsed\n", time.Since(start).Seconds())

}

func fetchUrl(url string, ch chan string) {

}
