package main

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	// If the array is empty or has only one interval, it's already merged.
	if len(intervals) <= 1 {
		return intervals
	}

	// Step 1: Sort the intervals based on their start values.
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// Step 2: Initialize the result slice with the first interval.
	var result [][]int
	result = append(result, intervals[0])

	// Step 3: Iterate through the remaining intervals.
	for i := 1; i < len(intervals); i++ {
		current := intervals[i]
		lastMerged := result[len(result)-1] // A reference to the last interval in our result list

		// Check for overlap: does the current interval start before or when the last one ends?
		if current[0] <= lastMerged[1] {
			// Overlap found. Update the end time of the last merged interval to be the maximum of both.
			if current[1] > lastMerged[1] {
				lastMerged[1] = current[1]
			}
		} else {
			// No overlap. Append the current interval to the result as a new independent interval.
			result = append(result, current)
		}
	}

	return result
}

func main(){
	fmt.Println(merge([][]int{{2,3},{4,5},{6,7},{8,9},{1,10}}))
}
