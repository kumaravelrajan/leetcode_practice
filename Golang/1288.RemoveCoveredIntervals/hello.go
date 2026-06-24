package main

import (
	"slices"
	"cmp"
)

func removeCoveredIntervals(intervals [][]int) int {
    if len(intervals) <= 1{
        return len(intervals)
    }

	result := 1
	coveredIndices := make(map[int]bool)
    
    fp, sp := 0, 1

	// Sort the slice according to index 0 in ascending
    slices.SortFunc(intervals, func (a, b []int) (int){
        return cmp.Compare(a[0], b[0])
    })


    for sp < len(intervals){
        if intervals[sp][1]<= intervals[fp][1]{
            // Covered interval found
			coveredIndices[sp] = true
            sp++
        } else {
			_, found := coveredIndices[fp+1]
			for found{
				fp++
				_, found = coveredIndices[fp+1]
			}
			fp++
			result++
        }
    }

    return fp + 1
}

func main(){
	removeCoveredIntervals([][]int{{3,10},{4,10},{5,11}})

}
