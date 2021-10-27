package map_test

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/copier"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

var stop chan bool

type MyStruct struct {
	test string
}

func  Predict1() (result []string, err error) {

	result[0]="hello"
	return result,nil
}


func  Predict() (result map[int64]string, err error) {

	result[1]="hello"
	return result,nil
}

func TestDo10(t *testing.T) {
	Predict1()
	// Predict()


	contains := strings.Contains("get dll resp error: Post \\\"http://10.85.130.198:16034/v1/model/tensorflow/predict\\\": context deadline exceeded", "context deadline exceeded")
	fmt.Println(contains)
	fmt.Println("-----")

	timer := time.AfterFunc(time.Duration(1)*time.Millisecond, func() {
		fmt.Println("caonima")
		return
	})

	defer timer.Stop()

	time.Sleep(10 *time.Second )
	fmt.Println("-----")
	LifeWindow := 10
	i := time.Duration(LifeWindow) * time.Hour
	fmt.Println(i)
	fmt.Println("-----")


	strs := []byte("{\"Test\":\"hello\"}")
	myStruct := &MyStruct{}
	err := json.Unmarshal(strs, myStruct)
	fmt.Println(myStruct)
	fmt.Println("-----")
	fmt.Println(err)
	fmt.Println("-----")

	serFeature := make([]float64, 117)
	fmt.Println(serFeature)
	serFeature = make([]float64, 0, 117)
	fmt.Println(serFeature)
	fmt.Println("-------")
	var test bool
	fmt.Println(test)
	fmt.Println("-------")
	curParams := Map{Dirty: map[interface{}]string{}}
	params := Map{Dirty: map[interface{}]string{}}
	curParams.Dirty["1"] = "1"

	fmt.Println(curParams.Dirty)
	fmt.Println(params.Dirty)

	copier.Copy(curParams, params)

	fmt.Println(curParams.Dirty)
	fmt.Println(params.Dirty)

	//a := "0"
	//atoi, err := strconv.Atoi(a)
	//fmt.Println(atoi)
	//fmt.Println(err)
	//
	//stop = make(chan bool)
	//test111()
	//for i := 0; i < 10; i++ {
	//
	//	stop <- true
	//}

}

type Map struct {

	// read contains the portion of the map's contents that are safe for
	// concurrent access (with or without mu held).
	//
	// The read field itself is always safe to load, but must only be stored with
	// mu held.
	//
	// Entries stored in read may be updated concurrently without mu, but updating
	// a previously-expunged entry requires that the entry be copied to the dirty
	// map and unexpunged with mu held.
	Read atomic.Value // readOnly

	// dirty contains the portion of the map's contents that require mu to be
	// held. To ensure that the dirty map can be promoted to the read map quickly,
	// it also includes all of the non-expunged entries in the read map.
	//
	// Expunged entries are not stored in the dirty map. An expunged entry in the
	// clean map must be unexpunged and added to the dirty map before a new value
	// can be stored to it.
	//
	// If the dirty map is nil, the next write to the map will initialize it by
	// making a shallow copy of the clean map, omitting stale entries.
	Dirty map[interface{}]string

	// misses counts the number of loads since the read map was last updated that
	// needed to lock mu to determine whether the key was present.
	//
	// Once enough misses have occurred to cover the cost of copying the dirty
	// map, the dirty map will be promoted to the read map (in the unamended
	// state) and the next store to the map will make a new dirty copy.
	Misses int
}

func test111() {

	go func() {
		for {
			// 控制器是一个通道 ，
			// 内部需要一个select 创建一个与控制器的链接
			select {
			case demo := <-stop:
				fmt.Println(demo)
				//continue
				return
				fmt.Println("stop111...")
			default:

			}
		}
	}()
}