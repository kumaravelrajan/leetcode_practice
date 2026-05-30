package main

func threeSum(nums []int) [][]int {
        
	triplets := [][]int{}

		slices.Sort(nums)

		for index, currNum := range nums{
			left, right := index + 1, len(nums)-1

							if nums[index] == nums[index+1]{
										continue
												}

														for left < len(nums) - 1 && left < right && nums[left] == nums[left + 1]{
																	left++
																			}

																					for right > 0 && left < right && nums[right] == nums[right-1]{
																								right--
																										}

																												if currNum + nums[left] + nums[right] == 0{
																															currTriplet := []int{currNum, nums[left], nums[right]}
																																		slices.Sort(currTriplet)

																																					isCurrTripletAlreadySaved := false
																																								for _, savedTriplet := range triplets{
																																												if slices.Equal(savedTriplet, currTriplet){
																																																	isCurrTripletAlreadySaved = true
																																																					}
																																																								}

																																																											if !isCurrTripletAlreadySaved{
																																																															triplets = append(triplets, currTriplet)
																																																																		}
																																																																				} else if currNum + nums[left] + nums[right] < 0{
																																																																							left++
																																																																									} else if currNum + nums[left] + nums[right] > 0{
																																																																												right--
																																																																														}
																																																																															}

																																																																																return triplets
																																																																																}
}

func main(){

}
