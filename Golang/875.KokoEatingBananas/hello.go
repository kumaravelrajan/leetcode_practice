package main

import ("slices")

func minEatingSpeed(piles []int, h int) int {
    lSpeed, rSpeed := slices.Min(piles), slices.Max(piles)
    resSpeed := 0

    for lSpeed <= rSpeed{
        midSpeed := (lSpeed + rSpeed) / 2
        currHours := 0

        for _, currPile := range piles{
            if midSpeed >= currPile{
                currHours++
            } else {
                // Ceiling function
                currHours += (currPile + midSpeed - 1) / midSpeed 
            }
        }

        if currHours <= h {
            
            resSpeed = midSpeed
            rSpeed = midSpeed - 1
        } else if currHours > h {
            lSpeed = midSpeed + 1
        }
    }

    return resSpeed
}

func main(){
	minEatingSpeed([]int{312884470}, 312884469)
}
