package main

func subarraySum(nums []int, k int) int {

    prefixSumLookup := make(map[int]int)
    currSum := 0
    result := 0

    for index, num := range nums{
        currSum += num
        prefixSumLookup[currSum] = index
    }

    for currPrefixSum, currIndex := range prefixSumLookup{
        if prevIndex, ok := prefixSumLookup[currPrefixSum - k]; ok{
            if prevIndex < currIndex{
                result++
            }
        }

        if currPrefixSum == k {
            result++
        }
    }

    return result
} 

func main(){

	subarraySum([]int{-1,-1,1}, 0)

}
