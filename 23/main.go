package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	bigList := []int{}
	for _, vv := range lists {
		for _, v := range readNode(vv) {
			bigList = append(bigList, v)
		}
	}
	for i := 0; i < len(bigList)-1; i++ {
		for j := 0; j < len(bigList)-i-1; j++ {
			if bigList[j] > bigList[j+1] {
				bigList[j], bigList[j+1] = bigList[j+1], bigList[j]
			}
		}
	}
	if len(bigList) == 0 {
		return &ListNode{}
	}
	//revBigList := make([]int, len(bigList))
	revBigList := []int{}
	for i := 0; i < len(bigList); i++ {
		revBigList[i] = bigList[len(bigList)-i-1]
	}
	res := *wNode(revBigList)
	return &res
}

func wNode(list []int) *ListNode {

	res := wDfs(list, nil)

	return res
}

func wDfs(list []int, res *ListNode) *ListNode {
	if len(list) != 0 {
		return wDfs(list[1:], &ListNode{list[0], res})
	} else {
		return res
	}
}

func readNode(node *ListNode) []int {
	res := []int{}
	rDfs(node, &res)

	return res
}

func rDfs(node *ListNode, res *[]int) {
	*res = append(*res, node.Val)
	if node.Next != nil {
		rDfs(node.Next, res)
	}
}

func main() {
	node := ListNode{
		Val: 1, Next: &ListNode{
			Val: 4, Next: &ListNode{5, nil},
		},
	}
	node2 := ListNode{
		Val: 1, Next: &ListNode{
			Val: 4, Next: &ListNode{5, nil},
		},
	}
	node3 := ListNode{
		Val: 2, Next: &ListNode{
			Val: 6, Next: nil},
	}
	a := mergeKLists([]*ListNode{&node, &node2, &node3})
	//a := mergeKLists([]*ListNode{})
	dfs(a)
}

func dfs(a *ListNode) {
	fmt.Println(a.Val)
	if a.Next != nil {
		dfs(a.Next)
	}
}
