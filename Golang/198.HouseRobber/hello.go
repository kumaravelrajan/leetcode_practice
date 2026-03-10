package main

// ! Solved using AI.

import (
	"fmt"
)

func rob(nums []int) int {
    ct_index0 := 0
	prev_index0 := -2
	sum0 := 0

	ct_index1 := 1
	prev_index1 := -1
	sum1 := 0

	for ct_index0 < len(nums) || ct_index1 < len(nums){
		if (ct_index0 < len(nums)) && (ct_index0 > prev_index0 + 1) {
			sum0 += nums[ct_index0]
			prev_index0 = ct_index0
			ct_index0 += 1
		}

		if (ct_index1 < len(nums)) && (ct_index1 > prev_index1 + 1){
			sum1 += nums[ct_index1]
			prev_index1 = ct_index1
			ct_index1 += 1
		}

		ct_index0++
		ct_index1++

	}

	if sum0 > sum1{
		return sum0
	} else if sum1 > sum0 {
		return sum1
	} else if sum1 == sum0{
		return sum1
	}

	return -1
}

func main(){
	fmt.Println(rob([]int{1,2,3,1}))
	fmt.Println(rob([]int{2,7,9,3,1}))
	fmt.Println(rob([]int{2,1,1,2}))
}