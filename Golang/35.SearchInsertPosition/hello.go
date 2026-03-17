package main

import "fmt"

func searchInsert(nums []int, target int) int {
    start_index := 0
	end_index := len(nums) - 1
	var mid int

	for start_index <= end_index{
		mid = start_index + ((end_index - start_index) / 2)

		if nums[mid] == target{
			return mid
		} else if nums[mid] < target{
			start_index = mid + 1
		} else if nums[mid] > target{
			end_index = mid - 1
		}
	}

	return start_index
}

func main(){

	fmt.Println(searchInsert([]int{1,3,5,6}, 5))
	fmt.Println(searchInsert([]int{1,3,5,6}, 2))
	fmt.Println(searchInsert([]int{1,3,5,6}, 7))
}
