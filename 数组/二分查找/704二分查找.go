package main

import "fmt"

func main() {
	num := []int{1, 5, 6, 7, 9}
	target := 7
	fmt.Println(search(num, target), num[search(num, target)])
}

//左闭右闭区间

func search(nums []int, target int) int {
	high := len(nums) - 1
	low := 0
	for low <= high {
		mid := low + (high-low)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

//左闭右开区间
// func search(nums []int ,target int )int {
// 	high :=len(nums)
// 	low :=0
// 	for low < high{
// 		mid := (high - low)/2 +low
// 		if nums[mid]==target {
// 			return mid
// 		}else if nums[mid] <target {
// 			low =mid +1
// 		}else{
// 			high =mid
// 		}
// 	}
// 	return -1
// }
