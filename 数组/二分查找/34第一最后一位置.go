// 给定一个按照升序排列的整数数组 nums，和一个目标值 target。
// 找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回 [-1, -1]。
// 进阶：
// 你可以设计并实现时间复杂度为 O(log n) 的算法解决此问题吗？
package main

import "fmt"

func main() {
	num := []int{1}
	target := 1
	fmt.Println(search(num, target))
}
func search(nums []int, target int) []int {
	right := len(nums) - 1
	left := 0
	result := -1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			result = mid
			break
		} else if nums[mid] >= target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	if result == -1 {
		return []int{-1, -1}
	}
	for i := result; i >= 0; i-- {
		if nums[i] == target {
			left = i
		} else {
			break
		}
	}
	for i := result; i < len(nums); i++ {
		if nums[i] == target {
			right = i
		} else {
			break
		}
	}
	return []int{left, right}
}
