from typing import Optional

# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution:
    def swapPairs(self, head: Optional[ListNode]) -> Optional[ListNode]:

        if not head:
            return None
        
        if head.next is None:
            return head
        
        current_node = head
        next_node = None
        prev_node = None

        if current_node.next is not None:
            next_node = current_node.next

            current_node.next = next_node.next
            next_node.next = current_node
            head = next_node
            current_node, next_node = next_node, current_node
            
            if current_node.next.next is None or next_node.next.next is None:
                return head
            else:
                prev_node = next_node
                current_node = current_node.next.next
                next_node = next_node.next.next

        while True:
            prev_node.next = current_node.next
            current_node.next = next_node.next
            next_node.next = current_node
            current_node, next_node = next_node, current_node

            if current_node.next.next is None or next_node.next.next is None:
                return head
            else:
                prev_node = next_node
                current_node = current_node.next.next
                next_node = next_node.next.next
                        
one = ListNode(1)
two = ListNode(2)
three = ListNode(3)
four = ListNode(4)

one.next = two
two.next = three
three.next = four

sol = Solution()
print(sol.swapPairs(one))