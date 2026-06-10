package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
    monStack := make([]int, 0, len(nums2))
    lookup := make(map[int]int)
    // tos := -1

    for i := len(nums2) - 1; i >= 0; i--{
        currNum := nums2[i]
        if len(monStack) == 0{
            monStack = append(monStack, currNum)
            lookup[currNum] = -1
        } else {
            for len(monStack) > 0 && monStack[len(monStack)-1] < currNum{
                monStack = monStack[:len(monStack)-1]
            }

            monStack = append(monStack, currNum)

            if len(monStack) > 1 {
                lookup[currNum] = monStack[len(monStack) - 2]
            } else {
                lookup[currNum] = -1
            }
        }
    }

    result := make([]int, 0, len(nums1))

    for _, currNum := range nums1{
        result = append(result, lookup[currNum])
    }

    return result
}

func main(){
	nextGreaterElement([]int{2,4}, []int{1,2,3,4})
}
