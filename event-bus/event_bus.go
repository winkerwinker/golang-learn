package event_bus

import (
	"errors"
	"reflect"
	"sync"
)

// 最近项目中需要用到观察者模式来实现一些逻辑，
// 如某些操作的数据变更会影响到同项目中另一模块的数据。
// 操作数据会影响到另一块数据

//现在有两种消息类型，
//四个函数

// 注册锁
var (
	registerMu sync.RWMutex
	listeners  = make(map[string]map[string]reflect.Value)
)

func Register(handlerFunc interface{}) error {

	registerMu.Lock()
	defer registerMu.Unlock()

	// 处理方法
	funcType := reflect.TypeOf(handlerFunc)
	if err := checkFunc(funcType); err != nil {
		return err
	}
	paramName := generateParameterName(funcType)
	_ = paramName
	return nil
}

// 检查参数是否合法
func checkFunc(t reflect.Type) error {

	if t.Kind() != reflect.Func {
		return errors.New("only listen by func")
	}
	if t.NumIn() != 1 {
		return errors.New("not standard func, must have more than one param")
	}

	return nil
}

// 获取参数类型名
func generateParameterName(t reflect.Type) string {
	in := t.In(0)
	// get the type
	return in.PkgPath() + "." + in.Name()
}

// 获得函数名
//func getFunctionName(f interface{}) string {
//return runtime.FuncForPC()
//}
