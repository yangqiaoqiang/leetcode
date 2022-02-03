package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	res := 0
	q := []*TreeNode{}
	q = append(q, root)
	if root == nil {
		return res
	}
	for i := len(q); i > 0; {
		res++
		for ; i > 0; i-- {
			node := q[0]
			q = q[1:]
			if node.Left == nil && node.Right == nil {
				return res
			}
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		i = len(q)
	}
	return res
}
