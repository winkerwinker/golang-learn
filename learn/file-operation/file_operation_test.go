package echo_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

// 文件操作

func TestParamSend(t *testing.T) {
	dir := "/Users/didi/Desktop/shell/"
	files, _ := ioutil.ReadDir(dir)

	// 如果推送前一天的
	// fileDateMap := make(map[string]time.Time)
	// 最晚的时间
	var maxTime time.Time

	fmt.Println(maxTime)

	// 找到最小的日期
	for _, f := range files {
		fi, _ := os.Stat(dir + f.Name())
		// fileDateMap[fi.Name()] = fi.ModTime()

		tempTime := fi.ModTime()
		if tempTime.After(maxTime) {
			maxTime = tempTime
		}
	}

	for _, f := range files {
		filePath := dir + f.Name()
		fi, _ := os.Stat(filePath)
		if fi.ModTime().Year() != maxTime.Year() || fi.ModTime().Month() != maxTime.Month() || fi.ModTime().Day() != maxTime.Day() {
			os.RemoveAll(dir + f.Name())
		}
	}

}
