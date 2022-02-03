package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

//迭代
// func maxDepth(root *TreeNode) int {
// 	res := 0
// 	q := []*TreeNode{}
// 	if root == nil {
// 		return res
// 	}
// 	q = append(q, root)
// 	for i := len(q); i > 0; {
// 		for ; i > 0; i-- {
// 			node := q[0]
// 			q = q[1:]
// 			if node.Left != nil {
// 				q = append(q, node.Left)
// 			}
// 			if node.Right != nil {
// 				q = append(q, node.Right)
// 			}
// 		}
// 		i = len(q)
// 		res++
// 	}
// 	return res
// }
