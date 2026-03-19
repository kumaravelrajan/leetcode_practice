package main

import (
	"fmt"
	"math"
)

func isHappy(n int) bool {
    // Find the different digits in the number n
	lookup := make(map[int]bool)

	for true{
		c_sum := 0

		n_digits := []int{}

		// Store current n in lookup so that later we can refer to this table easily for repeated values.
		lookup[n] = true

		for n > 0{
			last_digit := n%10
			n = n/10
			n_digits = append([]int{last_digit}, n_digits...)
		}

		for _, digit := range n_digits{
			c_sum += (digit*digit)
		}

		if c_sum == 1{
			return true
		} else if _, status := lookup[c_sum]; status{
			return false
		}

		// get digits from c_sum
		c_sum_digits := []int{}

		for c_sum > 0{
			c_sum_last_digit := c_sum % 10
			c_sum = c_sum / 10
			c_sum_digits = append([]int{c_sum_last_digit}, c_sum_digits...)
		}

		// Convert c_sum digits into n with base 10
		n = 0
		j := 0
		for i := len(c_sum_digits) - 1; i >= 0; i--{
			n += c_sum_digits[i] * int(math.Pow(10, float64(j)))
			j++
		}
	}

	return false
}

func main(){

	fmt.Println(isHappy(19))
	fmt.Println(isHappy(2))	

}
