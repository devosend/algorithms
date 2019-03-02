package min_stack

type MinStack struct {
	data []int
	top  int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{nil, -1}
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
}

func (this *MinStack) Pop() {
	if this.top < 0 {
		return
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
	if this.top < 0 {
		return 0
	}
	v := this.data[0]
	for i := 0; i <= this.top; i++ {
		if this.data[i] < v {
			v = this.data[i]
		}
	}
	return v
}
