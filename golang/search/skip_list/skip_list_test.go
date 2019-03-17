package skip_list

import "testing"

func TestSkipList(t *testing.T) {
	list := NewSkipList()
	list.Insert(8)
	list.Insert(4)
	list.Insert(7)
	list.Insert(10)
	p := list.head
	for i := 0; i < list.length; i++ {
		t.Log(p.forwards[0].val)
		p = p.forwards[0]
	}

	node := list.Find(4)
	t.Log(node)

	node = list.Find(5)
	t.Log(node)

	list.Delete(4)
	p = list.head
	for i := 0; i < list.length; i++ {
		t.Log(p.forwards[0].val)
		p = p.forwards[0]
	}
}
