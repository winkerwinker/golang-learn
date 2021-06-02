package builder

import (
	"testing"
)

func TestBuild(t *testing.T) {

	builder := ResourcePoolConfigBuilder{
		name:     "",
		maxTotal: 0,
	}

	// 并不能连接在一起 ，需要每次返回自身
	builder.SetMaxIdle(1)
	builder.SetMinIdle(1)
	// 灵魂， 类似k8s中的do
	builder.Build()

}
