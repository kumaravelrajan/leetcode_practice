package main

func searchRange(nums []int, target int) []int {
    left, right := 0, len(nums) - 1
    mid := -1
    isTargetFound := false

    if len(nums) == 0 {
        return []int{-1,-1}
    } else if len(nums) == 1{
        if target == nums[0]{
            return []int{0,0}
        } else {
            return []int{-1,-1}
        }
    }

    for left < right{
        mid = (left + right) / 2

        if nums[mid] < target{
            left = mid + 1
        } else if nums[mid] > target{
            right = mid - 1
        } else if nums[mid] == target{
            isTargetFound = true
            break
        }
    }

    if !isTargetFound{
        return []int{-1, -1}
    }

    midCopy := mid

    if midCopy == -1{
        return []int{-1,-1}
    } else {
        left, right = midCopy, midCopy

        for left >= 0 && nums[left] == target{
            left--
        }

        for right <= len(nums) - 1 && nums[right] == target{
            right++
        }
    }

    return []int{left+1, right-1}
}

func main(){
	searchRange([]int{1,4}, 4)
}
