package main

import (
	"fmt"
	// "slices"
)

func singleNumber(nums []int) int {

	// Sanity
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums)%2 == 0 {
		return 0
	}

	// Actual

	result := 0
	for _, num := range nums {
		result ^= num
	}
	return result
}

func main() {

	fmt.Println(singleNumber([]int{2, 2, 1}))
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))

}
