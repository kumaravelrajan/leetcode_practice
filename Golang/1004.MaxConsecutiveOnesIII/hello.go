package main

func longestOnes(nums []int, k int) int {
    left := 0
    maxLen := 0
	indicesOfZero := make([]int, len(nums))
	rightOfZeroIndex, leftOfZeroIndex := 0, 0

    for right, bit := range nums{
        if bit == 0{
            // bit is 0

			indicesOfZero[rightOfZeroIndex] = right
			rightOfZeroIndex++

            if k > 0 {
                k--
            } else {
				left = indicesOfZero[leftOfZeroIndex] + 1
				leftOfZeroIndex++
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
	longestOnes([]int{1,1,1,0,0,0,1,1,1,1,0}, 2)
}
