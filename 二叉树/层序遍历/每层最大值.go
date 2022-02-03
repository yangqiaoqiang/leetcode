package main

import (
	"container/list"
	"math"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	res := [][]int{}
	var finRes []int
	if root == nil {
		return finRes
	}
	queue := list.New()
	queue.PushBack(root)
	var tmpArr []int
	for queue.Len() > 0 {
		length := queue.Len()
		for i := 0; i < length; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
			tmpArr = append(tmpArr, node.Val)
		}
		res = append(res, tmpArr)
		tmpArr = []int{}
	}
	for i := 0; i < len(res); i++ {
		finRes = append(finRes, max(res[i]...))
	}
	return finRes
}
func max(vals ...int) int {
	max := int(math.Inf(-1))
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}
