package main

import (
	"sort"
)

func main() {

}
func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	num1 := []int{}
	for k, _ := range m {
		num1 = append(num1, k)
	}
	sort.Slice(num1, func(i, j int) bool {
		return m[num1[i]] > m[num1[j]]
	})
	num1 = num1[:k]
	return num1
}
