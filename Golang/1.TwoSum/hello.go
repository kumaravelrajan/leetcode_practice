package main

import (
	"fmt"
	// "slices"
)

func twoSum(nums []int, target int) []int {
    lookup := make(map[int]int)
	
	for index, num := range nums{
		if lookup_index, status := lookup[target - num]; status{
			// target - num + num = target; target - num already present in lookup. Now, num also present. Hence, return the indices. 

			return []int{index, lookup_index}
		} else {
			// target-num not yet present in lookup. Hence, just add num to the lookup and hope that some future num in nums considers the currently added num as the missing piece to achieve target.

			lookup[num] = index
		}
	}

	return []int{}
	
}

func main(){

	fmt.Println(twoSum([]int{2,7,11,15}, 9))
	fmt.Println(twoSum([]int{3,2,4}, 6))
	fmt.Println(twoSum([]int{3,3}, 6))

}