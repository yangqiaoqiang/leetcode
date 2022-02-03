package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//遍历
// func isSymmetric(root *TreeNode) bool {
// 	return ss(root.Left, root.Right)
// }
// func ss(left, right *TreeNode) bool {
// 	if left == nil && right == nil {
// 		return true
// 	}
// 	if left == nil || right == nil {
// 		return false
// 	}
// 	if left.Val != right.Val {
// 		return false
// 	}
// 	return ss(left.Left, right.Right) && ss(left.Right, right.Left)
// }

//迭代
func isSymmetric(root *TreeNode) bool {
	res := []*TreeNode{}
	if root != nil {
		res = append(res, root.Left, root.Right)
	}
	for len(res) > 0 {
		left := res[0]
		right := res[1]
		res = res[2:]
		if left == nil && right == nil {
			continue
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		res = append(res, left.Left, right.Right, left.Right, right.Left)
	}
	return true
}
