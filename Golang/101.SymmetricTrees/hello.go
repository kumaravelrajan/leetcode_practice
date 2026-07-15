package main

import "fmt"

type TreeNode struct{
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
    
    var findOutput func (left *TreeNode, right *TreeNode) bool

    findOutput = func (left *TreeNode, right *TreeNode) bool {
        // Base case conditions
        if left == nil && right == nil {
            return true
        } else if left != nil && right == nil{
            return false
        } else if left == nil && right != nil{
            return false
        } 

		boolVal := left.Val == right.Val

		if boolVal{
            boolVal = findOutput(left.Left, right.Right)

			if boolVal{
		        boolVal = findOutput(left.Right, right.Left)

				if !boolVal{
					return false
				}
			} else {
				return false
			}
        } else {
			return false
		}

		return true
    }

    return findOutput(root.Left, root.Right)
}

func main(){
	one := TreeNode{Val: 1}
	two1 := TreeNode{Val: 2}
	three1 := TreeNode{Val: 3}
	four1 := TreeNode{Val: 4}
	two2 := TreeNode{Val: 2}
	three2 := TreeNode{Val: 3}
	four2 := TreeNode{Val: 4}

	one.Left = &two1
	one.Right = &two2
	two1.Left = &three1
	two1.Right = &four1
	two2.Left = &four2
	two2.Right = &three2

	fmt.Println(isSymmetric(&one))

}
