package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mian() {
	e := &ListNode{7, nil}
	d := &ListNode{5, e}
	c := &ListNode{5, d}
	b := &ListNode{3, c}
	a := &ListNode{1, b}
	removeElements(a, 5)
}
func removeElements(head *ListNode, val int) *ListNode {
	vh := &ListNode{}
	vh.Next = head
	a := vh
	for a.Next != nil && a != nil {
		if a.Next.Val == val {
			a.Next = a.Next.Next
		} else {
			a = a.Next
		}
	}
	return vh.Next
}
