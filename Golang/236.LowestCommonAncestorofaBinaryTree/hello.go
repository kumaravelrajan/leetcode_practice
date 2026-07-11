package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    // Base case: if we hit a nil node, or if we find either p or q, 
    // we return the current node up.
    if root == nil || root == p || root == q {
        return root
    }

    // Search the left and right subtrees
    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)

    // If both left and right returned a non-nil node, it means 'p' was found 
    // in one subtree and 'q' was found in the other. This root IS the LCA!
    if left != nil && right != nil {
        return root
    }

    // If only one side returned a node, pass it up (it means both p and q 
    // are down that single path, or we've only found one of them so far)
    if left != nil {
        return left
    }
    return right
}

func main(){

	three := TreeNode{Val: 3}
	five := TreeNode{Val: 5}
	one := TreeNode{Val: 1}
	six := TreeNode{Val: 6}
	two := TreeNode{Val: 2}
	seven := TreeNode{Val: 7}
	four := TreeNode{Val: 4}
	zero := TreeNode{Val:0}
	eight := TreeNode{Val:8}

	three.Left = &five
	three.Right = &one
	five.Left = &six
	five.Right = &two
	two.Left = &seven
	two.Right = &four
	one.Left = &zero
	one.Right = &eight

	print(lowestCommonAncestor(&three, &five, &four).Val)

}
