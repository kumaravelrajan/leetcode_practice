package main

// import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    result := make([]int, len(nums1))

    for i := 0; i < len(result); i++{
        result[i]  = -1
    }

    // Monotonic decreasing stack
    monStack := []int{}

    lookupNums1 := make(map[int]int)

    for i, val := range nums1{
        lookupNums1[val] = i
    }

    for _, val := range nums2{
        if len(monStack) == 0 {
            monStack = append(monStack, val)
        } else {
            top := monStack[len(monStack) - 1]

            if val < top{
                monStack = append(monStack, val)
            } else {
                for val > top {
                    if _, found := lookupNums1[top]; found{
                        result[lookupNums1[top]] = val
                    }
                    monStack = monStack[: len(monStack) - 1]
                    if len(monStack) > 0 {
                        top = monStack[len(monStack) - 1]
                    } else {
                        break
                    }
                }
                monStack = append(monStack, val)
            }
        }
    }

    return result
}

func main(){
	nextGreaterElement([]int{1,3,5,2,4}, []int{6,5,4,3,2,1,7})
}
