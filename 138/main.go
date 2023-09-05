package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var testdata = Node{
	Val: 7,
	Next: &Node{
		Val: 13,
		Next: &Node{
			Val: 11,
			Next: &Node{
				Val:    10,
				Next:   nil,
				Random: nil,
			},
			Random: &Node{
				Val:    1,
				Next:   nil,
				Random: nil,
			},
		},
		Random: nil,
	},
	Random: nil,
}

func main() {
	copyRandomList(&testdata)
}

func dfs(copy *Node, i int, nodeMap *map[*Node]*Node, node *Node) {
	if node == nil {
		return
	}
	copy.Val = node.Val
	copy.Random = nil
	if node.Next != nil {
		copy.Next = &Node{}
	}
	(*nodeMap)[node] = copy

	if node.Next != nil {
		dfs(copy.Next, i+1, nodeMap, node.Next)
	}
}

func dfsForRandom(copy *Node, nodeMap *map[*Node]*Node, node *Node) {
	r, ok := (*nodeMap)[node.Random]
	if ok {
		copy.Random = r
	}
	if node.Next != nil {
		dfsForRandom(copy.Next, nodeMap, node.Next)
	}
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	nodeMap := map[*Node]*Node{}
	newHead := &Node{}
	dfs(newHead, 0, &nodeMap, head)
	dfsForRandom(newHead, &nodeMap, head)
	return newHead
}
