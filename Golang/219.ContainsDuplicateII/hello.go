package main

import (
	"fmt"
	"math"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	lookup := make(map[int]int)

	for c_index, num := range nums{
		if p_index, status := lookup[num]; status{
			// num already present in lookup. We now found a duplicate. 
			if math.Abs(float64(c_index - p_index)) <= float64(k){
				return true
			} else {
				lookup[num] = c_index
			}
		} else {
			lookup[num] = c_index
		}
	}

	return false
}

func main(){

	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1,0,1,1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1,2,3,1,2,3}, 2))

}
