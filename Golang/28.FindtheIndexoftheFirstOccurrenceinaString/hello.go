package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
    n_len := len(needle)

	for i := 0; (i + n_len - 1) < len(haystack) ; i++{
		if haystack[i:i+n_len] == needle{
			return i
		}
	}

	return -1
}

func main(){

	fmt.Println(strStr("sadbutsad", "sad"))
	fmt.Println(strStr("leetcode", "leeto"))


}
