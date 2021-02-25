package linkedlist

import (
	"reflect"
	"testing"
)

func TestReverse(t *testing.T) {
	tests := []struct {
		name string
		head *Node
	}{
		{
			name: "testcase 1",
			head: func() *Node {
				head := &Node{}
				a := NewNode(1)
				b := NewNode(2)
				c := NewNode(3)
				d := NewNode(4)
				e := NewNode(5)
				head.Next = a
				a.Next = b
				b.Next = c
				c.Next = d
				d.Next = e
				return head
			}(),
		},
		{
			name: "testcase 2",
			head: func() *Node {
				head := &Node{}
				return head
			}(),
		},
		{
			name: "testcase 3",
			head: func() *Node {
				head := &Node{}
				a := NewNode(1)
				head.Next = a
				return head
			}(),
		},
		{
			name: "testcase 4",
			head: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Print(tt.head)
			got := Reverse(tt.head)
			Print(got)
		})
	}
}

func TestFindMid(t *testing.T) {

	head := &Node{}
	a := NewNode(1)
	b := NewNode(2)
	c := NewNode(3)
	d := NewNode(4)
	e := NewNode(5)
	f := NewNode(6)
	head.Next = a
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = e
	e.Next = f

	tests := []struct {
		name string
		head *Node
		want *Node
	}{
		{
			name: "testcase 1",
			head: head,
			want: d,
		},
		{
			name: "testcase 1",
			head: nil,
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindMid(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReverseHalf(t *testing.T) {
	tests := []struct {
		name string
		head *Node
	}{
		// TODO: Add test cases.
		{
			name: "testcase 1",
			head: func() *Node {
				head := &Node{}
				a := NewNode(1)
				b := NewNode(2)
				c := NewNode(3)
				d := NewNode(4)
				e := NewNode(5)
				f := NewNode(6)
				head.Next = a
				a.Next = b
				b.Next = c
				c.Next = d
				d.Next = e
				e.Next = f
				return head
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseHalf(tt.head)
			Print(got)
		})
	}
}

func Test_isPalindrome(t *testing.T) {
	tests := []struct {
		name string
		head *Node
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "testcase 1",
			head: func() *Node {
				head := &Node{}
				a := NewNode(1)
				b := NewNode(2)
				c := NewNode(3)
				x := NewNode(4)
				d := NewNode(3)
				e := NewNode(2)
				f := NewNode(1)
				head.Next = a
				a.Next = b
				b.Next = c
				c.Next = x
				x.Next = d
				d.Next = e
				e.Next = f
				return head
			}(),
			want: true,
		},
		{
			name: "testcase 2",
			head: func() *Node {
				head := &Node{}
				a := NewNode(1)
				b := NewNode(2)
				c := NewNode(3)
				head.Next = a
				a.Next = b
				b.Next = c
				return head
			}(),
			want: false,
		},
		{
			name: "testcase 2",
			head: func() *Node {
				head := &Node{}
				a := NewNode(1)
				b := NewNode(1)
				c := NewNode(1)
				head.Next = a
				a.Next = b
				b.Next = c
				return head
			}(),
			want: true,
		},
		{
			name: "testcase 2",
			head: func() *Node {
				head := &Node{}
				return head
			}(),
			want: true,
		},
		{
			name: "testcase 2",
			head: func() *Node {
				return nil
			}(),
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Print(tt.head)
			if got := isPalindrome(tt.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
			Print(tt.head)
		})
	}
}

func Test_helpReverse(t *testing.T) {
	a := NewNode(1)
	b := NewNode(2)
	c := NewNode(3)
	a.Next = b
	b.Next = c
	tests := []struct {
		name  string
		head  *Node
		tail  *Node
		want  *Node
		want1 *Node
	}{
		// TODO: Add test cases.
		{
			name:  "testcase 1",
			head:  a,
			tail:  c,
			want:  c,
			want1: a,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Print(tt.head)
			got, got1 := helpReverse(tt.head, tt.tail)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("helpReverse() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("helpReverse() got1 = %v, want %v", got1, tt.want1)
			}
			Print(got)
		})
	}
}
