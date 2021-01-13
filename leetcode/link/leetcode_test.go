package link

import (
	"fmt"
	"testing"
)

func createList(n int) *ListNode {
	dummy := new(ListNode)
	cur := dummy

	for i := 0; i < n; i++ {
		cur.Next = &ListNode{Val: i + 1}
		cur = cur.Next
	}

	return dummy.Next
}

func Traverse(n *ListNode) {
	for n != nil {
		fmt.Printf("%v\t", n.Val)
		n = n.Next
	}
}

func TestAddTwoNumbers(t *testing.T) {
	a := &ListNode{Val: 2}
	b := &ListNode{Val: 4}
	c := &ListNode{Val: 3}
	a.Next = b
	b.Next = c

	d := &ListNode{Val: 5}
	e := &ListNode{Val: 6}
	f := &ListNode{Val: 4}

	d.Next = e
	e.Next = f

	r := addTwoNumbers(a, d)
	Traverse(r)

}

func TestReverseR(t *testing.T) {
	a := createList(5)
	r := Reverse(a)

	Traverse(r)
}

func TestReverseKGroup(t *testing.T) {
	a := createList(1)

	r := reverseKGroup(a, 1)

	Traverse(r)
}
