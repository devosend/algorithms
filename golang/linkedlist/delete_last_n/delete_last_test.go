package delete_last

import (
	"testing"
)

func TestListDeleteLast(t *testing.T) {
	node := ListNode{1, nil}
	list := LinkedList{&node, 1}
	nums := []int{2, 3, 4, 5}
	for _, num := range nums {
		node := ListNode{num, nil}
		curr := list.head
		for i := 0; curr.Next != nil; i++ {
			curr = curr.Next
		}
		curr.Next = &node
		list.size++
	}

	list.DeleteLastButN(2)
	if list.size != 4 {
		t.FailNow()
		t.Log("Failed Not Expected Size")
	}
	curr := list.head
	for i := 1; i < 5; i++ {
		if i == 4 {
			if curr.Val != 5 {
				t.FailNow()
				t.Log("Failed.")
			}
		} else {
			if curr.Val != i {
				t.FailNow()
				t.Log("Failed.")
			}
		}
		curr = curr.Next
	}
}
