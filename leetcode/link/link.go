package link

import (
	"fmt"
)

type Node struct {
	Val  interface{}
	Next *Node
}

type LinkedList struct {
	head *Node
}

func NewNode(val interface{}, next *Node) *Node {
	return &Node{
		Val:  val,
		Next: next,
	}
}

func NewLinkedList() *LinkedList {
	return &LinkedList{
		head: NewNode(nil, nil),
	}
}

func (l *LinkedList) Empty() bool {
	return l.head.Next == nil
}

func (l *LinkedList) Len() int {
	temp := l.head
	count := -1
	for temp != nil {
		count++
		temp = temp.Next
	}
	return count
}

func (l *LinkedList) Traverse() {
	p := l.head.Next
	for p != nil {
		fmt.Printf("%v", p.Val)
		p = p.Next
	}
}

func (l *LinkedList) Push(val interface{}) {
	node := NewNode(val, nil)

	p := l.head
	for p.Next != nil {
		p = p.Next
	}
	p.Next = node
}

func (l *LinkedList) Insert(index int, val interface{}) bool {
	node := NewNode(val, nil)

	if index < 0 || index > l.Len() {
		return false
	}

	p := l.head

	for i := 0; i < index; i++ {
		p = p.Next
	}

	node.Next = p.Next
	p.Next = node
	return true
}

func (l *LinkedList) Delete(index int) bool {
	if index < 0 || index >= l.Len() {
		return false
	}

	p := l.head
	for i := 0; i < index; i++ {
		p = p.Next
	}

	p.Next = p.Next.Next
	return true
}

func (l *LinkedList) Get(index int) interface{} {
	if node := l.GetNode(index); node == nil {
		return nil
	} else {
		return node.Val
	}
}

func (l *LinkedList) GetNode(index int) *Node {
	if index < 0 || index >= l.Len() {
		return nil
	}

	p := l.head

	for i := -1; i < index; i++ {
		p = p.Next
	}

	return p
}
