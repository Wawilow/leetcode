package idk

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func abs(ans *[]int, tree TreeNode) {
	*ans = append(*ans, tree.Val)
	if tree.Left != nil {
		abs(ans, *tree.Left)
	}
	if tree.Right != nil {
		abs(ans, *tree.Right)
	}
}

func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	abs(&ans, *root)
	return ans
}
