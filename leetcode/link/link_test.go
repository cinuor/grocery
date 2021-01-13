package link

import (
	"fmt"
	"testing"
)

func TestEmpty(t *testing.T) {
	l := NewLinkedList()
	fmt.Println(l.Empty())
}

func TestLen(t *testing.T) {
	l := NewLinkedList()
	fmt.Println(l.Len())
}

func TestPush(t *testing.T) {
	l := NewLinkedList()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Traverse()
}

func TestInsert(t *testing.T) {
	l := NewLinkedList()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Insert(4, 9)
	l.Traverse()
}

func TestDelete(t *testing.T) {
	l := NewLinkedList()
	l.Push(1)
	l.Push(2)
	l.Push(3)
	l.Push(4)
	l.Delete(3)
	l.Traverse()
}
