package main

import "fmt"

func main() {
	i := 3
	fmt.Println(generateMatrix(i))
}
func generateMatrix(n int) [][]int {
	sum := 1
	top, bot := 0, n-1
	left, right := 0, n-1
	target := n * n
	nums := make([][]int, n)
	for i := 0; i <= bot; i++ {
		nums[i] = make([]int, n)
	}
	for sum <= target {
		for i := left; i <= right; i++ {
			nums[top][i] = sum
			sum++
		}
		top++
		for i := top; i <= bot; i++ {
			nums[i][right] = sum
			sum++
		}
		right--
		for i := right; i >= left; i-- {
			nums[bot][i] = sum
			sum++
		}
		bot--
		for i := bot; i >= top; i-- {
			nums[i][left] = sum
			sum++
		}
		left++
	}
	return nums
}
