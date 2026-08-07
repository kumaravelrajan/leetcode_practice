package main

import "fmt"

/*
- Take a BFS like approach but without the queue. Instead with pointers.
- At every index, we know the max extent of the jump possible. This max extent becomes the boundary of the current level.
- Browse through the current level and find max value. If current level includes last index, we already have our answer.
- Else, find the new next level boundary by using the highest value in the current level.
*/

func jump(nums []int) int {
    currentEnd, nextEnd, jumps := 0, 0, 0

    for i := 0; i < len(nums); i++{
        if i <= currentEnd{

            if nextEnd < nums[i] + i{
                nextEnd = nums[i] + i
            }
        } else {
            jumps++

            currentEnd = nextEnd

            if currentEnd >= len(nums) - 1 {
                return jumps
            }

            if nextEnd < i + nums[i]{
                nextEnd = i + nums[i]
            }
        }
    }

    return -1
}

func main(){
	fmt.Println(jump([]int{2,3,1,1,4}))
	fmt.Println(jump([]int{2,3,0,1,4}))
	fmt.Println(jump([]int{7,0,9,6,9,6,1,7,9,0,1,2,9,0,3}))
}
