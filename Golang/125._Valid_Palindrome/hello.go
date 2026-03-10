package main

import (
	"fmt"
	// "unicode/utf8"
	"unicode"
)

func isPalindrome(s string) bool {

	var s_rune_slice []rune = []rune(s)

	// fmt.Println(len(s), "; ", len(s_rune_slice))

	for i, j := 0, len(s_rune_slice) - 1; i < j; i, j = i+1, j-1{
		for (!unicode.IsLetter(s_rune_slice[i]) && !unicode.IsDigit(s_rune_slice[i])) && (i < j){
			i++
		}

		for (!unicode.IsLetter(s_rune_slice[j]) && !unicode.IsDigit(s_rune_slice[j])) && (i < j){
			j--
		}

		if unicode.ToLower(s_rune_slice[i]) == unicode.ToLower(s_rune_slice[j]){
			continue
		} else {
			return false
		}
	}

	return true
}

func main(){
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
	fmt.Println(isPalindrome(" "))
}