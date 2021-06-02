package context

import (
	"flag"
	"fmt"
	"sync"
	"testing"
	"time"
)

var (
	spendTime = make([]int64, 10000)
)

var (
	total     = *flag.Int("totalNUm", 30, "总请求次数")
	concurNum int //flag.Int("concurNum", 10000, "并发总数")

	rTNum  time.Duration // 响应时间
	rTTNum time.Duration // 平均响应时间

	successNum int // 成功次数
	failNum    int // 失败次数

	beginTime time.Time // 开始时间
	endTime   time.Time // 开始时间
	secNum    int       // 秒数

	uid string
)

// 解析压测
func TestParseStress(t *testing.T) {

	//  起协程 预先准备好一部分数据
	fmt.Scanln(&uid)

	wg := sync.WaitGroup{}

	wg.Add(total)

	startTime := time.Now()
	for i := 0; i < total; i++ {
		go func() {
			defer wg.Done()
			begin := time.Now()

			randomParseRequest()

			milliseconds := time.Since(begin).Milliseconds()
			spendTime = append(spendTime, milliseconds)

		}()
	}
	// 等待所有协程结束
	wg.Wait()
	totalUsedTime := time.Since(startTime).Milliseconds()

	var avgSpendTime int64
	for _, value := range spendTime {
		avgSpendTime += value
	}
	avgSpendTime = avgSpendTime / int64(len(spendTime))

	fmt.Println("---------- summary >")
	fmt.Printf("总用时：%d\n", totalUsedTime)
	fmt.Printf("总次数：%d， 成功次数 %d ,成功比例为%d \n", total, successNum, (successNum/total)*100)
	fmt.Printf("平均用时间%d \n", avgSpendTime)

	fmt.Println("---------- >")
	// 分别计算耗时

	//go func() {
	//	// 每两秒发起 100个请求
	//	for range time.Tick(2 * time.Second) {
	//		SecNum += 2
	//	}
	//}()

}

// 生成随机的解析请求
// func randomParseRequest(uid string, field []DWDField) {
func randomParseRequest() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("hhhh")
	successNum++
}
