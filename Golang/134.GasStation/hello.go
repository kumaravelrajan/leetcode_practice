package main

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {

    if len(gas) == 1 {
        if gas[0] >= cost[0]{
			return 0
		} else {
			return -1
		}
    }

    availGas := 0

    for i := 0; i < len(gas); i++{
        if gas[i] < cost[i]{
            continue
        } else {
            availGas = gas[i]
			isTripStarted := false

			circleStartI := i

            for j := i; availGas >= cost[j]; j = (j + 1) % len(gas){

                nextJ := (j + 1) % len(gas)

                availGas = availGas - cost[j] + gas[nextJ]

				if isTripStarted {
					if nextJ == circleStartI {
						return circleStartI
					}
				} else {
					isTripStarted = true
				}
            }
        }
    }

    return -1
}

func main(){

	fmt.Println(canCompleteCircuit([]int{1,2,3,4,5}, []int{3,4,5,1,2}))
	fmt.Println(canCompleteCircuit([]int{2,3,4}, []int{3,4,3}))

}
