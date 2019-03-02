package array_stack

type MyStack struct {
	data []int
	top  int
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{nil, -1}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
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

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	if this.top < 0 {
		return 0
	}
	v := this.data[this.top]
	this.top--
	return v
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if this.top < 0 {
		return 0
	}
	v := this.data[this.top]
	return v
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.top < 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
