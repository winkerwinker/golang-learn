package gonum_test

import (
	"errors"
	"fmt"
	"sync"
	"testing"
)

func Test_1(t *testing.T) {

	result := make([]interface{}, 100)

	// resCh := make(chan interface{}, 100)
	errCh := make(chan error, 100)

	var wg sync.WaitGroup
	wg.Add(100)

	for i := 0; i < 100; i++ {
		index := i
		go func() {
			defer wg.Done()
			if index%10 == 0 {
				result[4] = 0
				result[index] = 0
				errCh <- errors.New("xxx")
				return
			}
			result[index] = index
		}()
	}

	wg.Wait()
	fmt.Println(result)

}
