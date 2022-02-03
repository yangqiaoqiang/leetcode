package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		traversal(node.Left)
		traversal(node.Right)
	}
	traversal(root)
	return res
}

//迭代
// func preorderTraversal(root *TreeNode) []int {
//     res:=[]int{}
//     if root ==nil{
//         return res
//     }
//     li:=list.New()
//     li.PushBack(root)
//     for li.Len()>0{
//         node:=li.Back().Value.(*TreeNode)
//         li.Remove(li.Back())
//         res=append(res,node.Val)
//         if node.Right!=nil{
//             li.PushBack(node.Right)
//         }
//         if node.Left!=nil{
//             li.PushBack(node.Left)
//         }
//     }
//     return res
// }
