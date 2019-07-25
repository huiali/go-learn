package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func creatTreeNode(value int) *treeNode {
	return &treeNode{value: value}
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = creatTreeNode(2)
	root.right.left.setValue(6)

	//root.print()
	//root.right.left.print()

	root.traverse()

}
