package useful_go

import (
	"testing"
)

//为整型常量值增加一个 String () 方法
//如果你使用 iota 的自定义整型用于枚举，
//请始终增加一个 String () 方法。
//假设你写了这样的代码：

// todo  让枚举从 iota + 1开始自增，不然会给一个默认值
// todo 或者引入unknow 枚举
type State int

// 状态 枚举
const (
	Running State = iota
	Stopped
	Rebooting
	Terminated
)

func (s State) String() string {
	switch s {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	case Rebooting:
		return "Rebooting"
	case Terminated:
		return "Terminated"
	default:
		return "Unknown"
	}
}

//=== RUN   Test_5
//0
//--- PASS: Test_5 (0.00s)
//PASS
// 返回0 没有逻辑上的意义
func Test_5(t *testing.T) {
	//running := Running
	//
	//fmt.Println(running)

	//for i := 0; i < 1000; i++ {
	//	i:=i
	//	go func(index int) {
	//		fmt.Println(index)
	//	}(i)
	//
	//}

	// todo slice是否设置默认大小的区别
	//now := time.Now()
	//result := make([]float64, 10000)
	//
	//for i := 0; i < 10000; i++ {
	//	result[i] = rand.Float64()
	//}
	//
	//logger.Logger.Info(time.Since(now))
	//
	////time="2021-06-21T20:16:53+08:00" level=info msg="259.798µs"
	////time="2021-06-21T20:16:53+08:00" level=info msg="378.656µs"
	//now = time.Now()
	//var result1 []float64
	//
	//for i := 0; i < 10000; i++ {
	//	result1 = append(result1, rand.Float64())
	//}
	//
	//logger.Logger.Info(time.Since(now))

	// todo map是否设置默认大小的区别
	// todo defer
}
