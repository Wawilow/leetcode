package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func dfs(root *TreeNode) (*TreeNode, int) {
	if root == nil {
		return nil, 0
	}

	ll, l := dfs(root.Left)
	rr, r := dfs(root.Right)

	if l == r {
		return root, l + 1
	} else if l > r {
		return ll, l + 1
	} else {
		return rr, r + 1
	}
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	res, _ := dfs(root)
	// _, res := deep(root)
	return res
}
