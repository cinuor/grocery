package lru

import (
	"fmt"
	"strings"
)

type DoubleList struct {
	Head *Node
	Tail *Node
	size int
}

func NewDoubleList() *DoubleList {
	d := &DoubleList{size: 0}
	d.Head = NewNode("", 0)
	d.Tail = NewNode("", 0)
	d.Head.Next = d.Tail
	d.Tail.Prev = d.Head

	return d
}

func (d *DoubleList) Size() int {
	return d.size
}

func (d *DoubleList) AddLast(n *Node) {
	n.Prev = d.Tail.Prev
	n.Next = d.Tail
	n.Prev.Next = n
	n.Next.Prev = n
	d.size++
}

func (d *DoubleList) Remove(n *Node) {
	n.Prev.Next = n.Next
	n.Next.Prev = n.Prev
	d.size--
}

func (d *DoubleList) RemoveFirst() *Node {
	if d.Size() == 0 {
		return nil
	}

	n := d.Head.Next
	d.Remove(n)
	return n
}

func (d *DoubleList) String() string {
	s := make([]string, 0)

	h := d.Head.Next
	t := d.Tail
	for h != t {
		s = append(s, fmt.Sprintf("(%s, %d)", h.Key, h.Value))
		h = h.Next
	}

	return strings.Join(s, "->")
}
