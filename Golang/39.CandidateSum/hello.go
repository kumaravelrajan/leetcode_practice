package main

func combinationSum(candidates []int, target int) [][]int {
    result := [][]int{}
    current := []int{}

    backtrack(candidates, target, 0, &result, current)

    return result
}

func backtrack (candidates []int, target int, start int, result *[][]int, current []int){

    if target == 0 {
        // Combination found
        temp := make([]int, len(current))
        copy(temp, current)
        *result = append(*result, temp)
    }

    if target < 0 {
        return
    }

    for i := start; i < len(candidates); i++{
        current = append(current, candidates[i])
        backtrack(candidates, target - candidates[i], i, result, current)
		current = current[: len(current)-1]
    }
} 

func main(){
	combinationSum([]int{2,3,6,7},7)
}