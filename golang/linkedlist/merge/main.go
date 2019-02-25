package merge_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	list := &ListNode{}

	if l1.Val < l2.Val {
		list.Val = l1.Val
		l1 = l1.Next
	} else {
		list.Val = l2.Val
		l2 = l2.Next
	}
	curr := list
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			curr = curr.Next
			l1 = l1.Next
		} else {
			curr.Next = l2
			curr = curr.Next
			l2 = l2.Next
		}
	}
	if l1 == nil {
		curr.Next = l2
	} else {
		curr.Next = l1
	}
	return list
}
