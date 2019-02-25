package middle_node

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	curr := head
	currDouble := head
	for currDouble != nil && currDouble.Next != nil {
		curr = curr.Next
		currDouble = currDouble.Next.Next
	}
	return curr
}
