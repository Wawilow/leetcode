package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeList struct {
	list []int
}

func (r *nodeList) readNote(node *TreeNode) {
	if node != nil {
		r.list = append(r.list, node.Val)
		r.readNote(node.Left)
		r.readNote(node.Right)
	}
}

func inorderTraversal(root *TreeNode) []int {
	r := nodeList{}
	r.readNote(root)
	return r.list
}

func main() {
	node_3 := TreeNode{Val: 3, Left: nil, Right: nil}
	node_2 := TreeNode{Val: 2, Left: nil, Right: &node_3}
	node_1 := &TreeNode{Val: 1, Left: &node_2, Right: nil}

	fmt.Print(inorderTraversal(node_1), "\n")
}
