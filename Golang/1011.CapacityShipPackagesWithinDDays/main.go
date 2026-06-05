package main

import ("fmt"
"math")

func shipWithinDays(weights []int, days int) int {

    // Maximum possible capacity = sum of all weights (all packages delivered in 1 day)
    // Minimum possible capacity = weight of heaviest package (all packages delivered in n days)

    minCapacity, maxCapacity := math.MinInt, 0

    // O(N)
    for _, currWeight := range weights{
        if currWeight > minCapacity {
            minCapacity = currWeight
        }

        maxCapacity += currWeight
    }

    fmt.Println("maxCapacity = ", maxCapacity, "; minCapacity = ", minCapacity)

    lCapacity, rCapacity := minCapacity, maxCapacity
    resCapacity := math.MaxInt

    for lCapacity <= rCapacity{
        midCapacity := (lCapacity + rCapacity)/2

        numOfDays := 0
		sumOfWeights := 0

        for _, currWeight := range weights{
            if sumOfWeights + currWeight <= midCapacity{
				sumOfWeights += currWeight
			} else {
				numOfDays++
				sumOfWeights = currWeight
			}
        }

		if sumOfWeights > 0 {
			numOfDays++
		}

        if numOfDays < days {
            rCapacity = midCapacity - 1
            
            if midCapacity < resCapacity {
                resCapacity = midCapacity
            }
        } else if numOfDays > days {
            lCapacity = midCapacity + 1
        } else {
            if midCapacity < resCapacity{
                resCapacity = midCapacity
				rCapacity = midCapacity - 1
            }
        }

    }

    return resCapacity
    
}

func main(){

	shipWithinDays([]int{3,2,2,4,1,4}, 3)

}
