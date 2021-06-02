package singeton

import (
	"sync"
)

// 单例模式

var lazySingleton *Singleton
var once = &sync.Once{}

func GetLazyInstance() *Singleton {

	if lazySingleton == nil {
		// 保证只创建一次
		once.Do(func() {
			lazySingleton = &Singleton{}
		})
	}
	return lazySingleton
}
