package main

import (
	"fmt"
	"unicode/utf8"
)

func lengthOfLastWord(s string) int {
    lastwordlen := 0

	for len(s) > 0{
		r, size := utf8.DecodeLastRuneInString(s)

		if size > 0 && r != ' '{
			lastwordlen++
		}

		if lastwordlen > 0 && r == ' '{
			break
		}

		s = s[: len(s) - size]
	}

	return lastwordlen
}

func main(){

	fmt.Println(lengthOfLastWord("Hello World"))
	fmt.Println(lengthOfLastWord("   fly me   to   the moon  "))
	fmt.Println(lengthOfLastWord("luffy is still joyboy"))
	fmt.Println(lengthOfLastWord(" 兴 "))

}
