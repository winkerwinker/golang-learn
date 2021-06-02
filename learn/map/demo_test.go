package map_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

var BifrostGoodFeatureLoadTask *BifrostCacheLoadTask

// 组装map
var BifrostTaskChannel = make(chan *BifrostCacheLoadTask)
var BifrostLoadTasks = make(map[string]*BifrostCacheLoadTask)

func UpdateBifrostCacheLoadTask(task *BifrostCacheLoadTask) {
	BifrostTaskChannel <- task
}


func TestDo1(t *testing.T) {
	fmt.Println(BifrostLoadTasks)
	fmt.Println(BifrostLoadTasks["GoodGeneralFeatureTask"])


	for i := 0; i < 2000; i++ {
		go func() {
			task := BifrostLoadTasks["GoodGeneralFeatureTask"]
			task.TaskName =strconv.Itoa(i) +"updateOK"
			//UpdateBifrostCacheLoadTask(task)
			//BifrostLoadTasks["GoodGeneralFeatureTask"]= task
			//fmt.Println(strconv.Itoa(i) +"updateOK")
		}()
	}
	fmt.Println("allDone")
	time.Sleep(4 * time.Second)
	loadTask := BifrostLoadTasks["GoodGeneralFeatureTask"]
	fmt.Println(loadTask)
}

// 不更改map ，不会发生改变
func TestDo(t *testing.T) {
	fmt.Println(BifrostLoadTasks)
	for i := 0; i < 2000; i++ {
		go func() {
			task := BifrostLoadTasks["GoodGeneralFeatureTask"]
			task.TaskName = "qqqqq"
			task.TaskName =strconv.Itoa(i) +"updateOK"
			//BifrostLoadTasks["GoodGeneralFeatureTask"]= task
			fmt.Println("updateOK")
		}()
	}
	fmt.Println("allDone")
	time.Sleep(4 * time.Second)
	loadTask := BifrostLoadTasks["GoodGeneralFeatureTask"]
	fmt.Println(loadTask)
}

func listenForUpdate() {
	for {
		// todo 如果有panic是否需要recover ，再继续for循环
		newTask := <-BifrostTaskChannel
		_, ok := BifrostLoadTasks[newTask.TaskName]
		if !ok {
			continue
		}

		BifrostLoadTasks[newTask.TaskName] = newTask
	}
}

func init() {

	//BifrostGoodFeatureLoadTask = &BifrostCacheLoadTask{
	//	TaskName:    "GoodFeature",
	//	CacheName:   localcache.Cache_1,
	//	LoadDate:    "-1",
	//	LoadTime:    "-1",
	//	LocalFolder: "/tmp/features/",
	//}

	BifrostGoodFeatureLoadTask = &BifrostCacheLoadTask{
		TaskName:    "GoodGeneralFeatureTask",
		CacheName:   "hhh",
		LoadDate:    "-1",
		LoadTime:    "-1",
		LocalFolder: "....",
	}

	BifrostLoadTasks[BifrostGoodFeatureLoadTask.TaskName] = BifrostGoodFeatureLoadTask

	// 起协程监听更新
	go listenForUpdate()
}

type BifrostCacheLoadTask struct {
	TaskName    string
	CacheName   string
	LoadDate    string
	LoadTime    string
	LocalFolder string
}
