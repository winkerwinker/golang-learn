package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(3333))
}

// 回文数
func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	//x= x
	flag := 10
	bits := make([]int, 0)
	//判断一个数是几位数
	for {

		t := x / flag   //10
		bit := x % flag //0
		bits = append(bits, bit*10/flag)

		flag *= 10

		if t < 10 {
			bits = append(bits, t)
			break
		}
	}

	fmt.Println(bits)
	for i := 0; i < len(bits)/2; i++ {
		if bits[i] != bits[len(bits)-1-i] {
			return false
		}
	}

	return true
}
