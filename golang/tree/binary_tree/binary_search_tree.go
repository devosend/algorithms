package binary_tree

type Node struct {
	data  int
	left  *Node
	right *Node
}

func newTree(val int) *Node {
	return &Node{val, nil, nil}
}

func (this *Node) insert(val int) bool {
	p := this
	node := &Node{val, nil, nil}
	for p != nil {
		if val < p.data {
			if p.left == nil {
				p.left = node
				return true
			}
			p = p.left
		}
		if val > p.data {
			if p.right == nil {
				p.right = node
				return true
			}
		}
	}
	return false
}

func (this *Node) find(val int) *Node {
	p := this
	for p != nil {
		if val == p.data {
			return p
		}
		if val < p.data {
			p = this.left
		} else {
			p = this.right
		}
	}
	return nil
}

func (this *Node) delete(val int) bool {
	p := this
	parent := this
	for p != nil && p.data != val {
		parent = this
		if val < p.data {
			p = this.left
		} else {
			p = this.right
		}
	}

	if p == nil {
		return false
	}

	if p.left == nil && p.right == nil {
		if p == parent.left {
			parent.left = nil
		} else {
			parent.right = nil
		}
		return true
	}

	if p.right == nil {
		if p == parent.left {
			parent.left = p.left
		} else {
			parent.right = p.left
		}
		return true
	}

	if p.left == nil {
		if p == parent.left {
			parent.left = p.right
		} else {
			parent.right = p.right
		}
		return true
	}

	min := p
	minParent := p
	for min != nil {
		minParent = min
		min = min.left
	}
	p.data = min.data
	minParent.left = nil
	return true
}
