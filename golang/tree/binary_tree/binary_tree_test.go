package binary_tree

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := NewTree(6)
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(17)
	tree.Insert(2)
	tree.Insert(1)
	InTraversal(tree)
	tree.Delete(6)
	InTraversal(tree)
}
