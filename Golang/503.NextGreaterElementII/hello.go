package main

func nextGreaterElements(nums []int) []int {
    // key: (index x), value: (next greater of index x)
    lookup := make(map[int]int)
    monStack := make([]int, 0, len(nums))
    result := make([]int, 0, len(nums))
	indices := make([]int, 0, len(nums))

    isFirstPass := true

    for i := 0; i < len(nums); i++{
        currNum := nums[i]

        if len(monStack) == 0 {
            monStack = append(monStack, currNum)
			indices = append(indices, i)
        } else {
            for len(monStack) > 0 && monStack[len(monStack) - 1] < currNum{
				if _, found := lookup[indices[len(monStack) - 1]]; !found{
	                lookup[indices[len(monStack) - 1]] = currNum
				}
                monStack = monStack[ : len(monStack)-1]
				indices = indices[ : len(indices) - 1]
            }
        }

        monStack = append(monStack, currNum)
		indices = append(indices, i)

        if isFirstPass && i == len(nums) - 1{
            i = -1
            isFirstPass = false
        }
    }

    for i := range nums{
        if nextGreaterNum, found := lookup[i]; found{
            result = append(result, nextGreaterNum)
        } else {
            result = append(result, -1)
        }
    }

    return result
}

func main(){
	nextGreaterElements([]int{1,2,3,4,5,6,5,4,5,1,2,3})
}