package skip_list

import (
	"math"
	"math/rand"
)

const MAX_LEVEL = 16

type SkipNode struct {
	val      int
	level    int
	forwards []*SkipNode
}

type SkipList struct {
	head *SkipNode
	// 跳表当前深度
	level  int
	length int
}

func NewSkipNode(val int) *SkipNode {
	return &SkipNode{val, MAX_LEVEL, make([]*SkipNode, MAX_LEVEL, MAX_LEVEL)}
}

func NewSkipList() SkipList {
	node := NewSkipNode(math.MinInt32)
	return SkipList{node, 1, 0}
}

func GetNewNodeLevel() int {
	level := 1
	for i := level; i < MAX_LEVEL; i++ {
		if rand.Int31()%7 == 1 {
			level++
		}
	}
	return level
}

func (this *SkipList) Find(val int) *SkipNode {
	if this.length == 0 {
		return nil
	}

	p := this.head
	level := this.level
	for i := level - 1; i >= 0; i-- {
		for p.forwards[i] != nil {
			if p.forwards[i].val == val {
				return p.forwards[i]
			} else if p.forwards[i].val < val {
				p = p.forwards[i]
			} else {
				break
			}
		}
	}
	return nil
}

func (this *SkipList) Insert(val int) bool {
	level := GetNewNodeLevel()
	node := NewSkipNode(val)
	node.level = level
	p := this.head
	update := make([]*SkipNode, MAX_LEVEL, MAX_LEVEL)
	for i := level - 1; i >= 0; i-- {
		for p.forwards[i] != nil {
			if p.forwards[i].val == val {
				return false
			} else if p.forwards[i].val < val {
				p = p.forwards[i]
			} else {
				update[i] = p
				break
			}
		}
		if p.forwards[i] == nil {
			update[i] = p
		}
	}

	for i := 0; i < level; i++ {
		node.forwards[i] = update[i].forwards[i]
		update[i].forwards[i] = node
	}

	if level > this.level {
		this.level = level
	}

	this.length++
	return true
}

func (this *SkipList) Delete(val int) bool {
	p := this.head
	update := make([]*SkipNode, MAX_LEVEL, MAX_LEVEL)
	for i := this.level - 1; i >= 0; i-- {
		update[i] = this.head
		for p.forwards[i] != nil {
			if p.forwards[i].val == val {
				update[i] = p
				break
			} else {
				p = p.forwards[i]
			}
		}
	}

	if update[0].forwards[0] == nil || update[0].forwards[0].val != val {
		return true
	}

	for i := 0; i < this.level; i++ {
		if update[i].forwards[i] != nil && update[i].forwards[i].val == val {
			update[i].forwards[i] = update[i].forwards[i].forwards[i]
		}
		if update[i] == this.head && update[i].forwards[i] == nil {
			this.level = i
		}
	}
	this.length--
	return true
}
