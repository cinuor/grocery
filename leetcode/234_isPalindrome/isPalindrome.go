package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	var prev, cur *ListNode = nil, head

	for cur != nil {
		r := cur.Next
		cur.Next = prev

		prev = cur
		cur = r
	}
	return prev
}

func findHalf(head *ListNode) *ListNode {

	slow, fast := head, head

	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow
}

func isPalindrome(head *ListNode) bool {

	half := findHalf(head)

	left := head
	right := reverse(half.Next)

	result := true

	for result && right != nil {
		if left.Val != right.Val {
			result = false
		}
		left = left.Next
		right = right.Next
	}
	half.Next = reverse(right)
	return result
}

func main() {
	a := &ListNode{1, nil}
	b := &ListNode{2, nil}
	x := &ListNode{3, nil}
	c := &ListNode{2, nil}
	d := &ListNode{1, nil}

	a.Next = b
	b.Next = x
	x.Next = c
	c.Next = d

	fmt.Println(isPalindrome(a))
}
