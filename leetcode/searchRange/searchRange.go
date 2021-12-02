package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchRange([]int{5, 6, 6, 6, 7}, 6))

}

//你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
// 简化问题 为， 二分中找到charset中重复多个中的第一个
// 简化问题 为， 二分中找到charset中重复多个中的最后一个
func searchRange(nums []int, target int) []int {

	if len(nums) == 0 || nums[0] > target || nums[len(nums)-1] < target {
		return []int{-1, -1}
	}

	first := searchFirst(nums, target)

	second := searchSecond(nums, target)

	return []int{first, second}

}

func searchFirst(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	tmpRes := -1
	for {

		if left > right {
			return tmpRes
		}

		midIndex := (left + right + 1) / 2

		if nums[midIndex] < target {
			left = midIndex + 1
		} else if nums[midIndex] > target {
			right = midIndex - 1
		} else {
			right = midIndex - 1
			// 找到可能的结果了
			tmpRes = midIndex
		}

	}

}

func searchSecond(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	tmpRes := -1
	for {
		if left > right {
			return tmpRes
		}

		midIndex := (left + right + 1) / 2

		if nums[midIndex] < target {
			left = midIndex + 1
		} else if nums[midIndex] > target {
			right = midIndex - 1
		} else {
			left = midIndex + 1
			// 找到可能的结果了
			tmpRes = midIndex
		}

	}

}
