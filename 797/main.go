package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	res := [][]int{}

	//max := 0
	//for _, vv := range graph {
	//	for _, v := range vv {
	//		if v > max {
	//			max = v
	//		}
	//	}
	//}

	dfs(&res, graph, []int{}, 0)

	return res
}

func dfs(res *[][]int, graph [][]int, path []int, node int) {
	if node != len(graph)-1 {
		//fmt.Println(node)
		neightbour := graph[node]
		for _, v := range neightbour {
			//fmt.Println(v)
			dfs(res, graph, append(path, node), v)
		}
	} else {
		fmt.Println(append(path, node), node)
		add := append(path, node)
		if add[len(add)-1] != len(graph)-1 {
			add = append(add, len(graph))
		}
		*res = append(*res, add)
	}
}

func main() {
	//fmt.Println(allPathsSourceTarget([][]int{[]int{1, 2}, []int{3}, []int{3}, []int{}}))
	fmt.Println(allPathsSourceTarget([][]int{[]int{5, 1}, []int{5, 3, 7, 2}, []int{7, 4, 6, 3, 5}, []int{4, 7, 5}, []int{6, 5, 7}, []int{7, 6}, []int{7}, []int{}}))
}
