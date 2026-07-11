package main

/**
 * Definition for a binary tree node.*/
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
    result := -1

    var inorder func(*TreeNode)

    inorder = func (root *TreeNode){
        if root == nil{
            return
        }

        // Keep drilling down the left subtree until you reach end
        inorder(root.Left)

        k--

        if k == 0{
            // Node found
            result = root.Val
            return
        }

        if k < 0 {
            return
        }

        inorder(root.Right)

    }

    return result
}

func main(){
	
}
