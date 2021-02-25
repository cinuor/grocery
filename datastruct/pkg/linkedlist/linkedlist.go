package linkedlist

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func NewNode(val int) *Node {
	return &Node{Val: val}
}

func Print(head *Node) {
	if head == nil {
		return
	}
	p := head
	for p != nil {
		if p.Val == -1 {
			fmt.Printf("head->")
		}
		fmt.Printf("%v->", p.Val)
		p = p.Next
	}
	fmt.Println("nil")
}

// 迭代反转链表
func Reverse(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	p := head
	q := p.Next
	for q != nil {
		r := q.Next
		q.Next = p
		p, q = q, r
	}
	head.Next.Next = nil
	head.Next = p
	return head
}

// 找到链表中点
func FindMid(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}

	fast, slow := head.Next, head.Next
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 反转后半部分链表 (判断回文链表)
func ReverseHalf(head *Node) *Node {
	if head == nil || head.Next == nil {
		return head
	}
	half := FindMid(head)
	Reverse(half)
	return head
}

// 判断回文链表
func isPalindrome(head *Node) bool {
	if head == nil || head.Next == nil {
		return true
	}
	mid := FindMid(head)
	Reverse(mid)
	p := mid.Next
	q := head.Next

	for p != nil {
		if q.Val != p.Val {
			return false
		}
		p = p.Next
		q = q.Next
	}
	Reverse(mid)
	return true
}

/*
 * // K个一组反转链表
 * func ReverseK(p *Node, k int) *Node {
 *     if p == nil || p.Next == nil {
 *         return p
 *     }
 *
 *     // TODO
 *
 *     p.Next = ReverseK(, k)
 *
 * }
 */

func helpReverse(head, tail *Node) (*Node, *Node) {
	p := head

	for p != tail {
		q := p.Next
		r := q.Next

		q.Next = p
		p, q = q, r
	}
	head.Next = p.Next
	return p, head
}
