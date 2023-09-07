package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a := []int{1, 2, 3}
	fmt.Println(a[1:])
}

func get_len(head *ListNode) int {
	res := 0
	for head != nil {
		head = head.Next
		res++
	}
	return res
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	if k == 1 {
		return []*ListNode{head}
	}

	l := get_len(head)
	norm_size, num_inc, res := l/k, l%k, []*ListNode{}

	for i := 0; i < k; i++ {
		size := norm_size
		if num_inc > 0 {
			size++
			num_inc--
		}
		if size == 0 {
			res = append(res, nil)
			continue
		}

		res = append(res, head)
		for j := 0; j < size-1; j++ {
			head = head.Next
		}

		tmp := head.Next
		head.Next = nil
		head = tmp
	}

	return res
}
