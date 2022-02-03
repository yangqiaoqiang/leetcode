package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var tarversal func(node *TreeNode)
	tarversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		tarversal(node.Left)
		res = append(res, node.Val)
		tarversal(node.Right)
	}
	tarversal(root)
	return res
}

//迭代
// func inorderTraversal(root *TreeNode) []int {
// 	ans := []int{}
// 	if root == nil {
// 		return ans
// 	}

// 	st := list.New()
// 	cur := root

// 	for cur != nil || st.Len() > 0 {
// 		if cur != nil {
// 			st.PushBack(cur)
// 			cur = cur.Left
// 		} else {
// 			cur = st.Remove(st.Back()).(*TreeNode)
// 			ans = append(ans, cur.Val)
// 			cur = cur.Right
// 		}
// 	}

// 	return ans
// }
