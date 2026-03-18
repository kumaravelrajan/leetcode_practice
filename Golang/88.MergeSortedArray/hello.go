package main

import (
	"slices"
	"fmt"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	// Sanity checking
	if n == 0{
		return
	} else if m == 0{
		index_nums1, index_nums2 := 0, 0

		for ; index_nums2 < len(nums2); 
		
		index_nums1, index_nums2 = index_nums1 + 1, index_nums2 + 1{
			nums1[index_nums1] = nums2[index_nums2]
		}
		return
	}
	
	index_nums2 := 0

	for index_nums1, ctnums1 := range nums1{
		ctnums2 := nums2[index_nums2]

		if ctnums1 > ctnums2 && index_nums1 < m{
			nums1[index_nums1], nums2[index_nums2] = nums2[index_nums2], nums1[index_nums1]
			
			slices.Sort(nums2)
		}

		if index_nums1 >= m{
			nums1[index_nums1] = nums2[index_nums2]
			index_nums2++
		}
	}
}

func main() {

	nums1 := []int{1,2,3,0,0,0}
	merge(nums1, 3, []int{2,5,6}, 3)
	fmt.Println(nums1)

	nums1 = []int{1}
	merge(nums1, 1, []int{}, 0)
	fmt.Println(nums1)

	nums1 = []int{0}
	merge(nums1, 0, []int{1}, 1)
	fmt.Println(nums1)


}
