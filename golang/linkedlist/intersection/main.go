package intersection_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	lengthA := 0
	lengthB := 0
	currA := headA
	currB := headB
	for currA != nil {
		currA = currA.Next
		lengthA += 1
	}
	for currB != nil {
		currB = currB.Next
		lengthB += 1
	}
	if lengthA > lengthB {
		len := lengthA - lengthB
		for len > 0 {
			headA = headA.Next
			len -= 1
		}

	} else {
		len := lengthB - lengthA
		for len > 0 {
			headB = headB.Next
			len -= 1
		}
	}

	for headA != headB && headA != nil {
		headA = headA.Next
		headB = headB.Next
	}

	if headA != nil {
		return headA
	}
	return nil
}
