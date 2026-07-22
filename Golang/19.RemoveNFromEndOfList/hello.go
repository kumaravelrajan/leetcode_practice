package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	if head.Next == nil{
		return nil
	}

    front := head
    back := head
    i := 0

    for front != nil{
        if i > n {
            front = front.Next
            back = back.Next
            i++
        } else {
            front = front.Next
            i++
        }

        // fmt.Println("back.Val = ", back.Val)
    }

	if back.Next != nil && back.Next.Next != nil{
		back.Next = back.Next.Next
	} else if back.Next != nil && back.Next.Next ==  nil{
		back.Next = nil
	}

    return head
    
}

func main(){
	one := ListNode{Val: 1}
	two := ListNode{Val: 2}
	three := ListNode{Val: 3}
	four := ListNode{Val: 4}
	five := ListNode{Val: 5}

	one.Next = &two
	two.Next = &three
	three.Next = &four
	four.Next = &five

	newList := removeNthFromEnd(&one, 2)

	for newList != nil{
		fmt.Println(newList.Val)
		newList = newList.Next
	}
}
