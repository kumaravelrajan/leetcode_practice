package main

import "fmt"

func checkInclusion(s1 string, s2 string) bool {

    if len(s1) > len(s2){
        return false
    }

    s1Lookup := make([]int, 26)
    s2Lookup := make([]int, 26)

    for _, char := range s1{
        s1Lookup[char - 'a']++
    }

    for i:= 0; i < len(s1); i++{
        s2Lookup[s2[i] - 'a']++   
    }

    s2LeftIndex := 0

    if compareSlices(s2Lookup, s1Lookup){
        return true
    } else {
        s2RightIndex := len(s1)

        for s2RightIndex < len(s2){
            
            s2Lookup[s2[s2LeftIndex]-'a']--
            s2Lookup[s2[s2RightIndex]-'a']++

            if compareSlices(s1Lookup, s2Lookup){
                return true
            } 

            s2LeftIndex++            
            s2RightIndex++
        }
    }

    return false
}

func compareSlices(slice1 []int, slice2 []int) bool{
    for i := 0; i < len(slice1); i++{
        if slice1[i] != slice2[i]{
            return false
        }
    }
    return true
}

func main(){
	fmt.Println(checkInclusion("ab", "eidbaooo"))
	fmt.Println(checkInclusion("adc", "dcda"))
}