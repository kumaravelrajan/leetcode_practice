package main

func shortestPathBinaryMatrix(grid [][]int) int {
    // Sanity checks
    if len(grid) == 0{
        return -1
    } else if len(grid) > 0 && len(grid[0]) == 0 {
        return -1
    } else if grid[0][0] == 1 || grid[len(grid)-1][len(grid)-1] == 1 {
        return -1
    }

    q := make([][]int, 0, len(grid)*len(grid[0]))
    q = append(q, []int{0, 0})
    isDestReached := false
    result := 1

    for len(q) > 0 && !isDestReached{
        n := len(q)
        q = findAdj0(grid, q, &isDestReached)
        q = q[n:]
        result++
    }

    if isDestReached{
        return result
    } else {
        return -1
    }
}

func findAdj0(grid [][]int, q [][]int, isDestReached *bool)([][]int){
    for _, currIJArr := range q{
        i := currIJArr[0]
        j := currIJArr[1]

        // Check next and before row
        if i + 1 < len(grid) && grid[i + 1][j] == 0{
            q = append(q, []int{i+1,j})
            grid[i+1][j] = 2

            if isCurrIndexDest(grid, i+1, j){
                *isDestReached = true
                return q
            }
        }

        if i - 1 >= 0 && grid[i - 1][j] == 0{
            q = append(q, []int{i - 1,j})
            grid[i-1][j] = 2

            if isCurrIndexDest(grid, i - 1, j){
                *isDestReached = true
                return q
            }
        }

        // Check next and before column
        if j + 1 < len(grid[0]) && grid[i][j + 1] == 0{
            q = append(q, []int{i,j + 1})
            grid[i][j+1] = 2

            if isCurrIndexDest(grid, i, j + 1){
                *isDestReached = true
                return q
            }
        }

        if j - 1 >= 0 && grid[i][j - 1] == 0{
            q = append(q, []int{i,j - 1})
            grid[i][j-1] = 2

            if isCurrIndexDest(grid, i, j - 1){
                *isDestReached = true
                return q
            }
        }

        // Check 4 diagonals

        // Top left
        if i - 1 >= 0 && j - 1 >= 0 && grid[i - 1][j - 1] == 0{
            q = append(q, []int{i - 1,j - 1})
            grid[i-1][j-1] = 2

            if isCurrIndexDest(grid, i - 1, j - 1){
                *isDestReached = true
                return q
            }
        }

        // Top right
        if i - 1 >= 0 && j + 1 < len(grid[0]) && grid[i - 1][j + 1] == 0{
            q = append(q, []int{i - 1,j + 1})
            grid[i-1][j+1] = 2

            if isCurrIndexDest(grid, i - 1, j + 1){
                *isDestReached = true
                return q
            }
        }

        // Bottom right
        if i + 1 < len(grid) && j + 1 < len(grid[0]) && grid[i + 1][j + 1] == 0{
            q = append(q, []int{i + 1,j + 1})
            grid[i+1][j+1] = 2

            if isCurrIndexDest(grid, i + 1, j + 1){
                *isDestReached = true
                return q
            }
        }

        // Bottom left
        if i + 1 < len(grid) && j - 1 >= 0 && grid[i + 1][j - 1] == 0{
            q = append(q, []int{i + 1,j - 1})
            grid[i+1][j-1] = 2

            if isCurrIndexDest(grid, i + 1, j - 1){
                *isDestReached = true
                return q
            }
        }
    }
    return q
}

func isCurrIndexDest(grid [][]int, i int, j int) bool {
    if i == len(grid) - 1 && j == len(grid[0]) - 1{
        return true
    } else {
        return false
    }
}

func main(){
	shortestPathBinaryMatrix([][]int{{0,1},{1,0}})
}