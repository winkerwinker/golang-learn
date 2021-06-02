package context

import (
	"fmt"
	"strconv"
	"time"
)

func main() {


	m := make(map[string]int,2)
	m["heh"] =64

	for key,value := range m {

		fmt.Println(key)
		fmt.Println(value)
	}
	fmt.Println("--------")

	// 会自动设置大小[0 0] ，所以应该设置为0
	result := make([]int64, 2)

	num, err2 := strconv.ParseInt("uus", 10, 64)
	_ = err2
	fmt.Println(num)
	//if err2 == nil {
	result = append(result, num)
	// 	}

	fmt.Println(result)
	fmt.Println("--------")

	// 也会有默认的大小
	res := make([]string, 2)

	fmt.Println(num)
	//if err2 == nil {
	res = append(res, "tttr")
	// 	}

	fmt.Println(res)
	for key, value := range res {

		fmt.Println(key)
		fmt.Println(value)
	}
	fmt.Println("--------")

	//var uid string
	//res := `{"addcar":"1","buyNow":"1","toDetail":"1"}`
	//var resultMap = make(map[string]string)
	//err1 := json.Unmarshal([]byte(res), &resultMap)
	//fmt.Println(err1)

	//fmt.Scanln(&uid)
	//fmt.Println("--------")
	//fmt.Println(uid)
	//
	//fmt.Println("--------")
	data := make([]int, 0)
	data = append(data, 1)
	data = append(data, 2)
	data = append(data, 3)
	data = append(data, 4)

	loopData := func(handleData chan<- int) {
		defer close(handleData)
		for i := range data {
			time.Sleep(5 * time.Second)
			handleData <- data[i]
		}
	}

	handleData := make(chan int)

	go loopData(handleData)

	//handleData 一定发生在那之后吗
	for num := range handleData {
		fmt.Println(num)
	}

}
