package bibao

import (
	"fmt"
	"sync"
	"testing"
)

var ints []int

func TestContext1(t *testing.T) {
	sPool := NewInt64sPool()
	get := sPool.pools[0].pool.Get()
	fmt.Println(get)

	fmt.Println(ints)

}

func NewInt64sPool() pool {
	var pools []*sizedPool
	pools = append(pools, newEmptySizedPool(10))
	pools = append(pools, newEmptySizedPool(11))

	p := pool{minSize: 10, maxSize: 12, pools: pools}
	p.setNew(makeInt64sPointer)

	return p
}

func makeInt64sPointer(size int) interface{} {
	data := make([]int64, 0, size)
	return &data
}

// pool is actually multiple pools which store a specific data structure of specific size.
// i.e. it can be three pools which return []int64 32K, 64K and 128K.
type pool struct {
	minSize int
	maxSize int
	pools   []*sizedPool
}

type sizedPool struct {
	Size int
	pool sync.Pool
}

func (p *pool) setNew(new func(int) interface{}) {

	for _, sizedPool := range p.pools {
		sizedPool := sizedPool
		//size := sizedPool.Size
		ints = append(ints, sizedPool.Size)
		sizedPool.pool.New = func() interface{} { return new(sizedPool.Size) }
	}
}

func newEmptySizedPool(size int) *sizedPool {
	return &sizedPool{
		Size: size,
		pool: sync.Pool{},
	}
}
