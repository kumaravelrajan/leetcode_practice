package main

import (
	"fmt"
	// "slices"
)

func bfs(adj_list [][]int)([]int){
	v := len(adj_list)

	visited := make([]bool, v)
	result := make([]int, 0, v)
	q := make([]int, 0, v)

	src := 0

	q = append(q, src)
	visited[src] = true

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		result = append(result, curr)

		for _, adj := range adj_list[curr]{
			if visited[adj] == false{
				visited[adj] = true
				q = append(q, adj)
			}
		}
	}
	return result
}

func main(){
	fmt.Println(bfs([][]int{{1, 2}, {0, 2}, {0, 1, 3, 4}, {2}, {2}}))
}
