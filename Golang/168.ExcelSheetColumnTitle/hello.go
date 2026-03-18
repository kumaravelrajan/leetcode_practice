package main

import "fmt"

func convertToTitle(columnNumber int) string {
    base_char := 'A'
	
	var result []byte

	for columnNumber > 0{
		columnNumber--

		result = append([]byte{byte(base_char) + byte(columnNumber%26)}, result...)

		columnNumber /= 26
	}

	return string(result)

}

func main(){
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(701))
}
