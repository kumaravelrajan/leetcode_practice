package main

type ListNode struct{
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {

    if head == nil || head.Next == nil{
        return head
    }

    dummy := &ListNode{}
    dummy.Next = head

    back, front := dummy, head.Next

    for front != nil{
        // printList(&back)
        if back.Next.Val == front.Val {
            front = front.Next
        } else {
            back.Next = front
            back = front

            if front.Next != nil{
                front = front.Next.Next
            } else {
                break
            }
        }
    }

    return dummy.Next
    
}

func main(){
	one1 := ListNode{Val: 1}
	one2 := ListNode{Val: 1}
	one3 := ListNode{Val: 1}
	two := ListNode{Val: 2}
	three := ListNode{Val: 3}

	one1.Next = &one2
	one2.Next = &one3
	one3.Next = &two
	two.Next = &three

	deleteDuplicates(&one1)
}
