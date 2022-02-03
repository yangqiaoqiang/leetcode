package main

func main() {
	e := &ListNode{7, nil}
	d := &ListNode{5, e}
	c := &ListNode{5, d}
	b := &ListNode{3, c}
	a := &ListNode{1, b}
	removeNthFromEnd(a, 2)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	pre := &ListNode{}
	pre.Next = head
	slow, fast := pre, pre
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return pre.Next
}
