package singeton

// 单例模式

//  todo 不能暴露出创建
type Singleton struct {
}

var singleton *Singleton

func init() {
	singleton = &Singleton{}
}

func GetInstance() *Singleton {
	return singleton
}
