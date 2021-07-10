package event_bus

import (
	"errors"
	"reflect"
	"runtime"
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
	funcName := getFunctionName(handlerFunc)
	// 注册的参数，如果坚听的参数已经有了。那么
	// 参数  ->  不同的funcName的map
	funcMap, ok := listeners[paramName]
	if ok {
		_, ok := funcMap[funcName]
		if ok {
			return errors.New("该方法已被注册")
		}
		funcMap[funcName] = reflect.ValueOf(handlerFunc)
		//对应的参数
		listeners[paramName] = funcMap

	} else {
		// map[string]reflect.Value)
		listeners[paramName] = map[string]reflect.Value{funcName: reflect.ValueOf(handlerFunc)}
	}
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

//获得函数名
func getFunctionName(f interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

// 完成了监听函数的注册之后，接下来就是对发送过来的消息进行处理
// 监听并转发消息
var name = make(chan interface{}, 100)
