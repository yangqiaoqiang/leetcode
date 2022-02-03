package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 2, 1}
	nums2 := []int{2, 2}
	fmt.Println(intersection(nums1, nums2))
}
func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	m := make(map[int]int)
	for _, v := range nums1 {
		m[v] = 1
	}
	for _, v := range nums2 {
		if m[v] == 1 {
			res = append(res, v)
			m[v]--
		}
	}
	return res
}
