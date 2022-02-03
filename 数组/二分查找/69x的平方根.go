// 给你一个非负整数 x ，计算并返回 x 的 算术平方根 。

// 由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。

// 注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
package main

import (
	"fmt"
)

func main() {
	i := 8
	fmt.Println(sqrt(i))
}
func sqrt(x int) int {
	l, r := 1, x/2+1
	for l < r {
		mid := (r-l)/2 + l
		if mid == x/mid {
			return mid
		} else if mid > x/mid {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return r
}
