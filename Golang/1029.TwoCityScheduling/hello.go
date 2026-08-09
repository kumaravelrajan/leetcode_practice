package main

import (
	"fmt"
	"slices"
	"cmp"
)

func twoCitySchedCost(costs [][]int) int {

    // index 0 is corresponding index in costs. index 1 is diff between two costs.
    diff := [][]int{}

    // Initialize diff slice
    for i, curr := range costs {
        if curr[0] > curr[1]{
            diff = append(diff, []int{i, curr[0] - curr[1]})
        } else {
            diff = append(diff, []int{i, curr[1] - curr[0]})
        }
    }

    // Sort diff slice in descending order wrt index 1
    slices.SortFunc(diff, func(a, b []int) int {
        return cmp.Compare(b[1], a[1])
    })

    // Var declaration
    result := 0
    a, b := 0, 0

    for i, currDiff := range diff {
        if a == len(costs)/2 {
            // a is complete. Add all b costs to result and return it.
            return returnFinalResult(1, result, diff, i, costs)
        } else if b == len(costs) / 2 {
            return returnFinalResult(0, result, diff, i, costs)
        }

        if costs[currDiff[0]][0] < costs[currDiff[0]][1]{
            result += costs[currDiff[0]][0]
            a++
        } else {
            result += costs[currDiff[0]][1]
            b++
        }
    }
    return result
}

// If a count is already len(costs)/2 then then indexToTarget in costs should be set to 1 because we only want to target b values. 
// i gives us the index of diff to start our calculations from.
func returnFinalResult(indexToTarget int, result int, diff [][]int, i int, costs [][]int) (int){
    for i < len(diff) {
        currDiffElem := diff[i]

        result += costs[currDiffElem[0]][indexToTarget]
        i++
    }
    return result
}

func main(){
	fmt.Println(twoCitySchedCost([][]int{
		{259, 770},
		{448, 54},
		{926, 667},
		{184, 139},
		{840, 118},
		{577, 469},
	}))
}