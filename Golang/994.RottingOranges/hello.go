package main

func orangesRotting(grid [][]int) int {
    // Scan the grid once to record all indices for 2s and count of 1s.

    if len(grid) == 0 || (len(grid) == 1 && len(grid[0]) == 0){
        return -1
    }

    countOf1 := 0
    q := make([][]int, 0, len(grid))
    numOfMinutes := 0

    for i := 0; i < len(grid) ; i++{
        for j := 0; j < len(grid[i]); j++{
            if grid[i][j] == 1{
                countOf1++
            } else if grid[i][j] == 2{
                q = append(q, []int{i,j})
            }
        }
    }

    // Start with q and branch out as far as possible. For every fresh orange turned rotten in this process reduce countOf1 by 1. If at the end there are no elements left in q but countOf1 is more than 0 then return -1. Else, return the number of minutes.

    for len(q) > 0 && countOf1 > 0{

        i := 0
        n := len(q)

        for i < n{
            q = makeFreshOrangesRottenForI(grid, q, q[i], &countOf1)
			i++
        }

        q = q[n:]
        numOfMinutes++
    }

    if countOf1 > 0 {
        return -1
    } else {
        return numOfMinutes
    }
}

func makeFreshOrangesRottenForI(grid [][]int, q [][]int, currRottenIndex []int, countOf1 *int) ([][]int){
    i, j := currRottenIndex[0], currRottenIndex[1]

    if i+1 < len(grid) && grid[i+1][j] == 1 {
        grid[i+1][j] = 2
        *countOf1--
        q = append(q, []int{i + 1, j})
    }

    if i - 1 >= 0 && grid[i-1][j] == 1{
        grid[i-1][j] = 2
        *countOf1--
        q = append(q, []int{i - 1, j})
    }

    if j + 1 < len(grid[i]) && grid[i][j + 1] == 1{
        grid[i][j + 1] = 2
        *countOf1--
        q = append(q, []int{i, j + 1})
    }

    if j - 1 >= 0 && grid[i][j - 1] == 1{
        grid[i][j - 1] = 2
        *countOf1--
        q = append(q, []int{i, j - 1})
    }

    return q
}

func main(){
	orangesRotting([][]int{{2,1,1},{1,1,0},{0,1,1}})
}
