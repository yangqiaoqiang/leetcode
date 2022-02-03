package main

import "sort"

func main() {

}
func threeSum(nums []int) [][]int {
	flag := 0
	res := [][]int{}
	sort.Ints(nums)
	for flag <= len(nums)-3 {
		if nums[flag] > 0 {
			break
		}
		if flag > 0 && nums[flag] == nums[flag-1] {
			flag++
			continue
		}
		slow, fast := flag+1, len(nums)-1
		for slow < fast {
			if nums[flag]+nums[slow]+nums[fast] == 0 {
				res = append(res, []int{nums[flag], nums[slow], nums[fast]})
				slow++
				fast--
				for slow < fast {
					if nums[slow] == nums[slow-1] {
						slow++
					} else {
						break
					}
				}
				for slow < fast {
					if nums[fast] == nums[fast+1] {
						fast--
					} else {
						break
					}
				}
			} else if nums[flag]+nums[slow]+nums[fast] > 0 {
				fast--
			} else {
				slow++
			}
		}
		flag++
	}
	return res
}
