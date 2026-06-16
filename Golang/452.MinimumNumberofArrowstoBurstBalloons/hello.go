package main

import ("slices"
"cmp")

func findMinArrowShots(points [][]int) int {
    if len(points) <= 1{
        return len(points)
    }
    
    slices.SortFunc(points, func (a, b []int) (int) {
        return cmp.Compare(a[0], b[0])
    })

    result := 0
    result++
    var lastInterval []int
    lastInterval = points[0]

    for i := 1; i < len(points); i++{
        currPoint := points[i]

        if !(currPoint[0] <= lastInterval[1] && currPoint[1] >= lastInterval[0]){
            // New range detected
            result++
            lastInterval = currPoint
        }
    }

    return result
}

func main(){
	findMinArrowShots([][]int{
	{3, 9},
	{7, 12},
	{3, 8},
	{6, 8},
	{9, 10},
	{2, 9},
	{0, 9},
	{3, 9},
	{0, 6},
	{2, 8},
})
}