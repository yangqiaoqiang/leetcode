package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归法
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

//迭代法
// func invertTree(root *TreeNode) *TreeNode {
//     res :=[]*TreeNode{}
//     node:=root
//     for node !=nil||len(res)>0{
//         for node !=nil{
//            node.Left,node.Right=node.Right,node.Left
//            res =append(res,node)
//            node =node.Left
//         }
//         node =res[len(res)-1]
//         res =res[:len(res)-1]
//         node =node.Right

//     }
//     return root
// }
