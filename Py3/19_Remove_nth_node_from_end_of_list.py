from typing import Optional

class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

def print_ll(head: ListNode):
    vals = []
    while head != None:
        vals.append(str(head.val))
        head = head.next

    return " ".join(vals)

class Solution:
    def removeNthFromEnd(self, head: Optional[ListNode], n: int) -> Optional[ListNode]:

        node, pre = head, head
        ll_index = 0

        # If Head to be removed return None
        if head.next is None and n >= 1:
            return None

        while node is not None:
            node = node.next

            if ll_index > n:
                pre = pre.next

            ll_index += 1

        if n == 1: # Tail needs to be removed
            pre.next = None
            print_ll(head)
            return head
        
        if n == ll_index: # Head needs to be removed
            head = head.next
            print_ll(head)
            return head
            
        pre.next = pre.next.next
        print_ll(head)
        return head 

        

                    

one = ListNode(1)
two = ListNode(2)
# three = ListNode(3)
# four = ListNode(4)
# five = ListNode(5)

one.next = two
# two.next = three
# three.next = four
# four.next = five

sol = Solution()
sol.removeNthFromEnd(one, 2)