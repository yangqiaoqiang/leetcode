package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if slow == fast {
			pre := head
			for pre != slow {
				pre = pre.Next
				slow = slow.Next
			}
			return pre
		}
	}
	return nil

}
