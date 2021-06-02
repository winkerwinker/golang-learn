package chanel

import (
	"testing"
)

//用于超时处理
//“不要通过共享内存来通信，而应该通过通信来共享内存” 这是一句风靡golang社区的经典语
//Go从语言层面保证同一个时间只有一个goroutine能够访问channel里面的。
//todo 一个时间假如有两个gorutine 访问两个chanel会怎么样

//4.阻塞： 发送数据：chan <- data,阻塞的，直到另一条goroutine，读取数据来解除阻塞 读取数据：data <- chan,也是阻塞的。直到另一条goroutine，写出数据解除阻塞。
//5.本身channel就是同步的，意味着同一时间，只能有一条goroutine来操作。

// 阻塞： 发送的过程是阻塞的，只有等一头读取，才能解除阻塞
// chanel本身就是同步的，以为同一时间，
// todo  高亮！只能有一条goroutine 来操作
func TestChanel6(t *testing.T) {
	// c1 := make(chan)
	// c1 <- time.After(time.Second * 1)
	// todo 如何定义一个时间的chanel

}
