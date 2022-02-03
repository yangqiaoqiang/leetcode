package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	li := list.New()
	li.PushBack(root)
	for li.Len() > 0 {
		flag := []int{}
		l := li.Len()
		for i := 0; i < l; i++ {
			node := li.Remove(li.Front()).(*TreeNode)
			if node.Left != nil {
				li.PushBack(node.Left)
			}
			if node.Right != nil {
				li.PushBack(node.Right)
			}
			flag = append(flag, node.Val)
		}
		flag1 := flag[len(flag)-1]
		res = append(res, flag1)
	}
	return res
}
