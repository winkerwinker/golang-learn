package singeton_test

import (
	singeton "awesomeProject/design-model/classic/singleton"
	"testing"
)
// todo 会自动被添加到go mod 里面
import "github.com/stretchr/testify/assert"

// 单例模式
func TestGetInstance(t *testing.T) {
	assert.Equal(t, singeton.GetInstance(), singeton.GetInstance())
}

//
//Benchmark  基准测试
//go.exe test -benchmem -bench="." -v
func BenchMarkGetINstanceParaller(b *testing.B) {
	//
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if singeton.GetInstance() != singeton.GetInstance() {
				b.Errorf("test fail")
			}
		}
	})

}
