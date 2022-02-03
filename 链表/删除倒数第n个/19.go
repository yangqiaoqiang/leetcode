package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := &ListNode{}
	flag := l
	l.Next = head
	r := head
	for r != nil {
		r = r.Next
		if n < 1 {
			l = l.Next
		}
		n--
	}
	l.Next = l.Next.Next
	return flag.Next
}
