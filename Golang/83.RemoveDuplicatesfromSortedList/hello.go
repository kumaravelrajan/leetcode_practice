package main

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

func printList(head *ListNode){
	c_node := head

	for c_node != nil{
		fmt.Println(c_node.Val)
		c_node = c_node.Next
	}

	fmt.Println("############################")
}

func deleteDuplicates(head *ListNode) *ListNode {
    // Sanity check

	if head == nil {
		return head
	} else if head.Next == nil{
		return head
	}
	
	// Other cases

	var c_node, p_node *ListNode
	p_node = head
	c_node = head.Next

	for c_node != nil{
		if p_node.Val == c_node.Val{
			// Go along the chain using c_node and see how many consecutive nodes have the same value

			for c_node.Next != nil && p_node.Val == c_node.Next.Val{
				c_node = c_node.Next
			}

			if c_node.Next == nil{
				// c_node currently pointing to last index in the list.

				if p_node.Val == c_node.Val{
					p_node.Next = nil
					break
				} else{
					p_node.Next = c_node
					break
				}
			} else if p_node.Val != c_node.Next.Val {
				// c_node.Next != nil and p_node.Val != c_node.Next.Val; This means that p_node.Next must point to c_node.Next

				p_node.Next = c_node.Next
				c_node = c_node.Next
				continue

			}
		}

		p_node = c_node
		c_node = c_node.Next
		
	}

	return head
}

func main(){

	l1 := ListNode{Val: 1}
	l1.Next = &ListNode{Val: 1}
	l1.Next.Next = &ListNode{Val: 2}

	printList(deleteDuplicates(&l1))

	l2 := ListNode{Val: 1}
	l2.Next = &ListNode{Val: 1}
	l2.Next.Next = &ListNode{Val: 2}
	l2.Next.Next.Next = &ListNode{Val: 3}
	l2.Next.Next.Next.Next = &ListNode{Val: 3}

	printList(deleteDuplicates(&l2))

	l3 := ListNode{Val: 1}
	l3.Next = &ListNode{Val: 1}
	l3.Next.Next = &ListNode{Val: 1}
	l3.Next.Next.Next = &ListNode{Val: 1}
	l3.Next.Next.Next.Next = &ListNode{Val: 2}
	l3.Next.Next.Next.Next.Next = &ListNode{Val: 3}

	printList(deleteDuplicates(&l3))

}
