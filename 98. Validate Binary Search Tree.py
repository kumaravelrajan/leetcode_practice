from typing import Optional

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def isValidBST(self, root: TreeNode) -> bool:
        def validate(node, low=-float('inf'), high=float('inf')):
            # An empty tree is a valid BST

            if not node:
                return True
            
            # The current node's value must be strictly within the low and high bounds
            if node.val <= low or node.val >= high:
                return False
            
            # Recursively validate the left and right subtrees
            # Left subtree: max value allowed becomes the current node's value
            # Right subtree: min value allowed becomes the current node's value
            return (validate(node.left, low, node.val) and 
                    validate(node.right, node.val, high))
        
        return validate(root)

        
# Tree 1
root1 = TreeNode(2)
root1.left = TreeNode(1)
root1.right = TreeNode(3)

# Tree 2
root2 = TreeNode(5)
root2.left = TreeNode(1)
root2.right = TreeNode(4)
root2.right.left = TreeNode(3)
root2.right.right = TreeNode(6)

# Tree 3
root3 = TreeNode(5)
root3.left = TreeNode(4)
root3.right = TreeNode(6)
root3.right.left = TreeNode(3)
root3.right.right = TreeNode(7)

sol = Solution()

print(sol.isValidBST(root1))
print(sol.isValidBST(root2))
print(sol.isValidBST(root3))


