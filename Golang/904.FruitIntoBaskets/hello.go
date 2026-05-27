package main

// import (
// 	"math"
// 	"slices"
// )

func totalFruit(fruits []int) int {
    lookup := make(map[int]int)
    left := 0
    maxLen := 0

    for right, fruit := range fruits{
        if len(lookup) < 2 {
            lookup[fruit]++
        } else if len(lookup) == 2 {
            _, status := lookup[fruit]

            if status{
                lookup[fruit]++
            } else {
                // fruit not present in lookup. 

                for len(lookup) == 2 {
                    lookup[fruits[left]]--

                    if lookup[fruits[left]] == 0{
                        delete(lookup, fruits[left])                        
                    }

                    left++
                }

                lookup[fruit]++
            }
        }

        currLen := right - left + 1

        if currLen > maxLen{
            maxLen = currLen
        }
    }

    return maxLen
}

func main(){
	totalFruit([]int{1,2,1})
}
