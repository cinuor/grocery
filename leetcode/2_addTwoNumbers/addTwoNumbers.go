package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carry, total int

	n := new(ListNode)
	head := n

	for carry != 0 || l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			total = l1.Val + l2.Val + carry
			l1 = l1.Next
			l2 = l2.Next
		} else if l1 != nil && l2 == nil {
			total = l1.Val + carry
			l1 = l1.Next
		} else if l1 == nil && l2 != nil {
			total = l2.Val + carry
			l2 = l2.Next
		} else {
			total = carry
		}

		remainder := total % 10
		carry = total / 10

		n.Next = &ListNode{Val: remainder, Next: nil}
		n = n.Next

		total = 0
	}
	return head.Next
}
