package main

import "fmt"

// claude

func alternateNums(arr []int)([]int){
	
	result := make([]int, 0, len(arr))
	n := len(arr)

	// pi → next positive index, ni → next negative index
	pi, ni := 0, 0
	for pi < n && arr[pi] < 0 {
		pi++
	}
	for ni < n && arr[ni] >= 0 {
		ni++
	}

	for pi < n || ni < n {
		// Place next positive
		if pi < n {
			result = append(result, arr[pi])
			pi++
			for pi < n && arr[pi] < 0 {
				pi++
			}
		}
		// Place next negative
		if ni < n {
			result = append(result, arr[ni])
			ni++
			for ni < n && arr[ni] >= 0 {
				ni++
			}
		}
	}

	return result
}

func main(){

	fmt.Println(alternateNums([]int{-5,4,8,2,-9,10,3,-2}))

}
