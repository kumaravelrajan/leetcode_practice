class Solution:
    def mergeTwoLists(self, list1: Optional[ListNode], list2: Optional[ListNode]) -> Optional[ListNode]:

        list1_current_node = list1
        list2_current_node = list2

        while list1_current_node != None or list2_current_node != None:
            if list1_current_node.value < list2_current_node.value