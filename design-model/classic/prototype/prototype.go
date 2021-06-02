package prototype

import (
	"encoding/json"
	"time"
)

//接下来会按照一个类似课程中的例子使用深拷贝和浅拷贝结合的方式进行实现
//需求: 假设现在数据库中有大量数据，包含了关键词，关键词被搜索的次数等信息，模块 A 为了业务需要
//会在启动时加载这部分数据到内存中
//并且需要定时更新里面的数据
//同时展示给用户的数据每次必须要是相同版本的数据，不能一部分数据来自版本 1 一部分来自版本 2

// todo 原型和本地缓存的区别

type KeyWord struct {
	word     string
	visit    int
	UpdateAt *time.Time
}

// clone 使用序列化与凡序列化的方式进行拷贝
func (k *KeyWord) CLone() *KeyWord {

	var keyword KeyWord
	marshal, _ := json.Marshal(keyword)
	json.Unmarshal(marshal, &keyword)
	return &keyword
}

// 定义一个map
type Keywords map[string]*KeyWord

func (k *Keywords) CLone(updateWords []*KeyWord) *Keywords {
	keywords := Keywords{}

	for k, v := range keywords {
		// 这里是浅拷贝，只拷贝了地址
		keywords[k] = v
	}

	//如果是深拷贝，调用上文的Clone方法
	return nil
}
