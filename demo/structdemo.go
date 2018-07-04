package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("node is nil")
		return
	}
	node.value = value
}

func structDemo() {
	fmt.Println("==========structDemo=========")
	var root *treeNode
	root.setValue(123)
	root = &treeNode{
		value: 100,
	}
	root.setValue(123)
	fmt.Println(root.value)

}
