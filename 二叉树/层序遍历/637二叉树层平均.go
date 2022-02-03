package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	res := [][]int{}
	var finRes []float64
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
	length := len(res)
	for i := 0; i < length; i++ {
		var sum int
		for j := 0; j < len(res[i]); j++ {
			sum += res[i][j]
		}
		tmp := float64(sum) / float64(len(res[i]))
		finRes = append(finRes, tmp)
	}
	return finRes
}
