package main

import "fmt"

func dfsrec(adj_list [][]int, visited []bool, res []int, s int) []int {

	visited[s] = true
	res = append(res, s)

	for _, neighour := range adj_list[s] {
		if !visited[neighour] {
			res = dfsrec(adj_list, visited, res, neighour)
		}
	}

	return res
}

func dfs(adj_list [][]int) []int {
	vertex_count := len(adj_list)
	visited := make([]bool, vertex_count)
	res := []int{}

	return dfsrec(adj_list, visited, res, 0)
}

func main() {

	res := dfs([][]int{{1, 2}, {0, 2}, {0, 1, 3, 4}, {2}, {2}})

	fmt.Println(res)

}
