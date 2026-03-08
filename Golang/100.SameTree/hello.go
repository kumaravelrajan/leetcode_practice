package main

import "fmt"

// Definition for a binary tree node.
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
    
	// base case
	if p != nil && q != nil{

		var is_p_child_node bool = p.Left == nil && p.Right == nil
		var is_q_child_node bool = q.Left == nil && q.Right == nil
		
		if is_p_child_node && is_q_child_node {

			if p.Val != q.Val{
				return false
			} else {
				return true
			}
		} else if is_p_child_node != is_q_child_node{

			// One is a leaf node while other is not. Hence, trees are not same.
			return false
		}
	} else if (p == nil) != (q == nil) {

		// One is nil while other is not. Clearly, trees are different.
		return false
	} else if (p == nil) && (q == nil){
		
		return true
	}
	

	// Recursion 
	var result bool

	if p.Val == q.Val{

		result = isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		
		if !result{
			return false
		}
	} else {
		return false
	}

	return result
}

func main(){

	p := TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	q := TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}

	fmt.Println(isSameTree(&p, &q))

	r := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	s := &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}

	fmt.Println(isSameTree(r, s))

}
