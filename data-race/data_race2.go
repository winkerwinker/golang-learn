package main

import "os"

// 共享变量
func main() {
	ParallelWrite([]byte("xxx"))
}

//
//WARNING: DATA RACE
//Write at 0x00c00009c1e0 by main goroutine:
//main.ParallelWrite()
///Users/didi/go/src/awesomeProject/data-race/data_race2.go:24 +0x1dd
//main.main()
///Users/didi/go/src/awesomeProject/data-race/data_race2.go:7 +0x84
//

//todo 同时共享了err 变量
func ParallelWrite(data []byte) chan error {

	res := make(chan error, 2)
	f1, err := os.Create("/tmp/file1")
	if err != nil {
		res <- err
	} else {
		go func() {
			_, err = f1.Write(data)
			res <- err
			f1.Close()
		}()
	}

	f2, err := os.Create("/tmp/file2")
	if err != nil {
		res <- err
	} else {
		go func() {

			_, err = f2.Write(data)
			res <- err
			f2.Close()
		}()
	}

	return res



}
