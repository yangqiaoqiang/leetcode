package main

import "container/list"

func main() {

}

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	queue := list.New()
	res := [][]int{}
	if root == nil {
		return res
	}
	queue.PushBack(root)
	for queue.Len() > 0 {
		length := queue.Len()
		var tmp []int
		for T := 0; T < length; T++ {
			myNode := queue.Remove(queue.Front()).(*Node)
			tmp = append(tmp, myNode.Val)
			for i := 0; i < len(myNode.Children); i++ {
				queue.PushBack(myNode.Children[i])
			}
		}
		res = append(res, tmp)
	}
	return res
}
