package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target))
}
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for k, v := range nums {
		if res, ok := m[target-v]; ok {
			return []int{res, k}
		} else {
			m[v] = k
		}
	}
	return []int{}
}
