class Node:
    def __init__(self, info):
        self.info = info  
        self.left = None  
        self.right = None 
        self.level = None 

    def __str__(self):
        return str(self.info) 

def preOrder(root):
    if root == None:
        return
    print (root.info, end=" ")
    preOrder(root.left)
    preOrder(root.right)
    
class BinarySearchTree:
    def __init__(self): 
        self.root = None

#Node is defined as
#self.left (the left child of the node)
#self.right (the right child of the node)
#self.info (the value of the node)

    def insert(self, val):
        #Enter you code here.
        
        node_to_add = Node(val)
        
        # Tree completely empty
        if self.root is None:
            self.root = node_to_add
            
        else:
            
            ct_node = self.root
            
            while True:
                if node_to_add.info < ct_node.info:
                    if ct_node.left:
                        ct_node = ct_node.left
                        
                    else:
                        ct_node.left = node_to_add
                        
                        break
                        
                elif node_to_add.info > ct_node.info:
                    if ct_node.right:
                        ct_node = ct_node.right
                        
                    else:
                        ct_node.right = node_to_add
                        break
                        
            return self.root
                        
                
                        
                                
            
        

tree = BinarySearchTree()
# t = int(input())

# arr = list(map(int, input().split()))

arr = [4, 2, 1, 3, 7, 6]

for elem in arr:
    tree.insert(elem)

preOrder(tree.root)
