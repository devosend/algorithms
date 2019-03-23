package binary_tree

import "fmt"

func PreTraversal(node *Node) {
	if node == nil {
		return
	}
	fmt.Println(node.data)
	PreTraversal(node.left)
	PreTraversal(node.right)
}

func InTraversal(node *Node) {
	if node == nil {
		return
	}
	InTraversal(node.left)
	fmt.Println(node.data)
	InTraversal(node.right)
}

func PostTraversal(node *Node) {
	if node == nil {
		return
	}
	PostTraversal(node.left)
	PostTraversal(node.right)
	fmt.Println(node.data)
}
