// 给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，
// 同时保持非零元素的相对顺序。
package main

import "fmt"

func main() {
	num := []int{0, 1, 5, 0, 8, 0, 4}
	moveZeroes(num)
	fmt.Println(num)
}
func moveZeroes(nums []int) {
	l := len(nums)
	res := 0
	for i := 0; i < l; i++ {
		if nums[i] != 0 {
			nums[res] = nums[i]
			res++
		}
	}
	// fmt.Println(res)
	for res < l {
		nums[res] = 0
		res++
	}
}
