package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {

	var x_str string = strconv.Itoa(x) 

	if len(x_str) == 0{
		return false
	}

	start_index := 0
	end_index := len(x_str) - 1

	for range x_str{

		if start_index == end_index {
			return true
		}

		if x_str[start_index] == x_str[end_index]{

			start_index++
			end_index--
			continue

		} else {

			return false

		}
	}
    
	return true
}

func main(){

	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))

}