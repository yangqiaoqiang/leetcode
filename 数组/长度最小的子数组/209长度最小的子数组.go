// 给定一个含有 n 个正整数的数组和一个正整数 target 。

// 找出该数组中满足其和 ≥ target 的长度最小的
// 连续子数组 [numsl, numsl+1, ..., numsr-1, numsr]
// ，并返回其长度。如果不存在符合条件的子数组，返回 0 。
package main

import (
	"fmt"
)

func main() {
	num := []int{2, 3, 1, 2, 4, 3}
	target := 7
	fmt.Println(minSubArrayLen(target, num))

}
func minSubArrayLen(target int, nums []int) int {
	d := 0
	s := len(nums) - 1
	h := 0
	sum := 0
	flag := s + 2
	for h <= s && d <= s {
		for i := d; i <= h; i++ {
			sum += nums[i]
		}
		if sum >= target {
			if flag > h-d+1 {
				flag = h - d + 1
			}

			d++
			if d > h {
				return 1
			}
		} else {
			h++
		}
		sum = 0
	}
	if flag == s+2 {
		return 0
	}
	return flag
}
