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
	swapPairs(a)
}
func swapPairs(head *ListNode) *ListNode {
	a := &ListNode{}
	a.Next = head
	l := a
	r := a.Next
	for r != nil && r.Next != nil {
		l.Next = r.Next
		flag := r.Next.Next
		r.Next.Next = r
		r.Next = flag
		l = r
		r = r.Next
	}
	return a.Next
}
