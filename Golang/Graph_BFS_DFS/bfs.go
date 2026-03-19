package main

import (
	"fmt"
	// "slices"
)

func bfs(adj_matrix [][]int)([]int){
	num_vertex := len(adj_matrix)

	visited := make([]bool, num_vertex)
	res := []int{}
	queue := []int{}

	src := 0
	visited[src] = true
	queue = append(queue, src)

	for len(queue) > 0{
		curr := queue[0]
		queue = queue[1:]

		res = append(res, curr)
		curr_neighbours := adj_matrix[curr]

		for _, neighbour := range curr_neighbours{
			if !visited[neighbour]{
				queue = append(queue, neighbour)
				visited[neighbour] = true
			}
		}
	}

	return res
}

func main(){
	fmt.Println(bfs([][]int{{1, 2}, {0, 2}, {0, 1, 3, 4}, {2}, {2}}))
}
