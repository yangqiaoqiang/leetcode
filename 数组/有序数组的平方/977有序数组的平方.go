// 给你一个按 非递减顺序 排序的整数数组 nums，
// 返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
package main

import "fmt"

func main() {
	num := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(num))
}
func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	zz := make([]int, r+1)
	res := r
	for r >= l {
		if nums[r]*nums[r] >= nums[l]*nums[l] {
			zz[res] = nums[r] * nums[r]
			res--
			r--
		} else {
			zz[res] = nums[l] * nums[l]
			res--
			l++
		}
	}
	return zz
}
