package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	if root.Left != nil {
		res += countNodes(root.Left)
	}
	if root.Right != nil {
		res += countNodes(root.Right)
	}
	return res
}
