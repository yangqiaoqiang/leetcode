package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	fmt.Println(removeElement(nums, val))
}
func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast <= len(nums)-1 {
		if nums[fast] == val {
			fast++
			continue
		}
		nums[slow] = nums[fast]
		slow++
		fast++
	}
	return slow
}
