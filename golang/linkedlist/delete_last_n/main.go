package delete_last

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	head *ListNode
	size int
}

func DeleteLastButN(list *LinkedList, n int) {
	if n < 0 || n > list.size {
		return
	}

	first := list.head
	for i := 1; i < n; i++ {
		first = first.Next
	}

	second := list.head
	for first != nil {
		first = first.Next
		second = second.Next
	}

	second.Next = second.Next.Next
}
