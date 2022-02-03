// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 请必须使用时间复杂度为 O(log n) 的算法。

package main

import "fmt"

func main() {
	num := []int{1, 5, 6, 7, 9}
	target := 10
	fmt.Println(searchInsert(num, target))
}
func searchInsert(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	for low <= high {
		mid := (high-low)/2 + low
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return low
}
