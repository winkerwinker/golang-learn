package slot

import (
	"context"
	"fmt"
	"golang.org/x/time/rate"
	"sync"
	"testing"
)

func Test(t *testing.T) {

	limiters := &sync.Map{}
	limiter := rate.NewLimiter(10, 10)

	l, _ := limiters.LoadOrStore("11", limiter)

	if err := l.(*rate.Limiter).Wait(context.Background()); err != nil {
		// 这里先不处理日志了，如果返回错误就直接 429
		fmt.Printf("%+v \n ", err)
	}

}
