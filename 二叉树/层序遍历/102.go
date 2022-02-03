package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	li := list.New()
	li.PushBack(root)
	for li.Len() > 0 {
		l := li.Len()
		flag := []int{}
		for i := 0; i <= l-1; i++ {
			node := li.Remove(li.Front()).(*TreeNode)
			if node.Left != nil {
				li.PushBack(node.Left)
			}
			if node.Right != nil {
				li.PushBack(node.Right)
			}
			flag = append(flag, node.Val)
		}
		res = append(res, flag)
	}
	return res
}
