package main

// import (
// 	"math"
// 	"slices"
// )

func totalFruit(fruits []int) int {
    last, secondLast := -1, -1
	consecutiveLastCount, curr, max := 0, 0, 0

	for _, fruit := range fruits{
		if fruit == last{
			curr++
			consecutiveLastCount++
		} else if fruit == secondLast{
			curr++
		} else {
			secondLast = last
			last = fruit
			curr = consecutiveLastCount + 1
			consecutiveLastCount = 1
		}

		if curr > max {
			max = curr
		}

	}

	return max
}

func main(){
	totalFruit([]int{1, 1})
}
