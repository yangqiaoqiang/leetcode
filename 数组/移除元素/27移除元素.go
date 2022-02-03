// 给你一个数组 nums 和一个值 val
// 你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

// 不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
// 元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。
package main

import "fmt"

func main() {
	num := []int{3, 2, 4, 5, 3, 2}
	val := 3
	fmt.Println(removeElement(num, val))
	fmt.Println(num)
}
func removeElement(nums []int, val int) int {
	l := len(nums)
	res := 0
	for i := 0; i < l; i++ {
		if nums[i] != val {
			nums[res] = nums[i]
			res++
		}
	}
	return res
}
