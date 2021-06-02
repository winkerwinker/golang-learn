package error_deal

import "log"

// 处理error

//1. 如果强依赖服务故障  ->  panic
//2. 不是不可恢复的错误  ->  error
//3. 在程序入口 例如， gin 中间件，需要使用recover 预防panic
//4. 避免使用也正go routine
//	1. 如果是在请求中需要执行异步任务，使用异步worker
//		消息通知的方式处理，避免请求量大时候大量创建goroutine
//	2. 如果需要使用goroutine ，应使用统一的GO函数创建，函数中会进行recover
//		避免因为野生goroutine panic 导致主进程推出



// 1. 全局panic 回复
// 2. 协程panic 回复
func Go(f func()){
	go func(){
		defer func(){
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
			}
		}()

		f()
	}()
}