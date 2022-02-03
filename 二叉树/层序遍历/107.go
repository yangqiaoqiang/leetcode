package main

import "container/list"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	res := [][]int{}
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
		res = append(res, flag)
	}
	i, j := 0, len(res)-1
	for i < j {
		res[i], res[j] = res[j], res[i]
		i++
		j--
	}
	return res
}
