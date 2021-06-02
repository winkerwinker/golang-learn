package main

import (
	"sync"
	"testing"
)

// sleepgroup，类似于java中的countdown
func Test1(t *testing.T) {
	wg.Add(1)
	go say("Hello World")
	wg.Wait()
}

var wg = sync.WaitGroup{}

func say(s string) {
	println(s)
	wg.Done()
}
