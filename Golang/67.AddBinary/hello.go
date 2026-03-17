package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	carry := 0
	
	// Use a byte slice to efficiently build the result string
	var result []byte

	// Loop until both strings are exhausted and no carry remains
	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		
		// Add the bit from string 'a' if available
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		
		// Add the bit from string 'b' if available
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}

		// Calculate the new bit (sum % 2) and append it
		result = append(result, byte((sum%2)+'0'))
		
		// Calculate the new carry (sum / 2)
		carry = sum / 2
	}

	// The result is built backwards (least significant bit first), so we must reverse it
	for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
		result[left], result[right] = result[right], result[left]
	}

	return string(result)
}

func main(){

	fmt.Println(addBinary("11", "1"))
	fmt.Println(addBinary("1010", "1011"))
}
