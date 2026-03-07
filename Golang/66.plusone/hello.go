package main

import (
	"fmt"
	// "string"
)

func plusOne(digits []int) []int {
	
	if len(digits) == 0{

		return []int{}

	} else if len(digits) == 1 {

		if digits[0] < 9 {

			digits[0]++
			return digits

		} else if digits[0] == 9 {

			digits[0] = 0
			digits = append([]int{1}, digits...)
			return digits

		}

	}

	for i := len(digits) - 1; i >= 0 ; i--{

		if i == len(digits) - 1 {

			if digits[i] < 9{

				digits[i]++
				return digits

			} else if digits[i] == 9{

				// is_carry_required = true
				digits[i] = 0

			}

		} else if i != 0{

			if digits[i] < 9{

				digits[i]++
				return digits

			} else if digits[i] == 9{

				digits[i] = 0

			}

		} else if i == 0{

			if digits[i] < 9 {

				digits[i]++
				return digits

			} else if digits[i] == 9 {

				digits[0] = 0
				digits = append([]int{1}, digits...)
				return digits

			}

		}
		
	}

	return []int{}

}

func main(){
	fmt.Println(plusOne([]int{1,2,3}))
	fmt.Println(plusOne([]int{4,3,2,1}))
	fmt.Println(plusOne([]int{9}))
	fmt.Println(plusOne([]int{9,9,9,9}))
}

