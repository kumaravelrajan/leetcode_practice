package main

// Definition for singly-linked list.
type ListNode struct {
	Val int
	Next *ListNode
}

func reorderList(head *ListNode)  {

	// Sanity check
	if head == nil || head.Next == nil || head.Next.Next == nil{
		return
	}

    var addresses_of_nodes []*ListNode
	var ct_node *ListNode = head
	// var new_head *ListNode= &ListNode{}
	// var new_ct_node *ListNode = new_head

	for ct_node != nil{
		addresses_of_nodes = append(addresses_of_nodes, ct_node)
		ct_node = ct_node.Next
	}

	i, j, ct_node := 0, len(addresses_of_nodes) - 1, head

	for i <= j{
		if i == 0{
			ct_node.Next = addresses_of_nodes[j]
			ct_node = ct_node.Next

			i++
			j--
			continue
		} else if i < j {
			ct_node.Next = addresses_of_nodes[i]
			ct_node.Next.Next = addresses_of_nodes[j]
			ct_node = ct_node.Next.Next
			
			i++
			j--
		} else if i == j{
			ct_node.Next = addresses_of_nodes[i]
			ct_node = ct_node.Next

			i++
			j--
		}
	}

	if ct_node.Next != nil{
		ct_node.Next = nil
	}
}

func main(){
	
	var l1 ListNode = ListNode{Val: 1}
	l1.Next = &ListNode{Val: 2}
	l1.Next.Next = &ListNode{Val: 3}
	l1.Next.Next.Next = &ListNode{Val: 4}
	reorderList(&l1)

	var l2 ListNode = ListNode{Val: 1}
	l2.Next = &ListNode{Val: 2}
	l2.Next.Next = &ListNode{Val: 3}
	l2.Next.Next.Next = &ListNode{Val: 4}
	l2.Next.Next.Next.Next = &ListNode{Val: 5}
	reorderList(&l2)
}