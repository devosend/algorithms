package array_queue

type ArrayQueue struct {
	data     []int
	capacity int
	tail     int
	head     int
}

// 初始化一个队列
func NewQueue(n int) ArrayQueue {
	return ArrayQueue{make([]int, n), n, 0, 0}
}

//入队
func (this *ArrayQueue) EnQueue(item int) bool {
	if this.tail >= this.capacity {
		return false
	}
	this.data[this.tail] = item
	this.tail += 1
	return true
}

//出队
func (this *ArrayQueue) DeQueue() int {
	if this.head >= this.tail {
		return ^int(^uint(0) >> 1)
	}
	val := this.data[this.head]
	this.head += 1
	return val
}
