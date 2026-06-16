package main

import (
	"fmt"
	"slices"
)

func merge(intervals [][]int) [][]int {

	lookup := make(map[int]bool)
    
    result := make([][]int, 0, len(intervals))

    slices.SortFunc(intervals, func (a []int, b []int) int {
        if a[0] < b[0]{
            return -1
        } else if a[0] > b[0]{
            return 1
        } else {
            return 0
        }
    })

    for i := 0; i < len(intervals) - 1; i++{
        for j := i + 1; j < len(intervals); j++{
            if intervals[j][0] <= intervals[i][1]{
                // Intervals overlap

                result = append(result, []int{intervals[i][0], max(intervals[i][1], intervals[j][1])})
				
				lookup[j] = true
				lookup[i] = true

            } else {
                break
            }
        }
    }

	for i := 0; i < len(intervals); i++{
		if _, found := lookup[i]; !found{
			result = append(result, intervals[i])
		}
	}

    return result
}

func main(){
	fmt.Println(merge([][]int{{1,3},{2,6},{8,10},{15,18}}))
	fmt.Println(merge([][]int{{0,2},{1,4},{3,5}}))
	
}
