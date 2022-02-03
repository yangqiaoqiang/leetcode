package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	l1 := headA
	l2 := headB
	i, k := 0, 0
	for headA != nil {
		i++
		headA = headA.Next
	}
	for headB != nil {
		k++
		headB = headB.Next
	}
	for a := k - i; a > 0; a-- {
		l2 = l2.Next
	}
	for a := i - k; a > 0; a-- {
		l1 = l1.Next
	}
	for l1 != l2 {
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1
}
