package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type readNodeList struct {
	res []int
}

func (r *readNodeList) readNode(node *TreeNode) {
	if node != nil {
		r.readNode(node.Left)
		r.readNode(node.Right)
		r.res = append(r.res, node.Val)
	}
}

func postorderTraversal(root *TreeNode) []int {
	r := readNodeList{}
	r.readNode(root)
	return r.res
}

func main() {
	one := &TreeNode{3, nil, nil}
	two := &TreeNode{2, one, nil}
	simpleNode := &TreeNode{1, nil, two}

	r := readNodeList{}
	r.readNode(simpleNode)
	fmt.Print(r.res, "\n")
}
