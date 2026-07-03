package main

import "fmt"

func permute(nums []int) [][]int {
    visited := make([]bool, len(nums))
    result := [][]int{}

    var backtrack func ([]int)

    backtrack = func (current []int){

        if len(current) == len(nums){
            // New permutation found
            temp := make([]int, len(nums))
            copy(temp, current)
            result = append(result, temp)
            return
        }

        for i := 0; i < len(nums); i++{
            if !visited[i]{
                visited[i] = true
                current = append(current, nums[i])
                backtrack(current)
				current = current[: len(current) - 1]
				visited[i] = false
            }
        }

    }

    backtrack([]int{})

    return result
}

func main(){
	fmt.Println(permute([]int{1,2,3}))
}
