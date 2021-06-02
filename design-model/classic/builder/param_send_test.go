package builder

import (
	"fmt"
	"testing"
)

// 参数发送的测试
func TestParamSend(t *testing.T) {

	type args struct {
		name string
		opts []ResourcePoolConfigOptFunc
	}

	realargs := &args{
		name: "test",
		opts: []ResourcePoolConfigOptFunc{
			func(option *ResourcePoolConfigOption) {
				option.minIdle = 2
			},
		},
	}

	// 返回真实的对象
	// 传入操作方法！！！  和 操作的对象！！！
	//todo   args + funcs ！！！ 可以test
	config, _ := NewResourcePoolConfig(realargs.name, realargs.opts...)
	fmt.Println(config)
}
