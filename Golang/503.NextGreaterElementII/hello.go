package main

func nextGreaterElements(nums []int) []int {
    // key: (x), value: (next greater of x)
    lookup := make(map[int]int)
    monStack := make([]int, 0, len(nums))
    result := make([]int, 0, len(nums))

    isFirstPass := true

    for i := 0; i < len(nums); i++{
        currNum := nums[i]

        if len(monStack) == 0 {
            monStack = append(monStack, currNum)
        } else {
            for len(monStack) > 0 && monStack[len(monStack)-1] < currNum{
				if _, found := lookup[monStack[len(monStack)-1]]; !found{
	                lookup[monStack[len(monStack)-1]] = currNum
				}
                monStack = monStack[ : len(monStack)-1]
            }
        }

        if len(monStack) > 0 && monStack[len(monStack) - 1] == currNum {
            
        } else {
            monStack = append(monStack, currNum)
        }

        if isFirstPass && i == len(nums) - 1{
            i = -1
            isFirstPass = false
        }
    }

    for _, currNum := range nums{
        if nextGreaterNum, found := lookup[currNum]; found{
            result = append(result, nextGreaterNum)
        } else {
            result = append(result, -1)
        }
    }

    return result
}

func main(){
	nextGreaterElements([]int{5,4,3,2,1})
}