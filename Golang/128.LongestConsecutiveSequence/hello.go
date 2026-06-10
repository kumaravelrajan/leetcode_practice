package main

func longestConsecutive(nums []int) int {

    lookup := make(map[int]bool)
    result := 0

    for _, num := range nums{
        lookup[num] = true
    }

    for num, isPresentInLookup := range lookup{

        numCopy := num
        currSeq := 0
        
        if isPresentInLookup{
            if _, exists := lookup[numCopy-1]; exists{                

                for exists{
                    lookup[numCopy] = false
                    numCopy--

                    _, exists = lookup[numCopy-1]
                }

                exists = true

                for exists{
                    currSeq++
                    numCopy++
                    _, exists = lookup[numCopy]
                }
            } else {
				currSeq++
			}
        } 

        if currSeq > result{
            result = currSeq
        }
        
    }

    return result
}

func main(){
	longestConsecutive([]int{0, 0})
}
