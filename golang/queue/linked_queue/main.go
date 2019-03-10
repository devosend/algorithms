package linked_queue

type Node struct {
	val  int
	next *Node
}

type LinkQueue struct {
	head *Node
	tail *Node
}

func NewQueue() LinkQueue {
	return LinkQueue{nil, nil}
}

func (this *LinkQueue) EnQueue(item int) {
	node := Node{item, nil}
	if this.tail == nil {
		this.tail = &node
		this.head = &node
	} else {
		this.tail.next = &node
		this.tail = &node
	}
}

func (this *LinkQueue) DeQueue() *Node {
	if this.head == nil {
		return nil
	}
	node := this.head
	this.head = this.head.next
	return node
}
