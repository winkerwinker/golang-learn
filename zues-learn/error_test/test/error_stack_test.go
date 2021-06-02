package test

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
)

// 测试打印堆栈
func TestErrorStack(t *testing.T) {
	level3()
}

// strconv.ParseInt: parsing "abcd": invalid syntax

//错误发生的 func: strconv.ParseInt
//发生错误 func 的参数: abcd
//错误的原因: invalid syntax

// 可以粗略知道方法，参数，原因

//内部定义了 A NumError records a failed conversion.
//需要自定义error给出运行建议，统一管理，也需要堆栈跟踪


func level3() {
	// 使用第三方库

}

func level1() {
	// todo 这个输出可以输出到真的出了问题到地方吗，可以打log的
	// 给我几个error 案例
	if _, err := strconv.ParseInt("abcd", 10, 64); err != nil {
		fmt.Println(err)
	}
}

func level2() {
	err := f()
	fmt.Print(err)
}

// NewTrace creates a simple traceable error.
func NewTrace(message string) error {
	return &trace{m: message, s: callers()}
}

func f() error {
	return NewTrace("ooops")
}

type trace struct {
	m string
	s []uintptr
}

func (e *trace) Error() string {

	var b strings.Builder
	b.WriteString(e.m)
	b.WriteString("\n\n")
	b.WriteString("Traceback:")

	// 根据保存的trace信息
	for _, pc := range e.s {

		fn := runtime.FuncForPC(pc)
		b.WriteString("\n")
		file, n := fn.FileLine(pc)
		b.WriteString(fmt.Sprintf("%s:%d", file, n))
	}
	return b.String()
}

func callers() []uintptr {
	var pcs [32]uintptr
	// 跳过几层caller
	n := runtime.Callers(3, pcs[:])
	st := pcs[0:n]
	return st
}
