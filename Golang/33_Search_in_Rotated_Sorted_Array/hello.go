package main

import (
	"fmt"
	// "slices"
)

func search(nums []int, target int) int {
    left, right := 0, len(nums)-1

    for left <= right {
        mid := left + (right-left)/2

        // Check if we found the target
        if nums[mid] == target {
            return mid
        }

        // Determine which half is strictly sorted
        if nums[left] <= nums[mid] {
            // The left half is sorted
            // Check if the target falls within this sorted left half
            if nums[left] <= target && target < nums[mid] {
                right = mid - 1 // Narrow search to the left half
            } else {
                left = mid + 1  // Target must be in the right half
            }
        } else {
            // The right half is sorted
            // Check if the target falls within this sorted right half
            if nums[mid] < target && target <= nums[right] {
                left = mid + 1  // Narrow search to the right half
            } else {
                right = mid - 1 // Target must be in the left half
            }
        }
    }

    // Target not found
    return -1
}

func main(){

	fmt.Println(search([]int{4,5,6,7,0,1,2}, 0))
	fmt.Println(search([]int{4,5,6,7,0,1,2}, 3))
	fmt.Println(search([]int{1}, 0))
	fmt.Println(search([]int{3,1}, 1))
}
