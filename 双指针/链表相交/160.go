package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l, r := headA, headB
	len1, len2 := 0, 0
	for l != nil {
		len1++
		l = l.Next
	}
	for r != nil {
		len2++
		r = r.Next
	}
	l, r = headA, headB
	if len1 > len2 {
		for i := 0; i < len1-len2; i++ {
			l = l.Next
		}
	} else {
		for i := 0; i < len2-len1; i++ {
			r = r.Next
		}
	}
	for l != nil {
		if l == r {
			return l
		}
		l = l.Next
		r = r.Next
	}
	return nil
}
