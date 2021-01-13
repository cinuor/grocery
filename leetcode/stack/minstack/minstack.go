package minstack

import (
	"encoding/json"
	"fmt"
	"math"
	"sync"
)

type node struct {
	val  int
	prev *node
}

type Stack struct {
	top    *node
	min    int
	length int
	lock   sync.Mutex
}

func NewStack() *Stack {
	return &Stack{top: nil, length: 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() int {
	if s.length == 0 {
		return -1
	}
	return s.top.val
}

func (s *Stack) Push(value int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := &node{val: value, prev: s.top}
	s.top = n
	if s.length == 0 {
		s.min = value
	} else {
		s.min = int(math.Min(float64(value), float64(s.min)))
	}
	s.length++
}

func (s *Stack) Pop() int {
	s.lock.Lock()
	defer s.lock.Unlock()

	n := s.top
	s.top = n.prev
	s.length--
	return n.val
}

func (s *Stack) GetMin() int {
	return s.min
}

func (s *Stack) MarshalJSON() ([]byte, error) {
	if s.length == 0 {
		return json.Marshal([]int{})
	}

	var nodes []int

	for i, head := 0, s.top; i < s.length && head != nil; i, head = i+1, head.prev {
		nodes = append(nodes, head.val)
	}

	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	}

	return json.Marshal(nodes)

}

func (s *Stack) UnmarshalJSON(data []byte) error {
	var vals []int

	if err := json.Unmarshal(data, &vals); err != nil {
		return err
	}

	for i := 0; i < len(vals); i++ {
		s.Push(vals[i])
	}

	return nil
}

func main() {
	s := NewStack()
	ss := NewStack()

	s.Push(49)
	s.Push(3)
	s.Push(34)
	s.Push(40)
	s.Push(2)

	data, _ := json.Marshal(s)

	if err := json.Unmarshal(data, ss); err != nil {
		fmt.Println(err.Error())
	}

	for i := 0; i < s.length; i++ {
		fmt.Println(s.top.val)
		s.top = s.top.prev
	}

	fmt.Println(s.GetMin())
}
