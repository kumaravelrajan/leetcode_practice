package main

import "fmt"

func bubblesort(sl []int) []int {

	for i := 0; i < len(sl)-1; i++ {

		for j := 0; j < len(sl)-i-1; j++ {

			if sl[j] > sl[j+1] {
				sl[j], sl[j+1] = sl[j+1], sl[j]
			}
		}
	}

	return sl
}

func main() {
	fmt.Println(bubblesort([]int{5, 6, 1, 3}))
}
