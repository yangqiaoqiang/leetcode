package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//递归
func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	var tarversal func(node *TreeNode)
	tarversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		tarversal(node.Left)
		tarversal(node.Right)
		res = append(res, node.Val)
	}
	tarversal(root)
	return res

}

//迭代
// func postorderTraversal(root *TreeNode) []int {
//     res:=[]int{}
//     if root==nil{
//         return res
//     }
//     li:=list.New()
//     li.PushBack(root)
//     for li.Len()>0{
//         node:=li.Remove(li.Back()).(*TreeNode)
//         res=append(res,node.Val)
//         if node.Left!=nil{
//             li.PushBack(node.Left)
//         }
//         if node.Right!=nil{
//             li.PushBack(node.Right)
//         }
//     }
//     for i,k:=0,len(res)-1;i<k;{
//         res[i],res[k]=res[k],res[i]
//         i++
//         k--
//     }
//     return res
// }
