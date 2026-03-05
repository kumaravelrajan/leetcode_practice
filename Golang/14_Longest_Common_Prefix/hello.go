package main

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0{
		return ""
	}
	
	if len(strs) == 1{
		return strs[0]
	}

	var shortest_string string = strs[0]
	var longest_common_prefix string

	// 1st pass: Find shortest string in the array first.
	for _, str := range strs[1:]{
		
		if len(shortest_string) > len(str){
			shortest_string = str
		}
	}

	// 2nd pass: Deduct characters from the shortest string to fit in the longest prefix possible.

	longest_common_prefix = shortest_string

	for _, str := range strs{
		
		// Run loop until longest_common_prefix contained in str. If longest_common_prefix becomes "" return "".
		for !strings.HasPrefix(str, longest_common_prefix){
			longest_common_prefix = longest_common_prefix[:len(longest_common_prefix)-1]

			if longest_common_prefix == ""{
				return ""
			}
		}

	}

	return longest_common_prefix
}

func main(){

	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog","racecar","car"}))

}
