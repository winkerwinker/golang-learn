package alltype

import (
	"fmt"
)

type File struct {
	fd   int    // file descriptor number
	name string // file name at Open time
}

//这将返回一个指向新File结构体的指针，结构体存有文件描述符和文件名。
//这段代码使用了GO的复合变量(composite literal)的概念，
// todo 复合变量(composite literal) 就是struct吗
//和创建内建的maps和arrays类型变量一样。要创建在堆（heap-allocated）中创建一个新的 对象，我们可以这样写：
func newFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	return &File{fd, name}
}

func newFileDemo(fd int, name string) *File {

	n := new(File)
	n.fd = fd
	n.name = name

	fmt.Print(Stdin.fd)
	return n

}

var (
	Stdin  = newFile(0, "/dev/stdin")
	Stdout = newFile(1, "/dev/stdout")
	Stderr = newFile(2, "/dev/stderr")
)

//func Open(name string, mode int, perm uint32) (file *File, err os.Error) {
//	r, e := syscall.Open(name, mode, perm)
//	if e != 0 {
//		err = os.Errno(e)
//	}
//	return newFile(r, name), err
//}
