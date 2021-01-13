package link

import "fmt"

type DNode struct {
	Val  interface{}
	Prev *DNode
	Next *DNode
}

type DLinkList struct {
	head   *DNode
	tail   *DNode
	length int
}

func NewDLinkList() *DLinkList {
	return &DLinkList{}
}

func (d *DLinkList) Len() int {
	return d.length
}

func (d *DLinkList) Traverse() {
	p := d.head
	for p != nil {
		fmt.Printf("%v\n", p.Val)
		p = p.Next
	}
}

func (d *DLinkList) Append(val interface{}) {
	node := &DNode{Val: val}
	if d.Len() == 0 {
		d.head = node
		d.tail = node
		d.length++
		return
	}

	p := d.tail

	d.tail = node
	p.Next = node
	node.Prev = p
	d.length++
	return
}

func (d *DLinkList) Prepend(val interface{}) {
	node := &DNode{Val: val}
	if d.Len() == 0 {
		d.head = node
		d.tail = node
		d.length++
		return
	}

	p := d.head
	d.head = node
	node.Next = p
	p.Prev = node
	d.length++
	return
}

func (d *DLinkList) Insert(index int, val interface{}) bool {
	if index < 0 || index > d.length {
		return false
	}

	if index == 0 {
		d.Prepend(val)
		d.length++
		return true
	} else if index == d.length {
		d.Append(val)
		d.length++
		return true
	}

	p := d.head
	for i := 0; i < index; i++ {
		p = p.Next
	}
	node := &DNode{Val: val}

	node.Next = p
	node.Prev = p.Prev
	p.Prev.Next = node
	p.Prev = node
	d.length++
	return true
}

func (d *DLinkList) Delete(index int) bool {
	if index < 0 || index > d.length {
		return false
	}

	if index == 0 {
		d.head = d.head.Next
		d.head.Prev = nil
		d.length--
		return true
	} else if index == d.length {
		d.tail = d.tail.Prev
		d.tail.Next = nil
		d.length--
		return true
	}

	p := d.head
	for i := 0; i < index; i++ {
		p = p.Next
	}

	p.Prev.Next = p.Next
	p.Next.Prev = p.Prev
	d.length--
	return true
}
