package main

import "fmt"

func main() {
	power := maxPower("hooraaaaaaaaaaay")
	fmt.Println("==========================")
	fmt.Println("res:", power)
}

func maxPower(s string) int {

	if len(s) <= 1 {
		return len(s)
	}
	max := 0
	tmpMax := 0
	var firstChar string
	for _, c := range s {
		if firstChar != string(c) {
			max = maxUtil(max, tmpMax)
			firstChar = string(c)
			tmpMax = 1
			//fmt.Println("==========================")
			//fmt.Println("tmpMax:", tmpMax, "\nmax:", max, "\nchar:", string(c))
		} else {
			tmpMax++
		}
		//fmt.Println("out:", index, "c:", string(c))
	}

	max = maxUtil(max, tmpMax)
	return max
}

func maxUtil(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
