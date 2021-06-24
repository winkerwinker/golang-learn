package concurent_graph_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//这里我们用一个通道代表乒乓球台。一个整型变量代表球，然后用两个
//goroutine 代表玩家，玩家通过增加整型变量的值（点击计数器）模拟击球动作。

// 适用于传递消息

//我们可以看到每个玩家都按照次序轮流操作，你可能会想为什么会这样。为什么多个玩家（goroutine）会按照严格的顺序接到 “球” 呢。
//
//答案是 Go 运行时环境维护了一个 接收者 FIFO 队列 (存储需要从某一通道上接收数据的 goroutine)，在我们的例子里，每个玩家在刚发出球后就做好了接球准备。我们来看一下更复杂的情况，加入 100 个玩家。

//————————————————
//原文作者：Summer
//转自链接：https://learnku.com/go/t/39490
//版权声明：著作权归作者所有。商业转载请联系作者获得授权，非商业转载请保留以上作者信息和原文链接。

//   todo 对于chan 来说 等待的协程是顺序消费的

func Test_3(t *testing.T) {

	num := 42

	go1 := make(chan int)
	go2 := make(chan int)
	// all goroutines are asleep - deadlock!
	//go1 <- 1

	go func() {
		for i := 0; i < 10; i++ {
			if i != 0 {
				<-go1
			}

			num++
			fmt.Println("go1：" + strconv.Itoa(num))
			// 可以控制让第二个程序加什么，加多少
			go2 <- num
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-go2
			num--
			fmt.Println("go2：" + strconv.Itoa(num))
			go1 <- num
		}
	}()

	time.Sleep(10 * time.Second)

}

func Test_00(t *testing.T) {

	go1 := make(chan int)

	//all goroutines are asleep - deadlock!
	//  todo 在声明chanel 前，先往里面写东西会死锁
	go1 <- 1

	fmt.Printf("-----")
	fmt.Printf("-----")
	fmt.Printf("-----")
	fmt.Printf("-----")
}
