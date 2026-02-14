class Node:
    def __init__(self, info): 
        self.info = info  
        self.left = None  
        self.right = None 
        self.level = None 

    def __str__(self):
        return str(self.info) 

class BinarySearchTree:
    def __init__(self): 
        self.root = None

    def create(self, val):  
        if self.root == None:
            self.root = Node(val)
        else:
            current = self.root
         
            while True:
                if val < current.info:
                    if current.left:
                        current = current.left
                    else:
                        current.left = Node(val)
                        break
                elif val > current.info:
                    if current.right:
                        current = current.right
                    else:
                        current.right = Node(val)
                        break
                else:
                    break

def height(root):
    # Base case: If the tree is empty or we reach a leaf's child
    if root is None:
        return -1
    
    # Recursively calculate the height of the left and right subtrees
    left_height = height(root.left)
    right_height = height(root.right)
    
    # The height of the current node is 1 + the maximum of the two children
    return 1 + max(left_height, right_height)

tree = BinarySearchTree()
# t = int(input())

arr = [3, 5, 2, 1, 4, 6, 7]

for i in arr:
    tree.create(i)

print(height(tree.root))