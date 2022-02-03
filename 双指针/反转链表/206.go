package main

func main() {
	e := &ListNode{7, nil}
	d := &ListNode{5, e}
	c := &ListNode{5, d}
	b := &ListNode{3, c}
	a := &ListNode{1, b}
	reverseList(a)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var slow *ListNode
	cur := head
	for cur != nil {
		fast := cur.Next
		cur.Next = slow
		slow = cur
		cur = fast
	}
	return slow
}
