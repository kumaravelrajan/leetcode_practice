package main

import (
	"fmt"
	"slices"
)

func majorityElement(nums []int) int {
	slices.Sort(nums)
	p_highfreq, p_num, c_highfreq, c_num := 0, 0, 0, 0

	for _, num := range nums {
		if c_num != num {
			c_num = num
			c_highfreq = 0
			c_highfreq++
		} else if c_num == num {
			c_highfreq++
		}

		if c_highfreq > p_highfreq {
			p_highfreq = c_highfreq
			p_num = c_num
		}
	}

	return p_num
}

func main() {

	fmt.Println(majorityElement([]int{3, 2, 3}))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))

}
