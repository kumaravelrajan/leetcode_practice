package main

import "fmt"

func removeElement(nums []int, val int) int {
    
	if len(nums) == 0{
		return 0
	}
	
	freq_val := 0
	index_to_insert := 0

	for _, value := range nums{
		if value == val{
			freq_val++
		} else {
			nums[index_to_insert] = value
			index_to_insert++
		}
	}

	return len(nums) - freq_val
}

func main(){

	fmt.Println(removeElement([]int{3,2,2,3}, 3))
	fmt.Println(removeElement([]int{0,1,2,2,3,0,4,2}, 2))

}
