package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {

	if len(s) != len(t){
		return false
	}

	char_mapping := make(map[byte]byte)
	destination_chars_already_mapped := make(map[byte]bool)

	// Only ASCII characters present here. Therefore, no need for UNICODE considerations. 
	s_i, t_j := 0, 0

	for ; s_i < len(s) && t_j < len(t); s_i, t_j = s_i + 1, t_j + 1{
		curr_s := s[s_i]
		curr_t := t[t_j]

		if curr_s != curr_t{
			_, status_s_char_mapped := char_mapping[curr_s]
			_, status_t_char_mapped := destination_chars_already_mapped[curr_t]
			
			if (!status_s_char_mapped && !status_t_char_mapped){
				char_mapping[curr_s] = curr_t
				destination_chars_already_mapped[curr_t] = true
			} else if (curr_t != char_mapping[curr_s]){
				return false
			}
		} else if curr_s == curr_t{
			_, status_s_char_mapped := char_mapping[curr_s]
			_, status_t_char_mapped := destination_chars_already_mapped[curr_t]

			if (!status_s_char_mapped && !status_t_char_mapped){
				char_mapping[curr_s] = curr_t
				destination_chars_already_mapped[curr_t] = true
			} else if (curr_t != char_mapping[curr_s]){
				return false
			}
		}
	}
    return true
}

func main(){
	fmt.Println(isIsomorphic("egg", "add"))
	fmt.Println(isIsomorphic("f11", "b23"))
	fmt.Println(isIsomorphic("paper", "title"))
}
