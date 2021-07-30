package main

/*
#cgo CFLAGS: -I./clibs/demo
#cgo LDFLAGS: -L./clibs/demo -ldemo
#include "demo.h" //非标准c头文件，所以用引号
*/
import "C"

func main() {
	C.demo()
}
