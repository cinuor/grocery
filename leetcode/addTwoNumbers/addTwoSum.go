package main

import (
    "fmt"
)

type ListNode struct {
    Val int
    Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

    dummyHead := new(ListNode)
    curr := dummyHead
    var carry int = 0
    for ;l1 != nil || l2 != nil;{
        var x, y int;
        if l1 != nil {
            x = l1.Val
        } else {
            x = 0
        }
        if l2 != nil {
            y = l2.Val
        } else {
            y = 0
        }
        sum := x + y + carry
        carry = sum / 10
        curr.Next = &ListNode{Val: sum % 10}
        curr = curr.Next
        if l1 != nil {
            l1 = l1.Next
        }
        if l2 != nil {
            l2 = l2.Next
        }
    }
    if carry > 0 {
        curr.Next = &ListNode{Val: carry}
    }
    return dummyHead.Next
}

func main() {
    a1 := &ListNode{Val: 1}
    a2 := &ListNode{Val: 8}
    a3 := &ListNode{Val: 9}

    a1.Next = a2
    a2.Next = a3

    b1 := &ListNode{Val: 9}
    b2 := &ListNode{Val: 3}

    b1.Next = b2

    for c := addTwoNumbers(a1, b1) ;c != nil; c = c.Next {
        fmt.Printf("%d\n", c.Val)
    }
}
