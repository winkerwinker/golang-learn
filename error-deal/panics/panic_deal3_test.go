package panics

import (
	"awesomeProject/zues-learn/logger"
	"fmt"
	"github.com/sirupsen/logrus"
	"math/rand"
	"runtime/debug"
	"sync"
	"testing"
	"time"
)

func Test_4(t *testing.T) {

	ids := make([]float64, 500)
	res := make([][]float64, 500)
	wg := sync.WaitGroup{}
	wg.Add(len(ids))
	for i := 0; i < 500; i++ {
		ids[i] = rand.Float64() * 100
	}
	sex := "user_sex&"
	sys := "sys_manufacturer&"
	levelKey := "city_level&"
	predGroup := "pred_group&"
	keys := []string{sex, sys, levelKey, predGroup}

	startTime := time.Now()
	for i := 0; i < len(ids); i++ {
		index := i
		go func(gid float64) {
			defer func() {
				if err := recover(); err != nil {
					logger.Logger.WithField("panic at :", string(debug.Stack())).Info("_panic")
				}
				wg.Done()
			}()

			fmt.Println(keys)
			time.Sleep(2 * time.Second)

			res[index] = make([]float64, 100)
			logger.Logger.WithFields(logrus.Fields{"msg": res[index]}).Info("_com_load_cache")

		}(ids[index])
	}
	wg.Wait()

	logger.Logger.WithFields(logrus.Fields{"msg": time.Since(startTime)}).Info("_com_load_cache")
}
