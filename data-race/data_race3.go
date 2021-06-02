package main




// 应该有提到过我们要写出可测性比较高的代码就要少用或者是尽量避免用全局变量，使用 map 作为全局变量比较常见的一种情况就是配置信息。
// 关于全局变量的话一般的做法就是加锁，、
//就本文这个问题也可以使用 sync.Map
//这个下一篇文章会讲，这里篇幅有限就不多讲了
 // todo 使用map作为全局 比较多的可能就是不安全的

// 不受保护的全局变量
var service = map[string]string{}


// RegisterService RegisterService
func RegisterService(name, addr string) {
	service[name] = addr
}

// LookupService LookupService
func LookupService(name string) string {
	return service[name]
}
