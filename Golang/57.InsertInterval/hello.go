package main

func insert(intervals [][]int, newInterval []int) [][]int {

	if len(intervals) == 0{
		return [][]int{newInterval}
	}

    result := make([][]int, 0, len(intervals))

	isOverlapFound := false

    for i := 0; i < len(intervals); i++{
        if intervals[i][1] >= newInterval[0] && !isOverlapFound{
            // Overlap found

			isOverlapFound = true

            startOfMergeInterval := min(intervals[i][0], newInterval[0])

            j := i + 1

            for j < len(intervals) && newInterval[1] >= intervals[j][0]{
                j++
            }

            endOfMergeInterval := max(newInterval[1], intervals[j-1][1])

            result = append(result, []int{startOfMergeInterval, endOfMergeInterval})
            
			i = j - 1

        } else {
            result = append(result, intervals[i])
        }
    }

	if !isOverlapFound{
		result = append(result, newInterval)
	}

    return result
}

func main(){
	insert([][]int{{1,5}}, []int{6,8})
}
