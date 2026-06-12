package main

import "fmt"

func largestRectangleArea(heights []int) int {
    pse := make([]int, len(heights))
    nse := make([]int, len(heights))
    stack := make([]int, 0, len(heights))

    for i := 0; i < len(pse); i++{
        pse[i] = -1
        nse[i] = -1
    }

    for i, currHeight := range heights{
        for len(stack) > 0 && heights[stack[len(stack)-1]] > currHeight{
            nse[stack[len(stack)-1]] = i
            stack = stack[: len(stack)-1]
        }
        stack = append(stack, i)
    }

    stack = make([]int, 0, len(heights))

    for i := len(heights)-1; i >= 0; i--{
        currHeight := heights[i]
        for len(stack) > 0 && heights[stack[len(stack)-1]] > currHeight{
            pse[stack[len(stack)-1]] = i
            stack = stack[: len(stack) - 1]
        }

        stack = append(stack, i)
    }

    fmt.Println("NSE = ", nse)
    fmt.Println("PSE = ", pse)

    maxArea := -1
    for i := 0; i < len(heights); i++{
		currArea := -1
		width := 0

		if pse[i] == -1 {
			width += i + 1
		} else {
			width += i - pse[i]
		}

		if nse[i] == -1 {
			width += len(heights) - 1 - i
		} else {
			width += nse[i] - i - 1
		}

		currArea = heights[i] * width

		if currArea > maxArea{
			maxArea = currArea
		}
    }

    return maxArea
}

func main(){

	largestRectangleArea([]int{0,9})

}
