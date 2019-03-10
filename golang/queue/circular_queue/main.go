package circular_queue

// head = (tail+1)%n

type CircularQueue struct {
	data []int
	n    int
	head int
	tail int
}

func NewQueue(n int) CircularQueue {
	return CircularQueue{make([]int, n), n, 0, 0}
}

func (this *CircularQueue) EnQueue(item int) bool {
	if (this.tail+1)%this.n == this.head {
		return false
	}
	this.data[this.tail] = item
	this.tail = (this.tail + 1) % this.n
	return true
}

func (this *CircularQueue) DeQueue() int {
	if this.tail == this.head {
		return ^int(^uint(0) >> 1)
	}
	val := this.data[this.head]
	this.head = (this.head + 1) % this.n
	return val
}
