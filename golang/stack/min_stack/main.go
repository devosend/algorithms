package min_stack

type MinStack struct {
	data []int
	min  []int
	top  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	var stack MinStack
	stack.min = append(stack.min, int(^uint(0)>>1))
	stack.top = -1
	return stack
}

func (this *MinStack) Push(x int) {
	if this.top < 0 {
		this.top = 0
	} else {
		this.top++
	}

	if this.top < len(this.data) {
		this.data[this.top] = x
	} else {
		this.data = append(this.data, x)
	}

	if x <= this.min[len(this.min)-1] {
		this.min = append(this.min, x)
	}
}

func (this *MinStack) Pop() {
	if this.top < 0 {
		return
	}
	if this.data[this.top] == this.min[len(this.min)-1] {
		this.min = this.min[:len(this.min)-1]
	}
	this.top--
}

func (this *MinStack) Top() int {
	if this.top < 0 {
		return 0
	}
	v := this.data[this.top]
	return v
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
