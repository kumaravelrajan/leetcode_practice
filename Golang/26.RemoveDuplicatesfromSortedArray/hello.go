package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
    lookup := make(map[int]int)
	dup_i_to_remove := []int{}

	// Loop runs until second last element
	for index, value := range nums[:len(nums) - 1]{
		next_val := nums[index + 1]
		next_index := index + 1

		if value == next_val{
			dup_i_to_remove = append(dup_i_to_remove, next_index)
		}

		lookup[index] = value
	}

	lookup[len(nums)-1] = nums[len(nums) - 1]

	for _, value := range dup_i_to_remove{
		// value contains the index in nums that needs to be removed. 

		delete(lookup, value)
	}

	clear(nums)

	i := 0
	j := 0

	for i < len(nums){
		if value, status := lookup[i]; status{
			nums[j] = value
			j++
		}
		i++
		
	}

	return j
}

func main(){

	fmt.Println(removeDuplicates([]int{1,1,2}))
	fmt.Println(removeDuplicates([]int{0,0,1,1,1,2,2,3,3,4}))
}
