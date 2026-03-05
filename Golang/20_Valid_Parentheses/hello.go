package main

import (
	"fmt"
)

func isValid(s string) bool {

	if len(s) % 2 != 0{
		return false
	}

	var closing_brackets = make(map[rune]rune)
	closing_brackets[']'] = '['
	closing_brackets['}'] = '{'
	closing_brackets[')'] = '('

	var stack = []rune{}

	for _, bracket := range s{

		if _, ok := closing_brackets[bracket]; !ok{
			//Opening bracket found; add to stack

			stack = append(stack, bracket)
		} else {
			// Closing bracket found. Check if last element of stack is compatible.
			var last_in_stack rune

			if len(stack) > 0{
				last_in_stack = stack[len(stack) - 1]
			} else {
				return false
			}
			
			if closing_brackets[bracket] != last_in_stack{

				return false

			} 

			stack = stack[:len(stack) - 1]
		}
		
	}

	if len(stack) == 0{
		return true
	} else {
		return false
	}
}

func main(){

fmt.Println(isValid("()"))
fmt.Println(isValid("()[]{}"))
fmt.Println(isValid("(]"))
fmt.Println(isValid("([])"))
}