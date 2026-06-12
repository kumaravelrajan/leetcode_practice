package main

import (
	"fmt"
	"slices"
)

func merge(intervals [][]int) [][]int {
    
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
            if intervals[i][0] <= intervals[j][0]{
                // Intervals overlap

                result = append(result, []int{intervals[i][0], max(intervals[i][1], intervals[j][1])})

            } else {
                break
            }
        }
    }

    return result
}

func main(){
	fmt.Println(merge([][]int{{1,3},{2,6},{8,10},{15,18}}))
}
