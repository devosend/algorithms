package loop

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	curr := head
	currDouble := head
	for currDouble != nil && currDouble.Next != nil {
		curr = curr.Next
		currDouble = currDouble.Next.Next
		if curr == currDouble {
			return true
		}
	}
	return false
}
