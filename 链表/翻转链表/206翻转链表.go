package main

func main() {
	e := &ListNode{7, nil}
	d := &ListNode{5, e}
	c := &ListNode{5, d}
	b := &ListNode{3, c}
	a := &ListNode{1, b}
	reverseList(a)
}
func reverseList(head *ListNode) *ListNode {
	var l *ListNode
	r := head
	for r != nil {
		flag := r.Next
		r.Next = l
		l = r
		r = flag
	}
	return l
}

type ListNode struct {
	Val  int
	Next *ListNode
}
