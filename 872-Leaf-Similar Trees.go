package main

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
//func dfs(res *[]int, node *TreeNode) {
//	if node.Left == nil && node.Right == nil {
//		*res = append(*res, node.Val)
//	}
//
//	if node.Left != nil {
//		dfs(res, node.Left)
//	}
//	if node.Right != nil {
//		dfs(res, node.Right)
//	}
//}
//
//func readNode(root *TreeNode) []int {
//	res := []int{}
//	dfs(&res, root)
//	fmt.Println(res)
//	return res
//}
//
//func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
//	return reflect.DeepEqual(readNode(root1), readNode(root2))
//}
//
//func main() {
//	r3 := TreeNode{3, nil, nil}
//	r1 := TreeNode{1, &r3, nil}
//	r2 := TreeNode{2, nil, nil}
//	fmt.Println(leafSimilar(&r1, &r2))
//}
