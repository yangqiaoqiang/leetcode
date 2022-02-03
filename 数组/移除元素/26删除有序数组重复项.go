// 给你一个有序数组 nums ，请你 原地 删除重复出现的元素，
// 使每个元素 只出现一次 ，返回删除后数组的新长度。

// 不要使用额外的数组空间，你必须在 原地 修改输入数组
// 并在使用 O(1) 额外空间的条件下完成。
package main

import "fmt"

func main() {
	num := []int{1, 1, 2, 2, 3, 3, 3, 4, 5}
	fmt.Println(removeDuplicates(num))
	fmt.Println(num)
}
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	res := 0
	for i := 1; i < l; i++ {
		if nums[i] != nums[res] {
			res++
			nums[res] = nums[i]
		}
	}
	return res + 1
}
