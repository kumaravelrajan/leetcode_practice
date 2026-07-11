package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0
    
    // We define a helper function to calculate height
    var maxDepth func(*TreeNode) int
    maxDepth = func(node *TreeNode) int {
        if node == nil {
            return 0 // Base case: an empty node has a height of 0
        }
        
        // Recursively get the height of left and right subtrees
        leftHeight := maxDepth(node.Left)
        rightHeight := maxDepth(node.Right)
        
        // The diameter passing through the CURRENT node is leftHeight + rightHeight.
        // We check if this is the largest diameter we've seen so far.
        if leftHeight + rightHeight > maxDiameter {
            maxDiameter = leftHeight + rightHeight
        }
        
        // Return the height of this subtree to the parent call
        return max(leftHeight, rightHeight) + 1
    }
    
    maxDepth(root)
    return maxDiameter
}

func main() {

	one := TreeNode{Val: 1}
	two := TreeNode{Val: 2}
	three := TreeNode{Val: 3}
	four := TreeNode{Val: 4}
	five := TreeNode{Val: 5}

	one.Left = &two
	one.Right = &three
	two.Left = &four
	two.Right = &five

	fmt.Println(diameterOfBinaryTree(&one))

}
