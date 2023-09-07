package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func dfs(node *ListNode, hashMap *map[int]*ListNode, i int) {
	nextNode := node.Next
	node.Next = nil
	(*hashMap)[i] = node

	if nextNode != nil {
		dfs(nextNode, hashMap, i+1)
	}
	return
}

func deDfs(node *ListNode, hashMap *map[int]*ListNode, i int) {
	nextNode, ok := (*hashMap)[i]
	node.Next = nextNode

	if nextNode != nil && ok {
		deDfs(nextNode, hashMap, i+1)
	}
	return
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// If no need to change anything than just return
	if left == right {
		return head
	}

	// Make list from node
	hashMap := map[int]*ListNode{}
	dfs(head, &hashMap, 0)

	// Set up right order
	for i := 0; i <= (right-left)/2; i++ {
		hashMap[left-1+i], hashMap[right-1-i] = hashMap[right-1-i], hashMap[left-1+i]
	}

	// Rebuild new node
	newHead := hashMap[0]
	deDfs(newHead, &hashMap, 1)

	// Return new node
	return newHead
}
