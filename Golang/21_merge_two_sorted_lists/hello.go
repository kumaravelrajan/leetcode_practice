package main

import (
	"fmt"
)

// Definition for singly-linked list.
type ListNode struct {
Val int
Next *ListNode
}

func forwardlist1orlist2(list1or2 **ListNode, ct_list3 **ListNode, prev_list3 **ListNode){
	(*ct_list3).Val = (*list1or2).Val
	(*ct_list3).Next = new(ListNode)
	*prev_list3 = *ct_list3
	*ct_list3 = (*ct_list3).Next

	*list1or2 = (*list1or2).Next
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    
	// Sanity checking
	if list1 != nil && list2 == nil{
		return list1
	} else if list2 != nil && list1 == nil{
		return list2		
	} else if list1 == nil && list2 == nil{
		return nil
	}

	var ct_list1 *ListNode = list1
	var ct_list2 *ListNode = list2
	var list3 *ListNode = new(ListNode)
	var ct_list3 *ListNode = list3
	var prev_list3 *ListNode

	for ct_list1 != nil || ct_list2 != nil{

		if ct_list1 != nil && ct_list2 != nil{

			if ct_list1.Val < ct_list2.Val{

			forwardlist1orlist2(&ct_list1, &ct_list3, &prev_list3)

			} else if ct_list2.Val < ct_list1.Val{

				forwardlist1orlist2(&ct_list2, &ct_list3, &prev_list3)


			} else if ct_list1.Val == ct_list2.Val {

				forwardlist1orlist2(&ct_list1, &ct_list3, &prev_list3)

			}
		} else {

			if ct_list1 != nil && ct_list2 == nil {

				forwardlist1orlist2(&ct_list1, &ct_list3, &prev_list3)

			} else{

				forwardlist1orlist2(&ct_list2, &ct_list3, &prev_list3)

			}

		}
		 
	}

	prev_list3.Next = nil

	return list3
}

func main(){

	var list1 *ListNode = &ListNode{Val: 1}
	list1.Next = &ListNode{Val: 2}
	list1.Next.Next = &ListNode{Val: 4}

	var list2 *ListNode = &ListNode{Val: 1}
	list2.Next = &ListNode{Val: 3}
	list2.Next.Next = &ListNode{Val: 4}

	list3 := mergeTwoLists(list1, list2)

	fmt.Println(list3.Val)

}
